package pay

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/usthooz/owechat/config"
)

// nonceStr 用时间戳生成随机字符串
func nonceStr() string {
	return strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
}

// calcSign 微信支付计算签名的函数
func calcSign(mReq map[string]interface{}, key string) string {
	// STEP 1, 对key进行升序排序.
	sorted_keys := make([]string, 0)
	for k, _ := range mReq {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)

	// STEP2, 对key=value的键值对用&连接起来，略过空值
	var (
		signStrings string
	)
	for _, k := range sorted_keys {
		value := fmt.Sprintf("%v", mReq[k])
		if value != "" && value != OmitIntString {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}

	// STEP3, 在键值对的最后加上key=API_KEY
	if key != "" {
		signStrings = signStrings + "key=" + key
	}

	// STEP4, 进行MD5签名并且将所有字符转为大写.
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := strings.ToUpper(hex.EncodeToString(cipherStr))
	return upperSign
}

// VerifySign 微信支付签名验证函数
func VerifySign(needVerifyM map[string]interface{}, sign string) bool {
	signCalc := calcSign(needVerifyM, cfg.BaseConf.PayKey)
	return sign == signCalc
}
