package wdk

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkChannelUserSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.channel.user.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkChannelUserSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
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
