package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabatclsaelophymerchantchannelordercancelAPIRequest 翱象商家自有渠道 交易订单取消 API请求
// alibaba.tcls.aelophy.merchant.channel.order.cancel
//
// 翱象小程序用户取消订单
type AlibabatclsaelophymerchantchannelordercancelAPIRequest struct {
	model.Params
	// 取消信息
	_userCancelInfo *OrderUserCancelInfo
}

// NewAlibabatclsaelophymerchantchannelordercancelRequest 初始化AlibabatclsaelophymerchantchannelordercancelAPIRequest对象
func NewAlibabatclsaelophymerchantchannelordercancelRequest() *AlibabatclsaelophymerchantchannelordercancelAPIRequest {
	return &AlibabatclsaelophymerchantchannelordercancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabatclsaelophymerchantchannelordercancelAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.channel.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabatclsaelophymerchantchannelordercancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabatclsaelophymerchantchannelordercancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserCancelInfo is UserCancelInfo Setter
// 取消信息
func (r *AlibabatclsaelophymerchantchannelordercancelAPIRequest) SetUserCancelInfo(_userCancelInfo *OrderUserCancelInfo) error {
	r._userCancelInfo = _userCancelInfo
	r.Set("user_cancel_info", _userCancelInfo)
	return nil
}

// GetUserCancelInfo UserCancelInfo Getter
func (r AlibabatclsaelophymerchantchannelordercancelAPIRequest) GetUserCancelInfo() *OrderUserCancelInfo {
	return r._userCancelInfo
}
