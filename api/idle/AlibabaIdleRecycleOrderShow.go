package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRecycleOrderShow 闲鱼回收订单查询V1.1
// alibaba.idle.recycle.order.show
//
// 查询回收订单
func AlibabaIdleRecycleOrderShow(clt *core.SDKClient, req *idle.AlibabaIdleRecycleOrderShowAPIRequest, resp *idle.AlibabaIdleRecycleOrderShowAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
