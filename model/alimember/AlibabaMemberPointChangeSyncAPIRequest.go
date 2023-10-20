package alimember

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamemberpointchangesyncAPIRequest 成长值/积分变更记录同步 API请求
// alibaba.member.point.change.sync
//
// 成长值/积分变更记录同步
type AlibabamemberpointchangesyncAPIRequest struct {
	model.Params
	// 变更对象
	_syncMemberPointChangeDto *SyncMemberPointChangeDto
}

// NewAlibabamemberpointchangesyncRequest 初始化AlibabamemberpointchangesyncAPIRequest对象
func NewAlibabamemberpointchangesyncRequest() *AlibabamemberpointchangesyncAPIRequest {
	return &AlibabamemberpointchangesyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamemberpointchangesyncAPIRequest) GetApiMethodName() string {
	return "alibaba.member.point.change.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamemberpointchangesyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamemberpointchangesyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncMemberPointChangeDto is SyncMemberPointChangeDto Setter
// 变更对象
func (r *AlibabamemberpointchangesyncAPIRequest) SetSyncMemberPointChangeDto(_syncMemberPointChangeDto *SyncMemberPointChangeDto) error {
	r._syncMemberPointChangeDto = _syncMemberPointChangeDto
	r.Set("sync_member_point_change_dto", _syncMemberPointChangeDto)
	return nil
}

// GetSyncMemberPointChangeDto SyncMemberPointChangeDto Getter
func (r AlibabamemberpointchangesyncAPIRequest) GetSyncMemberPointChangeDto() *SyncMemberPointChangeDto {
	return r._syncMemberPointChangeDto
}
