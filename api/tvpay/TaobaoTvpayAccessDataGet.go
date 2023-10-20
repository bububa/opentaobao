package tvpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvpay"
)

// TaobaoTvpayAccessDataGet tv支付
// taobao.tvpay.access.data.get
//
// 在匿名用户支付后尝试为其登陆绑定的淘宝账号
func TaobaoTvpayAccessDataGet(clt *core.SDKClient, req *tvpay.TaobaoTvpayAccessDataGetAPIRequest, resp *tvpay.TaobaoTvpayAccessDataGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
