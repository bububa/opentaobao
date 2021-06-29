package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
转呼控制接口 API请求
alibaba.aliqin.axb.vendor.call.control

转呼控制接口，用于查询小号绑定关系，控制呼叫转接目标
*/
type AlibabaAliqinAxbVendorCallControlRequest struct {
    model.Params
    // 转接控制接口request对象
    _startCallRequest   *StartCallRequest
}

// 初始化AlibabaAliqinAxbVendorCallControlRequest对象
func NewAlibabaAliqinAxbVendorCallControlRequest() *AlibabaAliqinAxbVendorCallControlRequest{
    return &AlibabaAliqinAxbVendorCallControlRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinAxbVendorCallControlRequest) GetApiMethodName() string {
    return "alibaba.aliqin.axb.vendor.call.control"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinAxbVendorCallControlRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// StartCallRequest Setter
// 转接控制接口request对象
func (r *AlibabaAliqinAxbVendorCallControlRequest) SetStartCallRequest(_startCallRequest *StartCallRequest) error {
    r._startCallRequest = _startCallRequest
    r.Set("start_call_request", _startCallRequest)
    return nil
}

// StartCallRequest Getter
func (r AlibabaAliqinAxbVendorCallControlRequest) GetStartCallRequest() *StartCallRequest {
    return r._startCallRequest
}
