package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRentOrderReceiveitem 确认签收
// alibaba.idle.rent.order.receiveitem
//
// 确认揽收/签收
func AlibabaIdleRentOrderReceiveitem(clt *core.SDKClient, req *idle.AlibabaIdleRentOrderReceiveitemAPIRequest, resp *idle.AlibabaIdleRentOrderReceiveitemAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
