package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHappytripTaxiOrderCancelAPIRequest
取消叫车 API请求
alibaba.happytrip.taxi.order.cancel

取消叫车订单,行程中的订单不能取消 */
type AlibabaHappytripTaxiOrderCancelAPIRequest struct {
	model.Params
	// 订单id
	_orderId string
	// 是否强制取消(true 或 false)默认false
	_force string
	// 取消类型，0：系统取消，非0：用户取消
	_type int64
}

// New
