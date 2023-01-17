package alihealth2

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDentalStatementQueryAPIRequest ISV查询对账单 API请求
// alibaba.alihealth.dental.statement.query
//
// ISV查询对账单
type AlibabaAlihealthDentalStatementQueryAPIRequest struct {
	model.Params
	// 订单ID
	_orderId string
	// 结算周期，单位月
	_statementTime string
}

// NewAlibabaAlihealthDentalStatementQueryRequest 初始化AlibabaAlihealthDentalStatementQueryAPIRequest对象
func NewAlibabaAlihealthDentalStatementQueryRequest() *AlibabaAlihealthDentalStatementQueryAPIRequest {
	return &AlibabaAlihealthDentalStatementQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDentalStatementQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.dental.statement.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDentalStatementQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDentalStatementQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetOrderId is OrderId Setter
// 订单ID
func (r *AlibabaAlihealthDentalStatementQueryAPIRequest) SetOrderId(_orderId string) error {
	r._orderId = _orderId
	r.Set("order_id", _orderId)
	return nil
}

// GetOrderId OrderId Getter
func (r AlibabaAlihealthDentalStatementQueryAPIRequest) GetOrderId() string {
	return r._orderId
}

// SetStatementTime is StatementTime Setter
// 结算周期，单位月
func (r *AlibabaAlihealthDentalStatementQueryAPIRequest) SetStatementTime(_statementTime string) error {
	r._statementTime = _statementTime
	r.Set("statement_time", _statementTime)
	return nil
}

// GetStatementTime StatementTime Getter
func (r AlibabaAlihealthDentalStatementQueryAPIRequest) GetStatementTime() string {
	return r._statementTime
}
