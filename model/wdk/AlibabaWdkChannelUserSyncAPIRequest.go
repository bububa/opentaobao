package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkChannelUserSyncAPIRequest 会员同步 API请求
// alibaba.wdk.channel.user.sync
//
// 会员同步
type AlibabaWdkChannelUserSyncAPIRequest struct {
	model.Params
	// 会员信息
	_userSyncInfo *UserSyncInfo
}

// NewAlibabaWdkChannelUserSyncRequest 初始化AlibabaWdkChannelUserSyncAPIRequest对象
func NewAlibabaWdkChannelUserSyncRequest() *AlibabaWdkChannelUserSyncAPIRequest {
	return &AlibabaWdkChannelUserSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkChannelUserSyncAPIRequest) Reset() {
	r._userSyncInfo = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkChannelUserSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.channel.user.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkChannelUserSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkChannelUserSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserSyncInfo is UserSyncInfo Setter
// 会员信息
func (r *AlibabaWdkChannelUserSyncAPIRequest) SetUserSyncInfo(_userSyncInfo *UserSyncInfo) error {
	r._userSyncInfo = _userSyncInfo
	r.Set("user_sync_info", _userSyncInfo)
	return nil
}

// GetUserSyncInfo UserSyncInfo Getter
func (r AlibabaWdkChannelUserSyncAPIRequest) GetUserSyncInfo() *UserSyncInfo {
	return r._userSyncInfo
}

var poolAlibabaWdkChannelUserSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkChannelUserSyncRequest()
	},
}

// GetAlibabaWdkChannelUserSyncRequest 从 sync.Pool 获取 AlibabaWdkChannelUserSyncAPIRequest
func GetAlibabaWdkChannelUserSyncAPIRequest() *AlibabaWdkChannelUserSyncAPIRequest {
	return poolAlibabaWdkChannelUserSyncAPIRequest.Get().(*AlibabaWdkChannelUserSyncAPIRequest)
}

// ReleaseAlibabaWdkChannelUserSyncAPIRequest 将 AlibabaWdkChannelUserSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkChannelUserSyncAPIRequest(v *AlibabaWdkChannelUserSyncAPIRequest) {
	v.Reset()
	poolAlibabaWdkChannelUserSyncAPIRequest.Put(v)
}
