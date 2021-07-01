package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleRentOrderReceiveitemAPIRequest
确认签收 API请求
alibaba.idle.rent.order.receiveitem

确认揽收/签收 */
type AlibabaIdleRentOrderReceiveitemAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
}

// New
