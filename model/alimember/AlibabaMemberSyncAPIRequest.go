package alimember

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMemberSyncAPIRequest) Reset() {
	r._syncMember = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMemberSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.member.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMemberSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMemberSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaMemberSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMemberSyncRequest()
	},
}

// GetAlibabaMemberSyncRequest 从 sync.Pool 获取 AlibabaMemberSyncAPIRequest
func GetAlibabaMemberSyncAPIRequest() *AlibabaMemberSyncAPIRequest {
	return poolAlibabaMemberSyncAPIRequest.Get().(*AlibabaMemberSyncAPIRequest)
}

// ReleaseAlibabaMemberSyncAPIRequest 将 AlibabaMemberSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaMemberSyncAPIRequest(v *AlibabaMemberSyncAPIRequest) {
	v.Reset()
	poolAlibabaMemberSyncAPIRequest.Put(v)
}
