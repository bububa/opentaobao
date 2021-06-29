package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
供应商推送通话结束事件 API请求
alibaba.aliqin.axb.vendor.push.call.release

通话结束挂断的时候，供应商推送通话结束事件给阿里侧
*/
type AlibabaAliqinAxbVendorPushCallReleaseRequest struct {
    model.Params
    // end_call_request
    endCallRequest   *EndCallRequest
}

// 初始化AlibabaAliqinAxbVendorPushCallReleaseRequest对象
func NewAlibabaAliqinAxbVendorPushCallReleaseRequest() *AlibabaAliqinAxbVendorPushCallReleaseRequest{
    return &AlibabaAliqinAxbVendorPushCallReleaseRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinAxbVendorPushCallReleaseRequest) GetApiMethodName() string {
    return "alibaba.aliqin.axb.vendor.push.call.release"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinAxbVendorPushCallReleaseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EndCallRequest Setter
// end_call_request
func (r *AlibabaAliqinAxbVendorPushCallReleaseRequest) SetEndCallRequest(endCallRequest *EndCallRequest) error {
    r.endCallRequest = endCallRequest
    r.Set("end_call_request", endCallRequest)
    return nil
}

// EndCallRequest Getter
func (r AlibabaAliqinAxbVendorPushCallReleaseRequest) GetEndCallRequest() *EndCallRequest {
    return r.endCallRequest
}
