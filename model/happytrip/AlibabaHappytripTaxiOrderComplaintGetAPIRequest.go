package happytrip

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHappytripTaxiOrderComplaintGetAPIRequest) Reset() {
	r._orderId = ""
	r._caseId = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHappytripTaxiOrderComplaintGetAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.order.complaint.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHappytripTaxiOrderComplaintGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHappytripTaxiOrderComplaintGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 供应商订单号
func (r *AlibabaHappytripTaxiOrderComplaintGetAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaHappytripTaxiOrderComplaintGetAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetCaseId is CaseId Setter
// 供应商工单号
func (r *AlibabaHappytripTaxiOrderComplaintGetAPIRequest) SetCaseId(_caseId string) error {
	r._caseId = _caseId
	r.Set("case_id", _caseId)
	return nil
}

// GetCaseId CaseId Getter
func (r AlibabaHappytripTaxiOrderComplaintGetAPIRequest) GetCaseId() string {
	return r._caseId
}

var poolAlibabaHappytripTaxiOrderComplaintGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHappytripTaxiOrderComplaintGetRequest()
	},
}

// GetAlibabaHappytripTaxiOrderComplaintGetRequest 从 sync.Pool 获取 AlibabaHappytripTaxiOrderComplaintGetAPIRequest
func GetAlibabaHappytripTaxiOrderComplaintGetAPIRequest() *AlibabaHappytripTaxiOrderComplaintGetAPIRequest {
	return poolAlibabaHappytripTaxiOrderComplaintGetAPIRequest.Get().(*AlibabaHappytripTaxiOrderComplaintGetAPIRequest)
}

// ReleaseAlibabaHappytripTaxiOrderComplaintGetAPIRequest 将 AlibabaHappytripTaxiOrderComplaintGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaHappytripTaxiOrderComplaintGetAPIRequest(v *AlibabaHappytripTaxiOrderComplaintGetAPIRequest) {
	v.Reset()
	poolAlibabaHappytripTaxiOrderComplaintGetAPIRequest.Put(v)
}
