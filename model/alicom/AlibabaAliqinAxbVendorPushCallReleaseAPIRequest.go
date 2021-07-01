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
type AlibabaAliqinAxbVendorPushCallReleaseAPIRequest struct {
    model.Params
    // end_call_request
    _endCallRequest   *EndCallRequest
}

// 初始化AlibabaAliqinAxbVendorPushCallReleaseAPIRequest对象
func NewAlibabaAliqinAxbVendorPushCallReleaseRequest() *AlibabaAliqinAxbVendorPushCallReleaseAPIRequest{
    return &AlibabaAliqinAxbVendorPushCallReleaseAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinAxbVendorPushCallReleaseAPIRequest) GetApiMethodName() string {
    return "alibaba.aliqin.axb.vendor.push.call.release"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinAxbVendorPushCallReleaseAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// EndCallRequest Setter
// end_call_request
func (r *AlibabaAliqinAxbVendorPushCallReleaseAPIRequest) SetEndCallRequest(_endCallRequest *EndCallRequest) error {
    r._endCallRequest = _endCallRequest
    r.Set("end_call_request", _endCallRequest)
    return nil
}

// EndCallRequest Getter
func (r AlibabaAliqinAxbVendorPushCallReleaseAPIRequest) GetEndCallRequest() *EndCallRequest {
    return r._endCallRequest
}
