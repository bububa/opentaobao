package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
流量发放(用户id) API请求
alibaba.aliqin.flow.publish

阿里通信流量下发功能，允许用户补发
*/
type AlibabaAliqinFlowPublishRequest struct {
    model.Params
    // 用户id
    _userId   string
    // 流量
    _flow   string
    // 原因
    _reason   string
    // 唯一流水号（字母+数字）
    _serial   string
    // 设置true为始终发送成功
    _always   string
}

// 初始化AlibabaAliqinFlowPublishRequest对象
func NewAlibabaAliqinFlowPublishRequest() *AlibabaAliqinFlowPublishRequest{
    return &AlibabaAliqinFlowPublishRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFlowPublishRequest) GetApiMethodName() string {
    return "alibaba.aliqin.flow.publish"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFlowPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// UserId Setter
// 用户id
func (r *AlibabaAliqinFlowPublishRequest) SetUserId(_userId string) error {
    r._userId = _userId
    r.Set("user_id", _userId)
    return nil
}

// UserId Getter
func (r AlibabaAliqinFlowPublishRequest) GetUserId() string {
    return r._userId
}
// Flow Setter
// 流量
func (r *AlibabaAliqinFlowPublishRequest) SetFlow(_flow string) error {
    r._flow = _flow
    r.Set("flow", _flow)
    return nil
}

// Flow Getter
func (r AlibabaAliqinFlowPublishRequest) GetFlow() string {
    return r._flow
}
// Reason Setter
// 原因
func (r *AlibabaAliqinFlowPublishRequest) SetReason(_reason string) error {
    r._reason = _reason
    r.Set("reason", _reason)
    return nil
}

// Reason Getter
func (r AlibabaAliqinFlowPublishRequest) GetReason() string {
    return r._reason
}
// Serial Setter
// 唯一流水号（字母+数字）
func (r *AlibabaAliqinFlowPublishRequest) SetSerial(_serial string) error {
    r._serial = _serial
    r.Set("serial", _serial)
    return nil
}

// Serial Getter
func (r AlibabaAliqinFlowPublishRequest) GetSerial() string {
    return r._serial
}
// Always Setter
// 设置true为始终发送成功
func (r *AlibabaAliqinFlowPublishRequest) SetAlways(_always string) error {
    r._always = _always
    r.Set("always", _always)
    return nil
}

// Always Getter
func (r AlibabaAliqinFlowPublishRequest) GetAlways() string {
    return r._always
}
