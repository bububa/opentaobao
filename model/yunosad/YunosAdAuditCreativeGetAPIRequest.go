package yunosad

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosAdAuditCreativeGetAPIRequest
获取单个创意审核状态 API请求
yunos.ad.audit.creative.get

获取单个创意审核状态 */
type YunosAdAuditCreativeGetAPIRequest struct {
	model.Params
	// 第三方的dspId
	_memberId int64
	// 第三方广告创意id
	_creativeId string
}

// NewYunosAdAuditCreativeGetRequest 初始化YunosAdAuditCreativeGetAPIRequest对象
func NewYunosAdAuditCreativeGetRequest() *YunosAdAuditCreativeGetAPIRequest {
	return &YunosAdAuditCreativeGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosAdAuditCreativeGetAPIRequest) GetApiMethodName() string {
	return "yunos.ad.audit.creative.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosAdAuditCreativeGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MemberId Setter
// 第三方的dspId
func (r *YunosAdAuditCreativeGetAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// Get MemberId Getter
func (r YunosAdAuditCreativeGetAPIRequest) GetMemberId() int64 {
	return r._memberId
}

// Set is CreativeId Setter
// 第三方广告创意id
func (r *YunosAdAuditCreativeGetAPIRequest) SetCreativeId(_creativeId string) error {
	r._creativeId = _creativeId
	r.Set("creative_id", _creativeId)
	return nil
}

// Get CreativeId Getter
func (r YunosAdAuditCreativeGetAPIRequest) GetCreativeId() string {
	return r._creativeId
}
