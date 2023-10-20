package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// Taobaoopenmalltradeagreepay openmall订单支付
// taobao.openmall.trade.agreepay
//
// openmall订单支付
func Taobaoopenmalltradeagreepay(clt *core.SDKClient, req *openmall.TaobaoopenmalltradeagreepayAPIRequest, session string) (*openmall.TaobaoopenmalltradeagreepayAPIResponse, error) {
	var resp openmall.TaobaoopenmalltradeagreepayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
