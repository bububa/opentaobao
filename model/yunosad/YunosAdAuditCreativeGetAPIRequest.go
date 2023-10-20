package yunosad

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosAdAuditCreativeGetAPIRequest 获取单个创意审核状态 API请求
// yunos.ad.audit.creative.get
//
// 获取单个创意审核状态
type YunosAdAuditCreativeGetAPIRequest struct {
	model.Params
	// 第三方广告创意id
	_creativeId string
	// 第三方的dspId
	_memberId int64
}

// NewYunosAdAuditCreativeGetRequest 初始化YunosAdAuditCreativeGetAPIRequest对象
func NewYunosAdAuditCreativeGetRequest() *YunosAdAuditCreativeGetAPIRequest {
	return &YunosAdAuditCreativeGetAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosAdAuditCreativeGetAPIRequest) Reset() {
	r._creativeId = ""
	r._memberId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosAdAuditCreativeGetAPIRequest) GetApiMethodName() string {
	return "yunos.ad.audit.creative.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosAdAuditCreativeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosAdAuditCreativeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCreativeId is CreativeId Setter
// 第三方广告创意id
func (r *YunosAdAuditCreativeGetAPIRequest) SetCreativeId(_creativeId string) error {
	r._creativeId = _creativeId
	r.Set("creative_id", _creativeId)
	return nil
}

// GetCreativeId CreativeId Getter
func (r YunosAdAuditCreativeGetAPIRequest) GetCreativeId() string {
	return r._creativeId
}

// SetMemberId is MemberId Setter
// 第三方的dspId
func (r *YunosAdAuditCreativeGetAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// GetMemberId MemberId Getter
func (r YunosAdAuditCreativeGetAPIRequest) GetMemberId() int64 {
	return r._memberId
}

var poolYunosAdAuditCreativeGetAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosAdAuditCreativeGetRequest()
	},
}

// GetYunosAdAuditCreativeGetRequest 从 sync.Pool 获取 YunosAdAuditCreativeGetAPIRequest
func GetYunosAdAuditCreativeGetAPIRequest() *YunosAdAuditCreativeGetAPIRequest {
	return poolYunosAdAuditCreativeGetAPIRequest.Get().(*YunosAdAuditCreativeGetAPIRequest)
}

// ReleaseYunosAdAuditCreativeGetAPIRequest 将 YunosAdAuditCreativeGetAPIRequest 放入 sync.Pool
func ReleaseYunosAdAuditCreativeGetAPIRequest(v *YunosAdAuditCreativeGetAPIRequest) {
	v.Reset()
	poolYunosAdAuditCreativeGetAPIRequest.Put(v)
}
