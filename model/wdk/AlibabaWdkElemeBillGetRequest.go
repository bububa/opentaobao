package wdk

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
饿了么日维度对账单查询 APIRequest
alibaba.wdk.eleme.bill.get

查询饿了么日维度对账单信息
*/
type AlibabaWdkElemeBillGetRequest struct {
    model.Params

    // 对账单查询参数
    eleBillRequest   *EleBillRequest 

}

func NewAlibabaWdkElemeBillGetRequest() *AlibabaWdkElemeBillGetRequest{
    return &AlibabaWdkElemeBillGetRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaWdkElemeBillGetRequest) GetApiMethodName() string {
    return "alibaba.wdk.eleme.bill.get"
}

func (r AlibabaWdkElemeBillGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaWdkElemeBillGetRequest) SetEleBillRequest(eleBillRequest *EleBillRequest) error {
    r.eleBillRequest = eleBillRequest
    r.Set("ele_bill_request", eleBillRequest)
    return nil
}

func (r AlibabaWdkElemeBillGetRequest) GetEleBillRequest() *EleBillRequest {
    return r.eleBillRequest
}

