package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
流量发放(用户id) APIRequest
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

func NewAlibabaAliqinFlowPublishRequest() *AlibabaAliqinFlowPublishRequest{
    return &AlibabaAliqinFlowPublishRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinFlowPublishRequest) GetApiMethodName() string {
    return "alibaba.aliqin.flow.publish"
}

func (r AlibabaAliqinFlowPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinFlowPublishRequest) SetUserId(userId string) error {
    r.userId = userId
    r.Set("user_id", userId)
    return nil
}

func (r AlibabaAliqinFlowPublishRequest) GetUserId() string {
    return r.userId
}

func (r *AlibabaAliqinFlowPublishRequest) SetFlow(flow string) error {
    r.flow = flow
    r.Set("flow", flow)
    return nil
}

func (r AlibabaAliqinFlowPublishRequest) GetFlow() string {
    return r.flow
}

func (r *AlibabaAliqinFlowPublishRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

func (r AlibabaAliqinFlowPublishRequest) GetReason() string {
    return r.reason
}

func (r *AlibabaAliqinFlowPublishRequest) SetSerial(serial string) error {
    r.serial = serial
    r.Set("serial", serial)
    return nil
}

func (r AlibabaAliqinFlowPublishRequest) GetSerial() string {
    return r.serial
}

func (r *AlibabaAliqinFlowPublishRequest) SetAlways(always string) error {
    r.always = always
    r.Set("always", always)
    return nil
}

func (r AlibabaAliqinFlowPublishRequest) GetAlways() string {
    return r.always
}

