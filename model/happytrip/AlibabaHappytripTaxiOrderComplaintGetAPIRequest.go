package happytrip

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahappytriptaxiordercomplaintgetAPIRequest 投诉详情 API请求
// alibaba.happytrip.taxi.order.complaint.get
//
// 获取投诉处理进度详情
type AlibabahappytriptaxiordercomplaintgetAPIRequest struct {
	model.Params
	// 供应商订单号
	_orderId string
	// 供应商工单号
	_caseId string
}

// NewAlibabahappytriptaxiordercomplaintgetRequest 初始化AlibabahappytriptaxiordercomplaintgetAPIRequest对象
func NewAlibabahappytriptaxiordercomplaintgetRequest() *AlibabahappytriptaxiordercomplaintgetAPIRequest {
	return &AlibabahappytriptaxiordercomplaintgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahappytriptaxiordercomplaintgetAPIRequest) GetApiMethodName() string {
	return "alibaba.happytrip.taxi.order.complaint.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahappytriptaxiordercomplaintgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahappytriptaxiordercomplaintgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 供应商订单号
func (r *AlibabahappytriptaxiordercomplaintgetAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabahappytriptaxiordercomplaintgetAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetCaseId is CaseId Setter
// 供应商工单号
func (r *AlibabahappytriptaxiordercomplaintgetAPIRequest) SetCaseId(_caseId string) error {
	r._caseId = _caseId
	r.Set("case_id", _caseId)
	return nil
}

// GetCaseId CaseId Getter
func (r AlibabahappytriptaxiordercomplaintgetAPIRequest) GetCaseId() string {
	return r._caseId
}
