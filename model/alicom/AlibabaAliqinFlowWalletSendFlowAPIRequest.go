package alicom

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAliqinFlowWalletSendFlowAPIRequest
流量发放 API请求
alibaba.aliqin.flow.wallet.send.flow

阿里通信流量下发功能，允许用户补发 */
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

// New
