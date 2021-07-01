package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaHappytripTaxiOrderComplaintGetAPIRequest
投诉详情 API请求
alibaba.happytrip.taxi.order.complaint.get

获取投诉处理进度详情 */
type AlibabaHappytripTaxiOrderComplaintGetAPIRequest struct {
	model.Params
	// 供应商订单号
	_orderId string
	// 供应商工单号
	_caseId string
}

// New
