package alihealth2

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
ISV查询对账单 APIRequest
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

func NewAlibabaAlihealthDentalStatementQueryRequest() *AlibabaAlihealthDentalStatementQueryRequest{
    return &AlibabaAlihealthDentalStatementQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDentalStatementQueryRequest) GetApiMethodName() string {
    return "alibaba.alihealth.dental.statement.query"
}

func (r AlibabaAlihealthDentalStatementQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDentalStatementQueryRequest) SetOrderId(orderId string) error {
    r.orderId = orderId
    r.Set("order_id", orderId)
    return nil
}

func (r AlibabaAlihealthDentalStatementQueryRequest) GetOrderId() string {
    return r.orderId
}

func (r *AlibabaAlihealthDentalStatementQueryRequest) SetStatementTime(statementTime string) error {
    r.statementTime = statementTime
    r.Set("statement_time", statementTime)
    return nil
}

func (r AlibabaAlihealthDentalStatementQueryRequest) GetStatementTime() string {
    return r.statementTime
}

