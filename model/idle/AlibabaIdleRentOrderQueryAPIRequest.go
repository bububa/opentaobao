package idle

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaIdleRentOrderQueryAPIRequest
查询订单 API请求
alibaba.idle.rent.order.query

查询订单信息 */
type AlibabaIdleRentOrderQueryAPIRequest struct {
	model.Params
	// 订单id
	_orderId int64
}

// New
