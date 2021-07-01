package blackvip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoBlackvipUserinfoGetAPIRequest
88VIP用户信息查询 API请求
taobao.blackvip.userinfo.get

查询88VIP用户信息，比如用户是否是88VIP，88VIP的失效时间等 */
type TaobaoBlackvipUserinfoGetAPIRequest struct {
	model.Params
}

// New
