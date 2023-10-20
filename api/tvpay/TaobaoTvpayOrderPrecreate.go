package tvpay

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tvpay"
)

// Taobaotvpayorderprecreate tv支付预下单
// taobao.tvpay.order.precreate
//
// tv支付预下单
func Taobaotvpayorderprecreate(clt *core.SDKClient, req *tvpay.TaobaotvpayorderprecreateAPIRequest, session string) (*tvpay.TaobaotvpayorderprecreateAPIResponse, error) {
	var resp tvpay.TaobaotvpayorderprecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
