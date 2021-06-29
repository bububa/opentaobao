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
    _endCallRequest   *EndCallRequest
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
func (r *AlibabaAliqinAxbVendorPushCallReleaseRequest) SetEndCallRequest(_endCallRequest *EndCallRequest) error {
    r._endCallRequest = _endCallRequest
    r.Set("end_call_request", _endCallRequest)
    return nil
}

// EndCallRequest Getter
func (r AlibabaAliqinAxbVendorPushCallReleaseRequest) GetEndCallRequest() *EndCallRequest {
    return r._endCallRequest
}
