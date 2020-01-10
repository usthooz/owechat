package cache

const (
	// defaultTokenCacheKey 默认的accesstoken缓存的key
	defaultTokenCacheKey = "owechat:cache:access_token"
	// defaultTicketCacheKey 默认的ticket缓存的key
	defaultTicketCacheKey = "owechat:cache:ticket"
)

var (
	// accessTokenKey accessToken 缓存key
	accessTokenKey string
	// ticketCacheKey ticket 缓存key
	ticketCacheKey string
)
