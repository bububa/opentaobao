package tanx

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTanxAuditCreativeAddAPIRequest 创意预审新增接口 API请求
// taobao.tanx.audit.creative.add
//
// 创意预审新增接口
type TaobaoTanxAuditCreativeAddAPIRequest struct {
	model.Params
	// dsp用户身份认证的TOKEN
	_token string
	// DSP的memberId
	_memberId int64
	// 当前时间戳，1970-01-01后的秒数
	_signTime int64
	// 预审核创意对象
	_creative *CreativeParamDto
}

// NewTaobaoTanxAuditCreativeAddRequest 初始化TaobaoTanxAuditCreativeAddAPIRequest对象
func NewTaobaoTanxAuditCreativeAddRequest() *TaobaoTanxAuditCreativeAddAPIRequest {
	return &TaobaoTanxAuditCreativeAddAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTanxAuditCreativeAddAPIRequest) Reset() {
	r._token = ""
	r._memberId = 0
	r._signTime = 0
	r._creative = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTanxAuditCreativeAddAPIRequest) GetApiMethodName() string {
	return "taobao.tanx.audit.creative.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTanxAuditCreativeAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTanxAuditCreativeAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// dsp用户身份认证的TOKEN
func (r *TaobaoTanxAuditCreativeAddAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaoTanxAuditCreativeAddAPIRequest) GetToken() string {
	return r._token
}

// SetMemberId is MemberId Setter
// DSP的memberId
func (r *TaobaoTanxAuditCreativeAddAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// GetMemberId MemberId Getter
func (r TaobaoTanxAuditCreativeAddAPIRequest) GetMemberId() int64 {
	return r._memberId
}

// SetSignTime is SignTime Setter
// 当前时间戳，1970-01-01后的秒数
func (r *TaobaoTanxAuditCreativeAddAPIRequest) SetSignTime(_signTime int64) error {
	r._signTime = _signTime
	r.Set("sign_time", _signTime)
	return nil
}

// GetSignTime SignTime Getter
func (r TaobaoTanxAuditCreativeAddAPIRequest) GetSignTime() int64 {
	return r._signTime
}

// SetCreative is Creative Setter
// 预审核创意对象
func (r *TaobaoTanxAuditCreativeAddAPIRequest) SetCreative(_creative *CreativeParamDto) error {
	r._creative = _creative
	r.Set("creative", _creative)
	return nil
}

// GetCreative Creative Getter
func (r TaobaoTanxAuditCreativeAddAPIRequest) GetCreative() *CreativeParamDto {
	return r._creative
}

var poolTaobaoTanxAuditCreativeAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTanxAuditCreativeAddRequest()
	},
}

// GetTaobaoTanxAuditCreativeAddRequest 从 sync.Pool 获取 TaobaoTanxAuditCreativeAddAPIRequest
func GetTaobaoTanxAuditCreativeAddAPIRequest() *TaobaoTanxAuditCreativeAddAPIRequest {
	return poolTaobaoTanxAuditCreativeAddAPIRequest.Get().(*TaobaoTanxAuditCreativeAddAPIRequest)
}

// ReleaseTaobaoTanxAuditCreativeAddAPIRequest 将 TaobaoTanxAuditCreativeAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoTanxAuditCreativeAddAPIRequest(v *TaobaoTanxAuditCreativeAddAPIRequest) {
	v.Reset()
	poolTaobaoTanxAuditCreativeAddAPIRequest.Put(v)
}
