package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdentalstatementqueryAPIRequest ISV查询对账单 API请求
// alibaba.alihealth.dental.statement.query
//
// ISV查询对账单
type AlibabaalihealthdentalstatementqueryAPIRequest struct {
	model.Params
	// 订单ID
	_orderId string
	// 结算周期，单位月
	_statementTime string
}

// NewAlibabaalihealthdentalstatementqueryRequest 初始化AlibabaalihealthdentalstatementqueryAPIRequest对象
func NewAlibabaalihealthdentalstatementqueryRequest() *AlibabaalihealthdentalstatementqueryAPIRequest {
	return &AlibabaalihealthdentalstatementqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdentalstatementqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.dental.statement.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdentalstatementqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdentalstatementqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单ID
func (r *AlibabaalihealthdentalstatementqueryAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaalihealthdentalstatementqueryAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetStatementTime is StatementTime Setter
// 结算周期，单位月
func (r *AlibabaalihealthdentalstatementqueryAPIRequest) SetStatementTime(_statementTime string) error {
	r._statementTime = _statementTime
	r.Set("statement_time", _statementTime)
	return nil
}

// GetStatementTime StatementTime Getter
func (r AlibabaalihealthdentalstatementqueryAPIRequest) GetStatementTime() string {
	return r._statementTime
}
