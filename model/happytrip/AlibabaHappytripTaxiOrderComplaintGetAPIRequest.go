package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHappytripTaxiOrderComplaintGetAPIRequest 投诉详情 API请求
// alibaba.happytrip.taxi.order.complaint.get
//
// 获取投诉处理进度详情
type AlibabaHappytripTaxiOrderComplaintGetAPIRequest struct {
	model.Params
	// 供应商订单号
	_orderId string
	// 供应商工单号
	_caseId string
}

// NewAlibabaHappytripTaxiOrderComplaintGetRequest 初始化AlibabaHappytripTaxiOrderComplaintGetAPIRequest对象
func NewAlibabaHappytripTaxiOrderComplaintGetRequest() *AlibabaHappytripTaxiOrderComplaintGetAPIRequest {
	return &AlibabaHappytripTaxiOrderComplaintGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiOrderComplaintGetAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.order.complaint.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiOrderComplaintGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is OrderId Setter
// 供应商订单号
func (r *AlibabaHappytripTaxiOrderComplaintGetAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// Get OrderId Getter
func (r AlibabaHappytripTaxiOrderComplaintGetAPIRequest) GetOrderId() string {
	return r._orderId
}

// Set is CaseId Setter
// 供应商工单号
func (r *AlibabaHappytripTaxiOrderComplaintGetAPIRequest) SetCaseId(_caseId string) error {
	r._caseId = _caseId
	r.Set("case_id", _caseId)
	return nil
}

// Get CaseId Getter
func (r AlibabaHappytripTaxiOrderComplaintGetAPIRequest) GetCaseId() string {
	return r._caseId
}
