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
    _buyerNick   string
    // 真实用户名称，如果填写这个字段，buyer_nick失效
    _realNick   string
    // 唯一流水号，字母+数字组合
    _serial   string
    // 流量
    _flow   string
    // 购物送
    _reason   string
    // 设置true为始终发送成功
    _always   string
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
func (r *AlibabaAliqinFlowWalletSendFlowRequest) SetBuyerNick(_buyerNick string) error {
    r._buyerNick = _buyerNick
    r.Set("buyer_nick", _buyerNick)
    return nil
}

// BuyerNick Getter
func (r AlibabaAliqinFlowWalletSendFlowRequest) GetBuyerNick() string {
    return r._buyerNick
}
// RealNick Setter
// 真实用户名称，如果填写这个字段，buyer_nick失效
func (r *AlibabaAliqinFlowWalletSendFlowRequest) SetRealNick(_realNick string) error {
    r._realNick = _realNick
    r.Set("real_nick", _realNick)
    return nil
}

// RealNick Getter
func (r AlibabaAliqinFlowWalletSendFlowRequest) GetRealNick() string {
    return r._realNick
}
// Serial Setter
// 唯一流水号，字母+数字组合
func (r *AlibabaAliqinFlowWalletSendFlowRequest) SetSerial(_serial string) error {
    r._serial = _serial
    r.Set("serial", _serial)
    return nil
}

// Serial Getter
func (r AlibabaAliqinFlowWalletSendFlowRequest) GetSerial() string {
    return r._serial
}
// Flow Setter
// 流量
func (r *AlibabaAliqinFlowWalletSendFlowRequest) SetFlow(_flow string) error {
    r._flow = _flow
    r.Set("flow", _flow)
    return nil
}

// Flow Getter
func (r AlibabaAliqinFlowWalletSendFlowRequest) GetFlow() string {
    return r._flow
}
// Reason Setter
// 购物送
func (r *AlibabaAliqinFlowWalletSendFlowRequest) SetReason(_reason string) error {
    r._reason = _reason
    r.Set("reason", _reason)
    return nil
}

// Reason Getter
func (r AlibabaAliqinFlowWalletSendFlowRequest) GetReason() string {
    return r._reason
}
// Always Setter
// 设置true为始终发送成功
func (r *AlibabaAliqinFlowWalletSendFlowRequest) SetAlways(_always string) error {
    r._always = _always
    r.Set("always", _always)
    return nil
}

// Always Getter
func (r AlibabaAliqinFlowWalletSendFlowRequest) GetAlways() string {
    return r._always
}
