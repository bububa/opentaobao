package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkchannelorderusercancelAPIRequest 用户发起售中取消 API请求
// alibaba.wdk.channel.order.usercancel
//
// 用户发起售中取消
type AlibabawdkchannelorderusercancelAPIRequest struct {
	model.Params
	// 取消信息
	_userCancelInfo *OrderUserCancelInfo
}

// NewAlibabawdkchannelorderusercancelRequest 初始化AlibabawdkchannelorderusercancelAPIRequest对象
func NewAlibabawdkchannelorderusercancelRequest() *AlibabawdkchannelorderusercancelAPIRequest {
	return &AlibabawdkchannelorderusercancelAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkchannelorderusercancelAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.channel.order.usercancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkchannelorderusercancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkchannelorderusercancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserCancelInfo is UserCancelInfo Setter
// 取消信息
func (r *AlibabawdkchannelorderusercancelAPIRequest) SetUserCancelInfo(_userCancelInfo *OrderUserCancelInfo) error {
	r._userCancelInfo = _userCancelInfo
	r.Set("user_cancel_info", _userCancelInfo)
	return nil
}

// GetUserCancelInfo UserCancelInfo Getter
func (r AlibabawdkchannelorderusercancelAPIRequest) GetUserCancelInfo() *OrderUserCancelInfo {
	return r._userCancelInfo
}
