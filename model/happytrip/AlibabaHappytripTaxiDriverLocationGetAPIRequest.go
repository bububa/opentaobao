package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHappytripTaxiDriverLocationGetAPIRequest
司机位置 API请求
alibaba.happytrip.taxi.driver.location.get

获取司机实时位置 */
type AlibabaHappytripTaxiDriverLocationGetAPIRequest struct {
	model.Params
	// 供应商订单号
	_orderId string
}

// New
