package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleRentOrderSenditemAPIRequest
确认发货 API请求
alibaba.idle.rent.order.senditem

确认发货 */
type AlibabaIdleRentOrderSenditemAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
	// 物流信息
	_logisticsList []LogisticsDto
}

// New
