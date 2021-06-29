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
    userId   string
    // 流量
    flow   string
    // 原因
    reason   string
    // 唯一流水号（字母+数字）
    serial   string
    // 设置true为始终发送成功
    always   string
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
func (r *AlibabaAliqinFlowPublishRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

// UserId Getter
func (r AlibabaAliqinFlowPublishRequest) GetUserId() string {
    return r.userId
}
// Flow Setter
// 流量
func (r *AlibabaAliqinFlowPublishRequest) SetFlow(flow string) error {
    r.flow = flow
    r.Set("flow", flow)
    return nil
}

// Flow Getter
func (r AlibabaAliqinFlowPublishRequest) GetFlow() string {
    return r.flow
}
// Reason Setter
// 原因
func (r *AlibabaAliqinFlowPublishRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

// Reason Getter
func (r AlibabaAliqinFlowPublishRequest) GetReason() string {
    return r.reason
}
// Serial Setter
// 唯一流水号（字母+数字）
func (r *AlibabaAliqinFlowPublishRequest) SetSerial(serial string) error {
    r.serial = serial
    r.Set("serial", serial)
    return nil
}

// Serial Getter
func (r AlibabaAliqinFlowPublishRequest) GetSerial() string {
    return r.serial
}
// Always Setter
// 设置true为始终发送成功
func (r *AlibabaAliqinFlowPublishRequest) SetAlways(always string) error {
    r.always = always
    r.Set("always", always)
    return nil
}

// Always Getter
func (r AlibabaAliqinFlowPublishRequest) GetAlways() string {
    return r.always
}
