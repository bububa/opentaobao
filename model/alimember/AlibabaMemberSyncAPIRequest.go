package alimember

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMemberSyncAPIRequest 会员信息同步 API请求
// alibaba.member.sync
//
// 会员信息同步
type AlibabaMemberSyncAPIRequest struct {
	model.Params
	// 会员同步信息
	_syncMember *SyncMemberDto
}

// NewAlibabaMemberSyncRequest 初始化AlibabaMemberSyncAPIRequest对象
func NewAlibabaMemberSyncRequest() *AlibabaMemberSyncAPIRequest {
	return &AlibabaMemberSyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMemberSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.member.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMemberSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetSyncMember is SyncMember Setter
// 会员同步信息
func (r *AlibabaMemberSyncAPIRequest) SetSyncMember(_syncMember *SyncMemberDto) error {
	r._syncMember = _syncMember
	r.Set("sync_member", _syncMember)
	return nil
}

// GetSyncMember SyncMember Getter
func (r AlibabaMemberSyncAPIRequest) GetSyncMember() *SyncMemberDto {
	return r._syncMember
}
