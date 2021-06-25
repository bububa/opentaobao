package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
饿了么对账单查询，带订单明细 APIRequest
alibaba.wdk.eleme.bill.detail.get

查询饿了么对账单信息，带订单明细
*/
type AlibabaWdkElemeBillDetailGetRequest struct {
    model.Params

    // 对账单查询参数
    eleBillRequest   *EleBillRequest 

}

func NewAlibabaWdkElemeBillDetailGetRequest() *AlibabaWdkElemeBillDetailGetRequest{
    return &AlibabaWdkElemeBillDetailGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkElemeBillDetailGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.eleme.bill.detail.get"
}

func (r AlibabaWdkElemeBillDetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkElemeBillDetailGetRequest) SetEleBillRequest(eleBillRequest *EleBillRequest) error {
    r.eleBillRequest = eleBillRequest
    r.Set("ele_bill_request", eleBillRequest)
    return nil
}

func (r AlibabaWdkElemeBillDetailGetRequest) GetEleBillRequest() *EleBillRequest {
    return r.eleBillRequest
}

