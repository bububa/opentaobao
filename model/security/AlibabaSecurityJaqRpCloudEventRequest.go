package security

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
事件上报 APIRequest
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

func NewAlibabaSecurityJaqRpCloudEventRequest() *AlibabaSecurityJaqRpCloudEventRequest{
    return &AlibabaSecurityJaqRpCloudEventRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaSecurityJaqRpCloudEventRequest) GetApiMethodName() string {
    return "alibaba.security.jaq.rp.cloud.event"
}

func (r AlibabaSecurityJaqRpCloudEventRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaSecurityJaqRpCloudEventRequest) SetVerifyToken(verifyToken string) error {
    r.verifyToken = verifyToken
    r.Set("verify_token", verifyToken)
    return nil
}

func (r AlibabaSecurityJaqRpCloudEventRequest) GetVerifyToken() string {
    return r.verifyToken
}

func (r *AlibabaSecurityJaqRpCloudEventRequest) SetEventCode(eventCode string) error {
    r.eventCode = eventCode
    r.Set("event_code", eventCode)
    return nil
}

func (r AlibabaSecurityJaqRpCloudEventRequest) GetEventCode() string {
    return r.eventCode
}

func (r *AlibabaSecurityJaqRpCloudEventRequest) SetEventData(eventData string) error {
    r.eventData = eventData
    r.Set("event_data", eventData)
    return nil
}

func (r AlibabaSecurityJaqRpCloudEventRequest) GetEventData() string {
    return r.eventData
}

