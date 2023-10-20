package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// Taobaoopenmalltradeconfirm 确认收货
// taobao.openmall.trade.confirm
//
// 确认订单收货
func Taobaoopenmalltradeconfirm(clt *core.SDKClient, req *openmall.TaobaoopenmalltradeconfirmAPIRequest, session string) (*openmall.TaobaoopenmalltradeconfirmAPIResponse, error) {
	var resp openmall.TaobaoopenmalltradeconfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
