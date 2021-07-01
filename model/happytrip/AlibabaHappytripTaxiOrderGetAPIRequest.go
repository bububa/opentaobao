package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHappytripTaxiOrderGetAPIRequest
订单详情 API请求
alibaba.happytrip.taxi.order.get

获取订单状态及详情 */
type AlibabaHappytripTaxiOrderGetAPIRequest struct {
	model.Params
	// 订单id
	_orderId string
}

// New
