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
type AlibabaSecurityJaqRpCloudEventAPIRequest struct {
    model.Params
    // 认证token
    _verifyToken   string
    // 事件编码
    _eventCode   string
    // 事件信息
    _eventData   string
}

// 初始化AlibabaSecurityJaqRpCloudEventAPIRequest对象
func NewAlibabaSecurityJaqRpCloudEventRequest() *AlibabaSecurityJaqRpCloudEventAPIRequest{
    return &AlibabaSecurityJaqRpCloudEventAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaSecurityJaqRpCloudEventAPIRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.cloud.event"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaSecurityJaqRpCloudEventAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// VerifyToken Setter
// 认证token
func (r *AlibabaSecurityJaqRpCloudEventAPIRequest) SetVerifyToken(_verifyToken string) error {
    r._verifyToken = _verifyToken
    r.Set("verify_token", _verifyToken)
    return nil
}

// VerifyToken Getter
func (r AlibabaSecurityJaqRpCloudEventAPIRequest) GetVerifyToken() string {
    return r._verifyToken
}
// EventCode Setter
// 事件编码
func (r *AlibabaSecurityJaqRpCloudEventAPIRequest) SetEventCode(_eventCode string) error {
    r._eventCode = _eventCode
    r.Set("event_code", _eventCode)
    return nil
}

// EventCode Getter
func (r AlibabaSecurityJaqRpCloudEventAPIRequest) GetEventCode() string {
    return r._eventCode
}
// EventData Setter
// 事件信息
func (r *AlibabaSecurityJaqRpCloudEventAPIRequest) SetEventData(_eventData string) error {
    r._eventData = _eventData
    r.Set("event_data", _eventData)
    return nil
}

// EventData Getter
func (r AlibabaSecurityJaqRpCloudEventAPIRequest) GetEventData() string {
    return r._eventData
}
