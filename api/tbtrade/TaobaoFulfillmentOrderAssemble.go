package tbtrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// TaobaoFulfillmentOrderAssemble 拆合单结果回传接口
// taobao.fulfillment.order.assemble
//
// 拆合单结果回传接口
func TaobaoFulfillmentOrderAssemble(clt *core.SDKClient, req *tbtrade.TaobaoFulfillmentOrderAssembleAPIRequest, session string) (*tbtrade.TaobaoFulfillmentOrderAssembleAPIResponse, error) {
	var resp tbtrade.TaobaoFulfillmentOrderAssembleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
