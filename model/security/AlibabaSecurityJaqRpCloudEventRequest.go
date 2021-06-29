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
    _verifyToken   string
    // 事件编码
    _eventCode   string
    // 事件信息
    _eventData   string
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
func (r *AlibabaSecurityJaqRpCloudEventRequest) SetVerifyToken(_verifyToken string) error {
    r._verifyToken = _verifyToken
    r.Set("verify_token", _verifyToken)
    return nil
}

// VerifyToken Getter
func (r AlibabaSecurityJaqRpCloudEventRequest) GetVerifyToken() string {
    return r._verifyToken
}
// EventCode Setter
// 事件编码
func (r *AlibabaSecurityJaqRpCloudEventRequest) SetEventCode(_eventCode string) error {
    r._eventCode = _eventCode
    r.Set("event_code", _eventCode)
    return nil
}

// EventCode Getter
func (r AlibabaSecurityJaqRpCloudEventRequest) GetEventCode() string {
    return r._eventCode
}
// EventData Setter
// 事件信息
func (r *AlibabaSecurityJaqRpCloudEventRequest) SetEventData(_eventData string) error {
    r._eventData = _eventData
    r.Set("event_data", _eventData)
    return nil
}

// EventData Getter
func (r AlibabaSecurityJaqRpCloudEventRequest) GetEventData() string {
    return r._eventData
}
