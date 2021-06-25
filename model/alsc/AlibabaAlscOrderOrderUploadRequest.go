package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单回流 APIRequest
alibaba.alsc.order.order.upload

第三方订单回流
*/
type AlibabaAlscOrderOrderUploadRequest struct {
    model.Params

    // 订单回流参数
    paramBackflowRequest   *BackflowRequest 

}

func NewAlibabaAlscOrderOrderUploadRequest() *AlibabaAlscOrderOrderUploadRequest{
    return &AlibabaAlscOrderOrderUploadRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlscOrderOrderUploadRequest) GetApiMethodName() string {
    return "alibaba.alsc.order.order.upload"
}

func (r AlibabaAlscOrderOrderUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlscOrderOrderUploadRequest) SetParamBackflowRequest(paramBackflowRequest *BackflowRequest) error {
    r.paramBackflowRequest = paramBackflowRequest
    r.Set("param_backflow_request", paramBackflowRequest)
    return nil
}

func (r AlibabaAlscOrderOrderUploadRequest) GetParamBackflowRequest() *BackflowRequest {
    return r.paramBackflowRequest
}

