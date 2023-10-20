package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest 翱象商家自有渠道 交易订单取消 API请求
// alibaba.tcls.aelophy.merchant.channel.order.cancel
//
// 翱象小程序用户取消订单
type AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest struct {
	model.Params
	// 取消信息
	_userCancelInfo *OrderUserCancelInfo
}

// NewAlibabaTclsAelophyMerchantChannelOrderCancelRequest 初始化AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest对象
func NewAlibabaTclsAelophyMerchantChannelOrderCancelRequest() *AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest {
	return &AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest) Reset() {
	r._userCancelInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.channel.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserCancelInfo is UserCancelInfo Setter
// 取消信息
func (r *AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest) SetUserCancelInfo(_userCancelInfo *OrderUserCancelInfo) error {
	r._userCancelInfo = _userCancelInfo
	r.Set("user_cancel_info", _userCancelInfo)
	return nil
}

// GetUserCancelInfo UserCancelInfo Getter
func (r AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest) GetUserCancelInfo() *OrderUserCancelInfo {
	return r._userCancelInfo
}

var poolAlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTclsAelophyMerchantChannelOrderCancelRequest()
	},
}

// GetAlibabaTclsAelophyMerchantChannelOrderCancelRequest 从 sync.Pool 获取 AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest
func GetAlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest() *AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest {
	return poolAlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest.Get().(*AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest)
}

// ReleaseAlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest 将 AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest 放入 sync.Pool
func ReleaseAlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest(v *AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest) {
	v.Reset()
	poolAlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest.Put(v)
}
