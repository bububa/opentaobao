package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTanxAuditCreativeAddAPIRequest
创意预审新增接口 API请求
taobao.tanx.audit.creative.add

创意预审新增接口 */
type TaobaoTanxAuditCreativeAddAPIRequest struct {
	model.Params
	// DSP的memberId
	_memberId int64
	// dsp用户身份认证的TOKEN
	_token string
	// 当前时间戳，1970-01-01后的秒数
	_signTime int64
	// 预审核创意对象
	_creative *CreativeParamDto
}

// NewTaobaoTanxAuditCreativeAddRequest 初始化TaobaoTanxAuditCreativeAddAPIRequest对象
func NewTaobaoTanxAuditCreativeAddRequest() *TaobaoTanxAuditCreativeAddAPIRequest {
	return &TaobaoTanxAuditCreativeAddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTanxAuditCreativeAddAPIRequest) GetApiMethodName() string {
	return "taobao.tanx.audit.creative.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTanxAuditCreativeAddAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is MemberId Setter
// DSP的memberId
func (r *TaobaoTanxAuditCreativeAddAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// Get MemberId Getter
func (r TaobaoTanxAuditCreativeAddAPIRequest) GetMemberId() int64 {
	return r._memberId
}

// Set is Token Setter
// dsp用户身份认证的TOKEN
func (r *TaobaoTanxAuditCreativeAddAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// Get Token Getter
func (r TaobaoTanxAuditCreativeAddAPIRequest) GetToken() string {
	return r._token
}

// Set is SignTime Setter
// 当前时间戳，1970-01-01后的秒数
func (r *TaobaoTanxAuditCreativeAddAPIRequest) SetSignTime(_signTime int64) error {
	r._signTime = _signTime
	r.Set("sign_time", _signTime)
	return nil
}

// Get SignTime Getter
func (r TaobaoTanxAuditCreativeAddAPIRequest) GetSignTime() int64 {
	return r._signTime
}

// Set is Creative Setter
// 预审核创意对象
func (r *TaobaoTanxAuditCreativeAddAPIRequest) SetCreative(_creative *CreativeParamDto) error {
	r._creative = _creative
	r.Set("creative", _creative)
	return nil
}

// Get Creative Getter
func (r TaobaoTanxAuditCreativeAddAPIRequest) GetCreative() *CreativeParamDto {
	return r._creative
}
