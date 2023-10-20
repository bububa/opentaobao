package alimember

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMemberPointChangeSyncAPIRequest 成长值/积分变更记录同步 API请求
// alibaba.member.point.change.sync
//
// 成长值/积分变更记录同步
type AlibabaMemberPointChangeSyncAPIRequest struct {
	model.Params
	// 变更对象
	_syncMemberPointChangeDto *SyncMemberPointChangeDto
}

// NewAlibabaMemberPointChangeSyncRequest 初始化AlibabaMemberPointChangeSyncAPIRequest对象
func NewAlibabaMemberPointChangeSyncRequest() *AlibabaMemberPointChangeSyncAPIRequest {
	return &AlibabaMemberPointChangeSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMemberPointChangeSyncAPIRequest) Reset() {
	r._syncMemberPointChangeDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMemberPointChangeSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.member.point.change.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMemberPointChangeSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMemberPointChangeSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncMemberPointChangeDto is SyncMemberPointChangeDto Setter
// 变更对象
func (r *AlibabaMemberPointChangeSyncAPIRequest) SetSyncMemberPointChangeDto(_syncMemberPointChangeDto *SyncMemberPointChangeDto) error {
	r._syncMemberPointChangeDto = _syncMemberPointChangeDto
	r.Set("sync_member_point_change_dto", _syncMemberPointChangeDto)
	return nil
}

// GetSyncMemberPointChangeDto SyncMemberPointChangeDto Getter
func (r AlibabaMemberPointChangeSyncAPIRequest) GetSyncMemberPointChangeDto() *SyncMemberPointChangeDto {
	return r._syncMemberPointChangeDto
}

var poolAlibabaMemberPointChangeSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMemberPointChangeSyncRequest()
	},
}

// GetAlibabaMemberPointChangeSyncRequest 从 sync.Pool 获取 AlibabaMemberPointChangeSyncAPIRequest
func GetAlibabaMemberPointChangeSyncAPIRequest() *AlibabaMemberPointChangeSyncAPIRequest {
	return poolAlibabaMemberPointChangeSyncAPIRequest.Get().(*AlibabaMemberPointChangeSyncAPIRequest)
}

// ReleaseAlibabaMemberPointChangeSyncAPIRequest 将 AlibabaMemberPointChangeSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaMemberPointChangeSyncAPIRequest(v *AlibabaMemberPointChangeSyncAPIRequest) {
	v.Reset()
	poolAlibabaMemberPointChangeSyncAPIRequest.Put(v)
}
