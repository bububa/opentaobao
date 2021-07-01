package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV查询对账单 API请求
alibaba.alihealth.dental.statement.query

ISV查询对账单
*/
type AlibabaAlihealthDentalStatementQueryAPIRequest struct {
    model.Params
    // 订单ID
    _orderId   string
    // 结算周期，单位月
    _statementTime   string
}

// 初始化AlibabaAlihealthDentalStatementQueryAPIRequest对象
func NewAlibabaAlihealthDentalStatementQueryRequest() *AlibabaAlihealthDentalStatementQueryAPIRequest{
    return &AlibabaAlihealthDentalStatementQueryAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDentalStatementQueryAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.dental.statement.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDentalStatementQueryAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单ID
func (r *AlibabaAlihealthDentalStatementQueryAPIRequest) SetOrderId(_orderId string) error {
    r._orderId = _orderId
    r.Set("order_id", _orderId)
    return nil
}

// OrderId Getter
func (r AlibabaAlihealthDentalStatementQueryAPIRequest) GetOrderId() string {
    return r._orderId
}
// StatementTime Setter
// 结算周期，单位月
func (r *AlibabaAlihealthDentalStatementQueryAPIRequest) SetStatementTime(_statementTime string) error {
    r._statementTime = _statementTime
    r.Set("statement_time", _statementTime)
    return nil
}

// StatementTime Getter
func (r AlibabaAlihealthDentalStatementQueryAPIRequest) GetStatementTime() string {
    return r._statementTime
}
