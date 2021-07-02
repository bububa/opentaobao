package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaInteractUserIsloginAPIRequest 校验用户是否已经登录 API请求
// alibaba.interact.user.islogin
//
// API的功能是校验用户是否登录，ISV调用接口的时候，通过此接口映射到mtop.interact.user.islogin接口上，因此接口只是做一个给ISV注册调用api的入口，没有发生真实的RPC
type AlibabaInteractUserIsloginAPIRequest struct {
	model.Params
	// 用户nick
	_buyerNick string
}

// NewAlibabaInteractUserIsloginRequest 初始化AlibabaInteractUserIsloginAPIRequest对象
func NewAlibabaInteractUserIsloginRequest() *AlibabaInteractUserIsloginAPIRequest {
	return &AlibabaInteractUserIsloginAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaInteractUserIsloginAPIRequest) GetApiMethodName() string {
	return "alibaba.interact.user.islogin"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaInteractUserIsloginAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetBuyerNick is BuyerNick Setter
// 用户nick
func (r *AlibabaInteractUserIsloginAPIRequest) SetBuyerNick(_buyerNick string) error {
	r._buyerNick = _buyerNick
	r.Set("buyer_nick", _buyerNick)
	return nil
}

// GetBuyerNick BuyerNick Getter
func (r AlibabaInteractUserIsloginAPIRequest) GetBuyerNick() string {
	return r._buyerNick
}
