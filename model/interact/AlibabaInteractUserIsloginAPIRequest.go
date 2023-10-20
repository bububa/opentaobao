package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabainteractuserisloginAPIRequest 校验用户是否已经登录 API请求
// alibaba.interact.user.islogin
//
// API的功能是校验用户是否登录，ISV调用接口的时候，通过此接口映射到mtop.interact.user.islogin接口上，因此接口只是做一个给ISV注册调用api的入口，没有发生真实的RPC
type AlibabainteractuserisloginAPIRequest struct {
	model.Params
	// 用户nick
	_buyerNick string
}

// NewAlibabainteractuserisloginRequest 初始化AlibabainteractuserisloginAPIRequest对象
func NewAlibabainteractuserisloginRequest() *AlibabainteractuserisloginAPIRequest {
	return &AlibabainteractuserisloginAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabainteractuserisloginAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.user.islogin"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabainteractuserisloginAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabainteractuserisloginAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetBuyerNick is BuyerNick Setter
// 用户nick
func (r *AlibabainteractuserisloginAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r AlibabainteractuserisloginAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}
