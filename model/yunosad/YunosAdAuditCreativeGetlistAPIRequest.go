package yunosad

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosAdAuditCreativeGetlistAPIRequest 批量获取创意审核状态 API请求
// yunos.ad.audit.creative.getlist
//
// 批量获取创意审核状态
type YunosAdAuditCreativeGetlistAPIRequest struct {
	model.Params
	// 创意列表
	_creativeIds []string
	// 状态
	_status string
	// 第三方DSP的id
	_memberId int64
}

// NewYunosAdAuditCreativeGetlistRequest 初始化YunosAdAuditCreativeGetlistAPIRequest对象
func NewYunosAdAuditCreativeGetlistRequest() *YunosAdAuditCreativeGetlistAPIRequest {
	return &YunosAdAuditCreativeGetlistAPIRequest{
		Params: model.NewParams(3),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosAdAuditCreativeGetlistAPIRequest) Reset() {
	r._creativeIds = r._creativeIds[:0]
	r._status = ""
	r._memberId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosAdAuditCreativeGetlistAPIRequest) GetApiMethodName() string {
	return "yunos.ad.audit.creative.getlist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosAdAuditCreativeGetlistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosAdAuditCreativeGetlistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreativeIds is CreativeIds Setter
// 创意列表
func (r *YunosAdAuditCreativeGetlistAPIRequest) SetCreativeIds(_creativeIds []string) error {
	r._creativeIds = _creativeIds
	r.Set("creative_ids", _creativeIds)
	return nil
}

// GetCreativeIds CreativeIds Getter
func (r YunosAdAuditCreativeGetlistAPIRequest) GetCreativeIds() []string {
	return r._creativeIds
}

// SetStatus is Status Setter
// 状态
func (r *YunosAdAuditCreativeGetlistAPIRequest) SetStatus(_status string) error {
	r._status = _status
	r.Set("status", _status)
	return nil
}

// GetStatus Status Getter
func (r YunosAdAuditCreativeGetlistAPIRequest) GetStatus() string {
	return r._status
}

// SetMemberId is MemberId Setter
// 第三方DSP的id
func (r *YunosAdAuditCreativeGetlistAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// GetMemberId MemberId Getter
func (r YunosAdAuditCreativeGetlistAPIRequest) GetMemberId() int64 {
	return r._memberId
}

var poolYunosAdAuditCreativeGetlistAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosAdAuditCreativeGetlistRequest()
	},
}

// GetYunosAdAuditCreativeGetlistRequest 从 sync.Pool 获取 YunosAdAuditCreativeGetlistAPIRequest
func GetYunosAdAuditCreativeGetlistAPIRequest() *YunosAdAuditCreativeGetlistAPIRequest {
	return poolYunosAdAuditCreativeGetlistAPIRequest.Get().(*YunosAdAuditCreativeGetlistAPIRequest)
}

// ReleaseYunosAdAuditCreativeGetlistAPIRequest 将 YunosAdAuditCreativeGetlistAPIRequest 放入 sync.Pool
func ReleaseYunosAdAuditCreativeGetlistAPIRequest(v *YunosAdAuditCreativeGetlistAPIRequest) {
	v.Reset()
	poolYunosAdAuditCreativeGetlistAPIRequest.Put(v)
}
