package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// Alibabaidlerecycleorderfulfillment 闲鱼回收订单履约V1
// alibaba.idle.recycle.order.fulfillment
//
// 外部回收商针对自有回收订单的履行
func Alibabaidlerecycleorderfulfillment(clt *core.SDKClient, req *idle.AlibabaidlerecycleorderfulfillmentAPIRequest, session string) (*idle.AlibabaidlerecycleorderfulfillmentAPIResponse, error) {
	var resp idle.AlibabaidlerecycleorderfulfillmentAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
