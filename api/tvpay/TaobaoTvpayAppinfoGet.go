package tvpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvpay"
)

// TaobaoTvpayAppinfoGet tv支付获取应用信息
// taobao.tvpay.appinfo.get
//
// tv支付获取应用信息
func TaobaoTvpayAppinfoGet(clt *core.SDKClient, req *tvpay.TaobaoTvpayAppinfoGetAPIRequest, resp *tvpay.TaobaoTvpayAppinfoGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
