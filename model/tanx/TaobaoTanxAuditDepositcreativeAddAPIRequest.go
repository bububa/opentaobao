package tanx

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTanxAuditDepositcreativeAddAPIRequest dsp托管创意新增接口 API请求
// taobao.tanx.audit.depositcreative.add
//
// dsp托管创意新增接口
type TaobaoTanxAuditDepositcreativeAddAPIRequest struct {
	model.Params
	// dsp用户身份认证的TOKEN
	_token string
	// DSP的memberId
	_memberId int64
	// 当前时间戳，1970-01-01后的秒数
	_signTime int64
	// 托管创意信息
	_creative *CreativeInfoDto
}

// NewTaobaoTanxAuditDepositcreativeAddRequest 初始化TaobaoTanxAuditDepositcreativeAddAPIRequest对象
func NewTaobaoTanxAuditDepositcreativeAddRequest() *TaobaoTanxAuditDepositcreativeAddAPIRequest {
	return &TaobaoTanxAuditDepositcreativeAddAPIRequest{
		Params: model.NewParams(4),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *TaobaoTanxAuditDepositcreativeAddAPIRequest) Reset() {
	r._token = ""
	r._memberId = 0
	r._signTime = 0
	r._creative = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTanxAuditDepositcreativeAddAPIRequest) GetApiMethodName() string {
	return "taobao.tanx.audit.depositcreative.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTanxAuditDepositcreativeAddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoTanxAuditDepositcreativeAddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// dsp用户身份认证的TOKEN
func (r *TaobaoTanxAuditDepositcreativeAddAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaoTanxAuditDepositcreativeAddAPIRequest) GetToken() string {
	return r._token
}

// SetMemberId is MemberId Setter
// DSP的memberId
func (r *TaobaoTanxAuditDepositcreativeAddAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// GetMemberId MemberId Getter
func (r TaobaoTanxAuditDepositcreativeAddAPIRequest) GetMemberId() int64 {
	return r._memberId
}

// SetSignTime is SignTime Setter
// 当前时间戳，1970-01-01后的秒数
func (r *TaobaoTanxAuditDepositcreativeAddAPIRequest) SetSignTime(_signTime int64) error {
	r._signTime = _signTime
	r.Set("sign_time", _signTime)
	return nil
}

// GetSignTime SignTime Getter
func (r TaobaoTanxAuditDepositcreativeAddAPIRequest) GetSignTime() int64 {
	return r._signTime
}

// SetCreative is Creative Setter
// 托管创意信息
func (r *TaobaoTanxAuditDepositcreativeAddAPIRequest) SetCreative(_creative *CreativeInfoDto) error {
	r._creative = _creative
	r.Set("creative", _creative)
	return nil
}

// GetCreative Creative Getter
func (r TaobaoTanxAuditDepositcreativeAddAPIRequest) GetCreative() *CreativeInfoDto {
	return r._creative
}

var poolTaobaoTanxAuditDepositcreativeAddAPIRequest = sync.Pool{
	New: func() any {
		return NewTaobaoTanxAuditDepositcreativeAddRequest()
	},
}

// GetTaobaoTanxAuditDepositcreativeAddRequest 从 sync.Pool 获取 TaobaoTanxAuditDepositcreativeAddAPIRequest
func GetTaobaoTanxAuditDepositcreativeAddAPIRequest() *TaobaoTanxAuditDepositcreativeAddAPIRequest {
	return poolTaobaoTanxAuditDepositcreativeAddAPIRequest.Get().(*TaobaoTanxAuditDepositcreativeAddAPIRequest)
}

// ReleaseTaobaoTanxAuditDepositcreativeAddAPIRequest 将 TaobaoTanxAuditDepositcreativeAddAPIRequest 放入 sync.Pool
func ReleaseTaobaoTanxAuditDepositcreativeAddAPIRequest(v *TaobaoTanxAuditDepositcreativeAddAPIRequest) {
	v.Reset()
	poolTaobaoTanxAuditDepositcreativeAddAPIRequest.Put(v)
}
