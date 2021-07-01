package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleRentOrderPackageAPIRequest
确认揽收商品 API请求
alibaba.idle.rent.order.package

确认揽收 */
type AlibabaIdleRentOrderPackageAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
	// 物流信息
	_logistics *LogisticsDto
}

// New
