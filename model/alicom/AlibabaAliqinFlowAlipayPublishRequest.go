package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
流量钱包流量发放-面向支付宝用户 APIRequest
alibaba.aliqin.flow.alipay.publish

用户淘宝流量钱包商家给支付宝用户发放流量
*/
type AlibabaAliqinFlowAlipayPublishRequest struct {
    model.Params

    // 用户的支付宝ID
    alipayId   string 

    // 外部流水号，用来做幂等校验
    serial   string 

    // 发放的流量数，单位MB
    flow   string 

    // 发放原因
    reason   string 

}

func NewAlibabaAliqinFlowAlipayPublishRequest() *AlibabaAliqinFlowAlipayPublishRequest{
    return &AlibabaAliqinFlowAlipayPublishRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinFlowAlipayPublishRequest) GetApiMethodName() string {
    return "alibaba.aliqin.flow.alipay.publish"
}

func (r AlibabaAliqinFlowAlipayPublishRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinFlowAlipayPublishRequest) SetAlipayId(alipayId string) error {
    r.alipayId = alipayId
    r.Set("alipay_id", alipayId)
    return nil
}

func (r AlibabaAliqinFlowAlipayPublishRequest) GetAlipayId() string {
    return r.alipayId
}

func (r *AlibabaAliqinFlowAlipayPublishRequest) SetSerial(serial string) error {
    r.serial = serial
    r.Set("serial", serial)
    return nil
}

func (r AlibabaAliqinFlowAlipayPublishRequest) GetSerial() string {
    return r.serial
}

func (r *AlibabaAliqinFlowAlipayPublishRequest) SetFlow(flow string) error {
    r.flow = flow
    r.Set("flow", flow)
    return nil
}

func (r AlibabaAliqinFlowAlipayPublishRequest) GetFlow() string {
    return r.flow
}

func (r *AlibabaAliqinFlowAlipayPublishRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

func (r AlibabaAliqinFlowAlipayPublishRequest) GetReason() string {
    return r.reason
}

