package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRecycleOrderFulfillment 闲鱼回收订单履约V1
// alibaba.idle.recycle.order.fulfillment
//
// 外部回收商针对自有回收订单的履行
func AlibabaIdleRecycleOrderFulfillment(clt *core.SDKClient, req *idle.AlibabaIdleRecycleOrderFulfillmentAPIRequest, session string) (*idle.AlibabaIdleRecycleOrderFulfillmentAPIResponse, error) {
	var resp idle.AlibabaIdleRecycleOrderFulfillmentAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
