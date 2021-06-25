package alicom

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
流量发放 APIRequest
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

func NewAlibabaAliqinFlowWalletSendFlowRequest() *AlibabaAliqinFlowWalletSendFlowRequest{
    return &AlibabaAliqinFlowWalletSendFlowRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAliqinFlowWalletSendFlowRequest) GetApiMethodName() string {
    return "alibaba.aliqin.flow.wallet.send.flow"
}

func (r AlibabaAliqinFlowWalletSendFlowRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAliqinFlowWalletSendFlowRequest) SetBuyerNick(buyerNick string) error {
    r.buyerNick = buyerNick
    r.Set("buyer_nick", buyerNick)
    return nil
}

func (r AlibabaAliqinFlowWalletSendFlowRequest) GetBuyerNick() string {
    return r.buyerNick
}

func (r *AlibabaAliqinFlowWalletSendFlowRequest) SetRealNick(realNick string) error {
    r.realNick = realNick
    r.Set("real_nick", realNick)
    return nil
}

func (r AlibabaAliqinFlowWalletSendFlowRequest) GetRealNick() string {
    return r.realNick
}

func (r *AlibabaAliqinFlowWalletSendFlowRequest) SetSerial(serial string) error {
    r.serial = serial
    r.Set("serial", serial)
    return nil
}

func (r AlibabaAliqinFlowWalletSendFlowRequest) GetSerial() string {
    return r.serial
}

func (r *AlibabaAliqinFlowWalletSendFlowRequest) SetFlow(flow string) error {
    r.flow = flow
    r.Set("flow", flow)
    return nil
}

func (r AlibabaAliqinFlowWalletSendFlowRequest) GetFlow() string {
    return r.flow
}

func (r *AlibabaAliqinFlowWalletSendFlowRequest) SetReason(reason string) error {
    r.reason = reason
    r.Set("reason", reason)
    return nil
}

func (r AlibabaAliqinFlowWalletSendFlowRequest) GetReason() string {
    return r.reason
}

func (r *AlibabaAliqinFlowWalletSendFlowRequest) SetAlways(always string) error {
    r.always = always
    r.Set("always", always)
    return nil
}

func (r AlibabaAliqinFlowWalletSendFlowRequest) GetAlways() string {
    return r.always
}

