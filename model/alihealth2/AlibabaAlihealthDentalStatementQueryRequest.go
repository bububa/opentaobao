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
type AlibabaAlihealthDentalStatementQueryRequest struct {
    model.Params
    // 订单ID
    orderId   string
    // 结算周期，单位月
    statementTime   string
}

// 初始化AlibabaAlihealthDentalStatementQueryRequest对象
func NewAlibabaAlihealthDentalStatementQueryRequest() *AlibabaAlihealthDentalStatementQueryRequest{
    return &AlibabaAlihealthDentalStatementQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDentalStatementQueryRequest) GetApiMethodName() string {
    return "alibaba.alihealth.dental.statement.query"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDentalStatementQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// OrderId Setter
// 订单ID
func (r *AlibabaAlihealthDentalStatementQueryRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

// OrderId Getter
func (r AlibabaAlihealthDentalStatementQueryRequest) GetOrderId() string {
    return r.orderId
}
// StatementTime Setter
// 结算周期，单位月
func (r *AlibabaAlihealthDentalStatementQueryRequest) SetStatementTime(statementTime string) error {
    r.statementTime = statementTime
    r.Set("statement_time", statementTime)
    return nil
}

// StatementTime Getter
func (r AlibabaAlihealthDentalStatementQueryRequest) GetStatementTime() string {
    return r.statementTime
}
