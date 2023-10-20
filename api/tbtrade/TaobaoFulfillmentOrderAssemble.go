package tbtrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbtrade"
)

// TaobaoFulfillmentOrderAssemble 拆合单结果回传接口
// taobao.fulfillment.order.assemble
//
// 拆合单结果回传接口
func TaobaoFulfillmentOrderAssemble(clt *core.SDKClient, req *tbtrade.TaobaoFulfillmentOrderAssembleAPIRequest, resp *tbtrade.TaobaoFulfillmentOrderAssembleAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
