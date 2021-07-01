package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleRentOrderLogisticsDeliverAPIRequest
创建揽收物流 API请求
alibaba.idle.rent.order.logistics.deliver

创建揽收物流
商家去物流公司创建物流订单 */
type AlibabaIdleRentOrderLogisticsDeliverAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
	// 物流信息
	_logistics *LogisticsDto
}

// New
