package alimember

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamembersyncAPIRequest 会员信息同步 API请求
// alibaba.member.sync
//
// 会员信息同步
type AlibabamembersyncAPIRequest struct {
	model.Params
	// 会员同步信息
	_syncMember *SyncMemberDto
}

// NewAlibabamembersyncRequest 初始化AlibabamembersyncAPIRequest对象
func NewAlibabamembersyncRequest() *AlibabamembersyncAPIRequest {
	return &AlibabamembersyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamembersyncAPIRequest) GetApiMethodName() string {
	return "alibaba.member.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamembersyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamembersyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncMember is SyncMember Setter
// 会员同步信息
func (r *AlibabamembersyncAPIRequest) SetSyncMember(_syncMember *SyncMemberDto) error {
	r._syncMember = _syncMember
	r.Set("sync_member", _syncMember)
	return nil
}

// GetSyncMember SyncMember Getter
func (r AlibabamembersyncAPIRequest) GetSyncMember() *SyncMemberDto {
	return r._syncMember
}
