package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHappytripTaxiIdGetAPIRequest
获取请求id API请求
alibaba.happytrip.taxi.id.get

获取订单号 */
type AlibabaHappytripTaxiIdGetAPIRequest struct {
	model.Params
	// 用户唯一标识
	_uid string
}

// New
