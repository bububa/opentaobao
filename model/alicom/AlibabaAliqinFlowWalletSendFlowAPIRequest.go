package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinFlowWalletSendFlowAPIRequest 流量发放 API请求
// alibaba.aliqin.flow.wallet.send.flow
//
// 阿里通信流量下发功能，允许用户补发
type AlibabaAliqinFlowWalletSendFlowAPIRequest struct {
	model.Params
	// 混淆用户名
	_buyerNick string
	// 真实用户名称，如果填写这个字段，buyer_nick失效
	_realNick string
	// 唯一流水号，字母+数字组合
	_serial string
	// 流量
	_flow string
	// 购物送
	_reason string
	// 设置true为始终发送成功
	_always string
}

// NewAlibabaAliqinFlowWalletSendFlowRequest 初始化AlibabaAliqinFlowWalletSendFlowAPIRequest对象
func NewAlibabaAliqinFlowWalletSendFlowRequest() *AlibabaAliqinFlowWalletSendFlowAPIRequest {
	return &AlibabaAliqinFlowWalletSendFlowAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAliqinFlowWalletSendFlowAPIRequest) GetApiMethodName() string {
	return "alibaba.aliqin.flow.wallet.send.flow"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAliqinFlowWalletSendFlowAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBuyerNick is BuyerNick Setter
// 混淆用户名
func (r *AlibabaAliqinFlowWalletSendFlowAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r AlibabaAliqinFlowWalletSendFlowAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}

// SetRealNick is RealNick Setter
// 真实用户名称，如果填写这个字段，buyer_nick失效
func (r *AlibabaAliqinFlowWalletSendFlowAPIRequest) SetRealNick(_realNick string) error {
	r._realNick = _realNick
	r.Set("real_nick", _realNick)
	return nil
}

// GetRealNick RealNick Getter
func (r AlibabaAliqinFlowWalletSendFlowAPIRequest) GetRealNick() string {
	return r._realNick
}

// SetSerial is Serial Setter
// 唯一流水号，字母+数字组合
func (r *AlibabaAliqinFlowWalletSendFlowAPIRequest) SetSerial(_serial string) error {
	r._serial = _serial
	r.Set("serial", _serial)
	return nil
}

// GetSerial Serial Getter
func (r AlibabaAliqinFlowWalletSendFlowAPIRequest) GetSerial() string {
	return r._serial
}

// SetFlow is Flow Setter
// 流量
func (r *AlibabaAliqinFlowWalletSendFlowAPIRequest) SetFlow(_flow string) error {
	r._flow = _flow
	r.Set("flow", _flow)
	return nil
}

// GetFlow Flow Getter
func (r AlibabaAliqinFlowWalletSendFlowAPIRequest) GetFlow() string {
	return r._flow
}

// SetReason is Reason Setter
// 购物送
func (r *AlibabaAliqinFlowWalletSendFlowAPIRequest) SetReason(_reason string) error {
	r._reason = _reason
	r.Set("reason", _reason)
	return nil
}

// GetReason Reason Getter
func (r AlibabaAliqinFlowWalletSendFlowAPIRequest) GetReason() string {
	return r._reason
}

// SetAlways is Always Setter
// 设置true为始终发送成功
func (r *AlibabaAliqinFlowWalletSendFlowAPIRequest) SetAlways(_always string) error {
	r._always = _always
	r.Set("always", _always)
	return nil
}

// GetAlways Always Getter
func (r AlibabaAliqinFlowWalletSendFlowAPIRequest) GetAlways() string {
	return r._always
}
