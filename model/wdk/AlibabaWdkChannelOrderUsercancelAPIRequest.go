package wdk

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkChannelOrderUsercancelAPIRequest) Reset() {
	r._userCancelInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkChannelOrderUsercancelAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.channel.order.usercancel"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkChannelOrderUsercancelAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkChannelOrderUsercancelAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserCancelInfo is UserCancelInfo Setter
// 取消信息
func (r *AlibabaWdkChannelOrderUsercancelAPIRequest) SetUserCancelInfo(_userCancelInfo *OrderUserCancelInfo) error {
	r._userCancelInfo = _userCancelInfo
	r.Set("user_cancel_info", _userCancelInfo)
	return nil
}

// GetUserCancelInfo UserCancelInfo Getter
func (r AlibabaWdkChannelOrderUsercancelAPIRequest) GetUserCancelInfo() *OrderUserCancelInfo {
	return r._userCancelInfo
}

var poolAlibabaWdkChannelOrderUsercancelAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkChannelOrderUsercancelRequest()
	},
}

// GetAlibabaWdkChannelOrderUsercancelRequest 从 sync.Pool 获取 AlibabaWdkChannelOrderUsercancelAPIRequest
func GetAlibabaWdkChannelOrderUsercancelAPIRequest() *AlibabaWdkChannelOrderUsercancelAPIRequest {
	return poolAlibabaWdkChannelOrderUsercancelAPIRequest.Get().(*AlibabaWdkChannelOrderUsercancelAPIRequest)
}

// ReleaseAlibabaWdkChannelOrderUsercancelAPIRequest 将 AlibabaWdkChannelOrderUsercancelAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkChannelOrderUsercancelAPIRequest(v *AlibabaWdkChannelOrderUsercancelAPIRequest) {
	v.Reset()
	poolAlibabaWdkChannelOrderUsercancelAPIRequest.Put(v)
}
