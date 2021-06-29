package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
流量发放 API请求
alibaba.aliqin.flow.wallet.send.flow

阿里通信流量下发功能，允许用户补发
*/
type AlibabaAliqinFlowWalletSendFlowRequest struct {
    model.Params
    // 混淆用户名
    buyerNick   string
    // 真实用户名称，如果填写这个字段，buyer_nick失效
    realNick   string
    // 唯一流水号，字母+数字组合
    serial   string
    // 流量
    flow   string
    // 购物送
    reason   string
    // 设置true为始终发送成功
    always   string
}

// 初始化AlibabaAliqinFlowWalletSendFlowRequest对象
func NewAlibabaAliqinFlowWalletSendFlowRequest() *AlibabaAliqinFlowWalletSendFlowRequest{
    return &AlibabaAliqinFlowWalletSendFlowRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFlowWalletSendFlowRequest) GetApiMethodName() string {
    return "alibaba.aliqin.flow.wallet.send.flow"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFlowWalletSendFlowRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BuyerNick Setter
// 混淆用户名
func (r *AlibabaAliqinFlowWalletSendFlowRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

// BuyerNick Getter
func (r AlibabaAliqinFlowWalletSendFlowRequest) GetBuyerNick() string {
    return r.buyerNick
}
// RealNick Setter
// 真实用户名称，如果填写这个字段，buyer_nick失效
func (r *AlibabaAliqinFlowWalletSendFlowRequest) SetRealNick(realNick string) error {
    r.realNick = realNick
    r.Set("real_nick", realNick)
    return nil
}

// RealNick Getter
func (r AlibabaAliqinFlowWalletSendFlowRequest) GetRealNick() string {
    return r.realNick
}
// Serial Setter
// 唯一流水号，字母+数字组合
func (r *AlibabaAliqinFlowWalletSendFlowRequest) SetSerial(serial string) error {
    r.serial = serial
    r.Set("serial", serial)
    return nil
}

// Serial Getter
func (r AlibabaAliqinFlowWalletSendFlowRequest) GetSerial() string {
    return r.serial
}
// Flow Setter
// 流量
func (r *AlibabaAliqinFlowWalletSendFlowRequest) SetFlow(flow string) error {
    r.flow = flow
    r.Set("flow", flow)
    return nil
}

// Flow Getter
func (r AlibabaAliqinFlowWalletSendFlowRequest) GetFlow() string {
    return r.flow
}
// Reason Setter
// 购物送
func (r *AlibabaAliqinFlowWalletSendFlowRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

// Reason Getter
func (r AlibabaAliqinFlowWalletSendFlowRequest) GetReason() string {
    return r.reason
}
// Always Setter
// 设置true为始终发送成功
func (r *AlibabaAliqinFlowWalletSendFlowRequest) SetAlways(always string) error {
    r.always = always
    r.Set("always", always)
    return nil
}

// Always Getter
func (r AlibabaAliqinFlowWalletSendFlowRequest) GetAlways() string {
    return r.always
}
