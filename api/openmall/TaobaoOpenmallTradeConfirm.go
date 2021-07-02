package openmall

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallTradeConfirm 确认收货
// taobao.openmall.trade.confirm
//
// 确认订单收货
func TaobaoOpenmallTradeConfirm(clt *core.SDKClient, req *openmall.TaobaoOpenmallTradeConfirmAPIRequest, session string) (*openmall.TaobaoOpenmallTradeConfirmAPIResponse, error) {
	var resp openmall.TaobaoOpenmallTradeConfirmAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
