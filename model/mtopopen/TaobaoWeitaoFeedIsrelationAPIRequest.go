package mtopopen

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoWeitaoFeedIsrelationAPIRequest
是否关注 API请求
taobao.weitao.feed.isrelation

判断用户是否关注对应的公共账号 */
type TaobaoWeitaoFeedIsrelationAPIRequest struct {
	model.Params
	// 要查询的粉丝的淘宝昵称
	_fansNick string
	// 要查询的公共账号的淘宝昵称
	_sellerNick string
}

// New
