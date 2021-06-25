package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
转呼控制接口 APIRequest
alibaba.aliqin.axb.vendor.call.control

转呼控制接口，用于查询小号绑定关系，控制呼叫转接目标
*/
type AlibabaAliqinAxbVendorCallControlRequest struct {
    model.Params

    // 转接控制接口request对象
    startCallRequest   *StartCallRequest 

}

func NewAlibabaAliqinAxbVendorCallControlRequest() *AlibabaAliqinAxbVendorCallControlRequest{
    return &AlibabaAliqinAxbVendorCallControlRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinAxbVendorCallControlRequest) GetApiMethodName() string {
    return "alibaba.aliqin.axb.vendor.call.control"
}

func (r AlibabaAliqinAxbVendorCallControlRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinAxbVendorCallControlRequest) SetStartCallRequest(startCallRequest *StartCallRequest) error {
    r.startCallRequest = startCallRequest
    r.Set("start_call_request", startCallRequest)
    return nil
}

func (r AlibabaAliqinAxbVendorCallControlRequest) GetStartCallRequest() *StartCallRequest {
    return r.startCallRequest
}

