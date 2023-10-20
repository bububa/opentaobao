package alimember

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaMemberIdentitySyncAPIRequest 会员身份信息同步 API请求
// alibaba.member.identity.sync
//
// 会员身份信息同步
type AlibabaMemberIdentitySyncAPIRequest struct {
	model.Params
	// 会员身份同步信息
	_syncDto *SyncMemberIdentityDto
}

// NewAlibabaMemberIdentitySyncRequest 初始化AlibabaMemberIdentitySyncAPIRequest对象
func NewAlibabaMemberIdentitySyncRequest() *AlibabaMemberIdentitySyncAPIRequest {
	return &AlibabaMemberIdentitySyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaMemberIdentitySyncAPIRequest) Reset() {
	r._syncDto = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaMemberIdentitySyncAPIRequest) GetApiMethodName() string {
	return "alibaba.member.identity.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaMemberIdentitySyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaMemberIdentitySyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSyncDto is SyncDto Setter
// 会员身份同步信息
func (r *AlibabaMemberIdentitySyncAPIRequest) SetSyncDto(_syncDto *SyncMemberIdentityDto) error {
	r._syncDto = _syncDto
	r.Set("sync_dto", _syncDto)
	return nil
}

// GetSyncDto SyncDto Getter
func (r AlibabaMemberIdentitySyncAPIRequest) GetSyncDto() *SyncMemberIdentityDto {
	return r._syncDto
}

var poolAlibabaMemberIdentitySyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaMemberIdentitySyncRequest()
	},
}

// GetAlibabaMemberIdentitySyncRequest 从 sync.Pool 获取 AlibabaMemberIdentitySyncAPIRequest
func GetAlibabaMemberIdentitySyncAPIRequest() *AlibabaMemberIdentitySyncAPIRequest {
	return poolAlibabaMemberIdentitySyncAPIRequest.Get().(*AlibabaMemberIdentitySyncAPIRequest)
}

// ReleaseAlibabaMemberIdentitySyncAPIRequest 将 AlibabaMemberIdentitySyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaMemberIdentitySyncAPIRequest(v *AlibabaMemberIdentitySyncAPIRequest) {
	v.Reset()
	poolAlibabaMemberIdentitySyncAPIRequest.Put(v)
}
