package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHappytripTaxiOrderConfirmAPIRequest
费用确认 API请求
alibaba.happytrip.taxi.order.confirm

1.司机点结束计费,欢行会收到正常支付待评论 回调,确认费用无误欢行可以通过此接口确认并支付。
2.如果欢行一直不调用此接口,订单会在48小时后自动支付。 */
type AlibabaHappytripTaxiOrderConfirmAPIRequest struct {
	model.Params
	// 要确认支付的订单号
	_orderId string
}

// New
