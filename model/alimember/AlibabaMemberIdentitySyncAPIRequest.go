package alimember

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabamemberidentitysyncAPIRequest 会员身份信息同步 API请求
// alibaba.member.identity.sync
//
// 会员身份信息同步
type AlibabamemberidentitysyncAPIRequest struct {
	model.Params
	// 会员身份同步信息
	_syncDto *SyncMemberIdentityDto
}

// NewAlibabamemberidentitysyncRequest 初始化AlibabamemberidentitysyncAPIRequest对象
func NewAlibabamemberidentitysyncRequest() *AlibabamemberidentitysyncAPIRequest {
	return &AlibabamemberidentitysyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabamemberidentitysyncAPIRequest) GetApiMethodName() string {
	return "alibaba.member.identity.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabamemberidentitysyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabamemberidentitysyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncDto is SyncDto Setter
// 会员身份同步信息
func (r *AlibabamemberidentitysyncAPIRequest) SetSyncDto(_syncDto *SyncMemberIdentityDto) error {
	r._syncDto = _syncDto
	r.Set("sync_dto", _syncDto)
	return nil
}

// GetSyncDto SyncDto Getter
func (r AlibabamemberidentitysyncAPIRequest) GetSyncDto() *SyncMemberIdentityDto {
	return r._syncDto
}
