package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkChannelOrderUsercancelAPIRequest 用户发起售中取消 API请求
// alibaba.wdk.channel.order.usercancel
//
// 用户发起售中取消
type AlibabaWdkChannelOrderUsercancelAPIRequest struct {
	model.Params
	// 取消信息
	_userCancelInfo *OrderUserCancelInfo
}

// NewAlibabaWdkChannelOrderUsercancelRequest 初始化AlibabaWdkChannelOrderUsercancelAPIRequest对象
func NewAlibabaWdkChannelOrderUsercancelRequest() *AlibabaWdkChannelOrderUsercancelAPIRequest {
	return &AlibabaWdkChannelOrderUsercancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkChannelOrderUsercancelAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.channel.order.usercancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkChannelOrderUsercancelAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is UserCancelInfo Setter
// 取消信息
func (r *AlibabaWdkChannelOrderUsercancelAPIRequest) SetUserCancelInfo(_userCancelInfo *OrderUserCancelInfo) error {
	r._userCancelInfo = _userCancelInfo
	r.Set("user_cancel_info", _userCancelInfo)
	return nil
}

// Get UserCancelInfo Getter
func (r AlibabaWdkChannelOrderUsercancelAPIRequest) GetUserCancelInfo() *OrderUserCancelInfo {
	return r._userCancelInfo
}
