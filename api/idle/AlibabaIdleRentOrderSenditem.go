package idle

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idle"
)

// AlibabaIdleRentOrderSenditem 确认发货
// alibaba.idle.rent.order.senditem
//
// 确认发货
func AlibabaIdleRentOrderSenditem(clt *core.SDKClient, req *idle.AlibabaIdleRentOrderSenditemAPIRequest, resp *idle.AlibabaIdleRentOrderSenditemAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
