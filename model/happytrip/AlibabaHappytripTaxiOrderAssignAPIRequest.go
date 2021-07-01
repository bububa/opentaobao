package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHappytripTaxiOrderAssignAPIRequest
订单指派 API请求
alibaba.happytrip.taxi.order.assign

通知供应商订单指派成功 */
type AlibabaHappytripTaxiOrderAssignAPIRequest struct {
	model.Params
	// 供应商订单号
	_orderId string
}

// New
