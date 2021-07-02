package yunosad

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosAdAuditCreativeAddAPIRequest 单个创意预审接口 API请求
// yunos.ad.audit.creative.add
//
// YunOS广告业务第三方DSP单个创意预审接口
type YunosAdAuditCreativeAddAPIRequest struct {
	model.Params
	// 外部dsp的id
	_memberId int64
	// 创意审核入参
	_creative *CreativeParamDto
}

// NewYunosAdAuditCreativeAddRequest 初始化YunosAdAuditCreativeAddAPIRequest对象
func NewYunosAdAuditCreativeAddRequest() *YunosAdAuditCreativeAddAPIRequest {
	return &YunosAdAuditCreativeAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosAdAuditCreativeAddAPIRequest) GetApiMethodName() string {
	return "yunos.ad.audit.creative.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosAdAuditCreativeAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MemberId Setter
// 外部dsp的id
func (r *YunosAdAuditCreativeAddAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// Get MemberId Getter
func (r YunosAdAuditCreativeAddAPIRequest) GetMemberId() int64 {
	return r._memberId
}

// Set is Creative Setter
// 创意审核入参
func (r *YunosAdAuditCreativeAddAPIRequest) SetCreative(_creative *CreativeParamDto) error {
	r._creative = _creative
	r.Set("creative", _creative)
	return nil
}

// Get Creative Getter
func (r YunosAdAuditCreativeAddAPIRequest) GetCreative() *CreativeParamDto {
	return r._creative
}
