package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest
翱象商家自有渠道 交易订单取消 API请求
alibaba.tcls.aelophy.merchant.channel.order.cancel

翱象小程序用户取消订单 */
type AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest struct {
	model.Params
	// 取消信息
	_userCancelInfo *OrderUserCancelInfo
}

// NewAlibabaTclsAelophyMerchantChannelOrderCancelRequest 初始化AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest对象
func NewAlibabaTclsAelophyMerchantChannelOrderCancelRequest() *AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest {
	return &AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest) GetApiMethodName() string {
	return "alibaba.tcls.aelophy.merchant.channel.order.cancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is UserCancelInfo Setter
// 取消信息
func (r *AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest) SetUserCancelInfo(_userCancelInfo *OrderUserCancelInfo) error {
	r._userCancelInfo = _userCancelInfo
	r.Set("user_cancel_info", _userCancelInfo)
	return nil
}

// Get UserCancelInfo Getter
func (r AlibabaTclsAelophyMerchantChannelOrderCancelAPIRequest) GetUserCancelInfo() *OrderUserCancelInfo {
	return r._userCancelInfo
}
