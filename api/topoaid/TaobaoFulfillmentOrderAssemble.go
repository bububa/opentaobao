package topoaid

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/topoaid"
)

// TaobaoFulfillmentOrderAssemble 拆合单结果回传接口
// taobao.fulfillment.order.assemble
//
// 拆合单结果回传接口
func TaobaoFulfillmentOrderAssemble(clt *core.SDKClient, req *topoaid.TaobaoFulfillmentOrderAssembleAPIRequest, session string) (*topoaid.TaobaoFulfillmentOrderAssembleAPIResponse, error) {
	var resp topoaid.TaobaoFulfillmentOrderAssembleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
