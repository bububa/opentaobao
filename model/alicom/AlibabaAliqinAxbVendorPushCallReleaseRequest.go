package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商推送通话结束事件 APIRequest
alibaba.aliqin.axb.vendor.push.call.release

通话结束挂断的时候，供应商推送通话结束事件给阿里侧
*/
type AlibabaAliqinAxbVendorPushCallReleaseRequest struct {
    model.Params

    // end_call_request
    endCallRequest   *EndCallRequest 

}

func NewAlibabaAliqinAxbVendorPushCallReleaseRequest() *AlibabaAliqinAxbVendorPushCallReleaseRequest{
    return &AlibabaAliqinAxbVendorPushCallReleaseRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinAxbVendorPushCallReleaseRequest) GetApiMethodName() string {
    return "alibaba.aliqin.axb.vendor.push.call.release"
}

func (r AlibabaAliqinAxbVendorPushCallReleaseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinAxbVendorPushCallReleaseRequest) SetEndCallRequest(endCallRequest *EndCallRequest) error {
    r.endCallRequest = endCallRequest
    r.Set("end_call_request", endCallRequest)
    return nil
}

func (r AlibabaAliqinAxbVendorPushCallReleaseRequest) GetEndCallRequest() *EndCallRequest {
    return r.endCallRequest
}

