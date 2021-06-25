package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
资金合规商家账单 APIRequest
alibaba.wdk.order.finance.bill.query

拉取资金合规商家账单
*/
type AlibabaWdkOrderFinanceBillQueryRequest struct {
    model.Params

    // 入参
    billQueryRequest   *WdkOpenOrderFinanceBillQueryRequest 

}

func NewAlibabaWdkOrderFinanceBillQueryRequest() *AlibabaWdkOrderFinanceBillQueryRequest{
    return &AlibabaWdkOrderFinanceBillQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkOrderFinanceBillQueryRequest) GetApiMethodName() string {
    return "alibaba.wdk.order.finance.bill.query"
}

func (r AlibabaWdkOrderFinanceBillQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkOrderFinanceBillQueryRequest) SetBillQueryRequest(billQueryRequest *WdkOpenOrderFinanceBillQueryRequest) error {
    r.billQueryRequest = billQueryRequest
    r.Set("bill_query_request", billQueryRequest)
    return nil
}

func (r AlibabaWdkOrderFinanceBillQueryRequest) GetBillQueryRequest() *WdkOpenOrderFinanceBillQueryRequest {
    return r.billQueryRequest
}

