package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkchannelusersyncAPIRequest 会员同步 API请求
// alibaba.wdk.channel.user.sync
//
// 会员同步
type AlibabawdkchannelusersyncAPIRequest struct {
	model.Params
	// 会员信息
	_userSyncInfo *UserSyncInfo
}

// NewAlibabawdkchannelusersyncRequest 初始化AlibabawdkchannelusersyncAPIRequest对象
func NewAlibabawdkchannelusersyncRequest() *AlibabawdkchannelusersyncAPIRequest {
	return &AlibabawdkchannelusersyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkchannelusersyncAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.channel.user.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkchannelusersyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkchannelusersyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetUserSyncInfo is UserSyncInfo Setter
// 会员信息
func (r *AlibabawdkchannelusersyncAPIRequest) SetUserSyncInfo(_userSyncInfo *UserSyncInfo) error {
	r._userSyncInfo = _userSyncInfo
	r.Set("user_sync_info", _userSyncInfo)
	return nil
}

// GetUserSyncInfo UserSyncInfo Getter
func (r AlibabawdkchannelusersyncAPIRequest) GetUserSyncInfo() *UserSyncInfo {
	return r._userSyncInfo
}
