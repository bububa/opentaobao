package icburfq

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
我的权益 APIRequest
alibaba.icbu.rfq.myequity

查询供应商权益接口
*/
type AlibabaIcbuRfqMyequityRequest struct {
    model.Params

}

func NewAlibabaIcbuRfqMyequityRequest() *AlibabaIcbuRfqMyequityRequest{
    return &AlibabaIcbuRfqMyequityRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaIcbuRfqMyequityRequest) GetApiMethodName() string {
    return "alibaba.icbu.rfq.myequity"
}

func (r AlibabaIcbuRfqMyequityRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


