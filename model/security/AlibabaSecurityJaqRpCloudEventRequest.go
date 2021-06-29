package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
事件上报 API请求
alibaba.security.jaq.rp.cloud.event

事件上报接口
*/
type AlibabaSecurityJaqRpCloudEventRequest struct {
    model.Params
    // 认证token
    verifyToken   string
    // 事件编码
    eventCode   string
    // 事件信息
    eventData   string
}

// 初始化AlibabaSecurityJaqRpCloudEventRequest对象
func NewAlibabaSecurityJaqRpCloudEventRequest() *AlibabaSecurityJaqRpCloudEventRequest{
    return &AlibabaSecurityJaqRpCloudEventRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpCloudEventRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.cloud.event"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpCloudEventRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VerifyToken Setter
// 认证token
func (r *AlibabaSecurityJaqRpCloudEventRequest) SetVerifyToken(verifyToken string) error {
    r.verifyToken = verifyToken
    r.Set("verify_token", verifyToken)
    return nil
}

// VerifyToken Getter
func (r AlibabaSecurityJaqRpCloudEventRequest) GetVerifyToken() string {
    return r.verifyToken
}
// EventCode Setter
// 事件编码
func (r *AlibabaSecurityJaqRpCloudEventRequest) SetEventCode(eventCode string) error {
    r.eventCode = eventCode
    r.Set("event_code", eventCode)
    return nil
}

// EventCode Getter
func (r AlibabaSecurityJaqRpCloudEventRequest) GetEventCode() string {
    return r.eventCode
}
// EventData Setter
// 事件信息
func (r *AlibabaSecurityJaqRpCloudEventRequest) SetEventData(eventData string) error {
    r.eventData = eventData
    r.Set("event_data", eventData)
    return nil
}

// EventData Getter
func (r AlibabaSecurityJaqRpCloudEventRequest) GetEventData() string {
    return r.eventData
}
