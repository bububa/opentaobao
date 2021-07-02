package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRecycleOrderShow 闲鱼回收订单查询V1.1
// alibaba.idle.recycle.order.show
//
// 查询回收订单
func AlibabaIdleRecycleOrderShow(clt *core.SDKClient, req *idle.AlibabaIdleRecycleOrderShowAPIRequest, session string) (*idle.AlibabaIdleRecycleOrderShowAPIResponse, error) {
	var resp idle.AlibabaIdleRecycleOrderShowAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
