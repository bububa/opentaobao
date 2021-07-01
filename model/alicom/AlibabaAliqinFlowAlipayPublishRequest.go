package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
流量钱包流量发放-面向支付宝用户 API请求
alibaba.aliqin.flow.alipay.publish

用户淘宝流量钱包商家给支付宝用户发放流量
*/
type AlibabaAliqinFlowAlipayPublishAPIRequest struct {
    model.Params
    // 用户的支付宝ID
    _alipayId   string
    // 外部流水号，用来做幂等校验
    _serial   string
    // 发放的流量数，单位MB
    _flow   string
    // 发放原因
    _reason   string
}

// 初始化AlibabaAliqinFlowAlipayPublishAPIRequest对象
func NewAlibabaAliqinFlowAlipayPublishRequest() *AlibabaAliqinFlowAlipayPublishAPIRequest{
    return &AlibabaAliqinFlowAlipayPublishAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFlowAlipayPublishAPIRequest) GetApiMethodName() string {
    return "alibaba.aliqin.flow.alipay.publish"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFlowAlipayPublishAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AlipayId Setter
// 用户的支付宝ID
func (r *AlibabaAliqinFlowAlipayPublishAPIRequest) SetAlipayId(_alipayId string) error {
    r._alipayId = _alipayId
    r.Set("alipay_id", _alipayId)
    return nil
}

// AlipayId Getter
func (r AlibabaAliqinFlowAlipayPublishAPIRequest) GetAlipayId() string {
    return r._alipayId
}
// Serial Setter
// 外部流水号，用来做幂等校验
func (r *AlibabaAliqinFlowAlipayPublishAPIRequest) SetSerial(_serial string) error {
    r._serial = _serial
    r.Set("serial", _serial)
    return nil
}

// Serial Getter
func (r AlibabaAliqinFlowAlipayPublishAPIRequest) GetSerial() string {
    return r._serial
}
// Flow Setter
// 发放的流量数，单位MB
func (r *AlibabaAliqinFlowAlipayPublishAPIRequest) SetFlow(_flow string) error {
    r._flow = _flow
    r.Set("flow", _flow)
    return nil
}

// Flow Getter
func (r AlibabaAliqinFlowAlipayPublishAPIRequest) GetFlow() string {
    return r._flow
}
// Reason Setter
// 发放原因
func (r *AlibabaAliqinFlowAlipayPublishAPIRequest) SetReason(_reason string) error {
    r._reason = _reason
    r.Set("reason", _reason)
    return nil
}

// Reason Getter
func (r AlibabaAliqinFlowAlipayPublishAPIRequest) GetReason() string {
    return r._reason
}
