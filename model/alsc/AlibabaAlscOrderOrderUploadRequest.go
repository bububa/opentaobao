package alsc

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
订单回流 API请求
alibaba.alsc.order.order.upload

第三方订单回流
*/
type AlibabaAlscOrderOrderUploadRequest struct {
    model.Params
    // 订单回流参数
    paramBackflowRequest   *BackflowRequest
}

// 初始化AlibabaAlscOrderOrderUploadRequest对象
func NewAlibabaAlscOrderOrderUploadRequest() *AlibabaAlscOrderOrderUploadRequest{
    return &AlibabaAlscOrderOrderUploadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlscOrderOrderUploadRequest) GetApiMethodName() string {
    return "alibaba.alsc.order.order.upload"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlscOrderOrderUploadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// ParamBackflowRequest Setter
// 订单回流参数
func (r *AlibabaAlscOrderOrderUploadRequest) SetParamBackflowRequest(paramBackflowRequest *BackflowRequest) error {
    r.paramBackflowRequest = paramBackflowRequest
    r.Set("param_backflow_request", paramBackflowRequest)
    return nil
}

// ParamBackflowRequest Getter
func (r AlibabaAlscOrderOrderUploadRequest) GetParamBackflowRequest() *BackflowRequest {
    return r.paramBackflowRequest
}
