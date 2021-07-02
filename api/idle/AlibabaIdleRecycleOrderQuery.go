package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRecycleOrderQuery 闲鱼回收订单查询V1
// alibaba.idle.recycle.order.query
//
// 查询回收订单
func AlibabaIdleRecycleOrderQuery(clt *core.SDKClient, req *idle.AlibabaIdleRecycleOrderQueryAPIRequest, session string) (*idle.AlibabaIdleRecycleOrderQueryAPIResponse, error) {
	var resp idle.AlibabaIdleRecycleOrderQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
