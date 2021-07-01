package opentrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/opentrade"
)

/* TaobaoOpentradeCreateOrder
订单创建
taobao.opentrade.create.order

交易开放创建订单 */
func TaobaoOpentradeCreateOrder(clt *core.SDKClient, req *opentrade.TaobaoOpentradeCreateOrderAPIRequest, session string) (*opentrade.TaobaoOpentradeCreateOrderAPIResponse, error) {
	var resp opentrade.TaobaoOpentradeCreateOrderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
