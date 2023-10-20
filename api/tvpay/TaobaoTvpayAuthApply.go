package tvpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvpay"
)

// TaobaoTvpayAuthApply tv支付申请设备授权
// taobao.tvpay.auth.apply
//
// 为用户在指定设备上申请支付授权
func TaobaoTvpayAuthApply(clt *core.SDKClient, req *tvpay.TaobaoTvpayAuthApplyAPIRequest, resp *tvpay.TaobaoTvpayAuthApplyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
