package pay

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	"github.com/usthooz/gutil"
	"github.com/usthooz/owechat/config"
	"github.com/usthooz/owechat/util"
)

// SendPack 发放成功: return nil,nil 发放失败: return 查询结果,err/nil
func SendPack(sp *SendredPackParams) (*SendredPackResp, error) {
	// 获取clien ip
	ip, err := gutil.GetLocalPublicIp()
	if err != nil {
		return nil, err
	}
	// 订单ID
	if len(sp.MchBillno) == 0 {
		sp.MchBillno = fmt.Sprintf("%s-%s-%d", cfg.BaseConf.MchId, sp.ReOpenid, time.Now().Unix())
	}
	// 封装请求参数
	params := &SendredPackReq{
		NonceStr:    gutil.RandString(16),
		MchBillno:   sp.MchBillno,
		MchId:       cfg.BaseConf.MchId,
		Wxappid:     cfg.BaseConf.Appid,
		SendName:    sp.SendName,
		ReOpenid:    sp.ReOpenid,
		TotalAmount: sp.TotalAmount,
		TotalNum:    sp.TotalNum,
		Wishing:     sp.Wishing,
		ClientIp:    ip,
		ActName:     sp.ActName,
		Remark:      sp.Remark,
		SceneId:     sp.SceneId,
	}
	// 处理活动信息
	if sp.RiskInfo != nil {
		risk := sp.RiskInfo
		params.RiskInfo = fmt.Sprintf("posttime=%d&clientversion=%s&mobile=%s&deviceid=%s", risk.Posttime, risk.ClientVersion, risk.Mobile, risk.Deviceid)
	}
	// 请求参数签名
	params.Sign = calcSign(map[string]interface{}{
		"nonce_str":    params.NonceStr,
		"mch_billno":   params.MchBillno,
		"mch_id":       params.MchId,
		"wxappid":      params.Wxappid,
		"send_name":    params.SendName,
		"re_openid":    params.ReOpenid,
		"total_amount": params.TotalAmount,
		"total_num":    params.TotalNum,
		"wishing":      params.Wishing,
		"client_ip":    params.ClientIp,
		"act_name":     params.ActName,
		"remark":       params.Remark,
		"scene_id":     params.SceneId,
		"risk_info":    params.RiskInfo,
	}, cfg.BaseConf.PayKey)

	var (
		bytesReq []byte
		res      []byte
	)

	bytesReq, err = xml.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("SendPack: %s", err.Error())
	}

	// 发起请求
	res, err = util.DoPostXml(SendPackUrl, strings.Replace(string(bytesReq), "SendredPackReq", "xml", -1))
	if err != nil {
		return nil, fmt.Errorf("SendPack: %s", err.Error())
	}

	// 解析响应
	xmlResp := SendredPackResp{}
	err = xml.Unmarshal(res, &xmlResp)
	if err != nil {
		return nil, fmt.Errorf("SendPack: %s", err.Error())
	}

	return &xmlResp, nil
}

// QueryPack
func QueryPack(mchBillno, billType string) (*QueryPackResp, error) {
	// 封装请求参数
	params := &QueryPackReq{
		NonceStr:  gutil.RandString(16),
		MchBillno: mchBillno,
		MchId:     cfg.BaseConf.MchId,
		Appid:     cfg.BaseConf.Appid,
		BillType:  billType,
	}
	// 签名
	params.Sign = calcSign(map[string]interface{}{
		"nonce_str":  params.NonceStr,
		"mch_billno": params.MchBillno,
		"mch_id":     params.MchId,
		"appid":      params.Appid,
		"bill_type":  params.BillType,
	}, cfg.BaseConf.PayKey)

	var (
		bytesReq []byte
		res      []byte
	)
	bytesReq, err := xml.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("QueryPack: %s", err.Error())
	}

	// 发起请求
	res, err = util.DoPostXml(QueryPackUrl, strings.Replace(string(bytesReq), "QueryPackReq", "xml", -1))
	if err != nil {
		return nil, fmt.Errorf("QueryPack: %s", err.Error())
	}

	// 解析响应
	xmlResp := QueryPackResp{}
	err = xml.Unmarshal(res, &xmlResp)
	if err != nil {
		return nil, fmt.Errorf("QueryPack: %s", err.Error())
	}
	return &xmlResp, nil
}
