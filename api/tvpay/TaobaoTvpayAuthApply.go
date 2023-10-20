package tvpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvpay"
)

// Taobaotvpayauthapply tv支付申请设备授权
// taobao.tvpay.auth.apply
//
// 为用户在指定设备上申请支付授权
func Taobaotvpayauthapply(clt *core.SDKClient, req *tvpay.TaobaotvpayauthapplyAPIRequest, session string) (*tvpay.TaobaotvpayauthapplyAPIResponse, error) {
	var resp tvpay.TaobaotvpayauthapplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
