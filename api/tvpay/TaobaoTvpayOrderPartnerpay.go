package tvpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvpay"
)

// Taobaotvpayorderpartnerpay tv支付第三方支付订单
// taobao.tvpay.order.partnerpay
//
// tv支付第三方发起并支付订单（使用设备授权）
func Taobaotvpayorderpartnerpay(clt *core.SDKClient, req *tvpay.TaobaotvpayorderpartnerpayAPIRequest, session string) (*tvpay.TaobaotvpayorderpartnerpayAPIResponse, error) {
	var resp tvpay.TaobaotvpayorderpartnerpayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
