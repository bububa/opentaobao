package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTanxAuditCreativeModifyAPIRequest 创意修改接口 API请求
// taobao.tanx.audit.creative.modify
//
// 创意修改接口
type TaobaoTanxAuditCreativeModifyAPIRequest struct {
	model.Params
	// DSP用户ID
	_memberId int64
	// dsp用户身份认证的TOKEN
	_token string
	// 当前时间戳，1970-01-01后的秒数
	_signTime int64
}

// NewTaobaoTanxAuditCreativeModifyRequest 初始化TaobaoTanxAuditCreativeModifyAPIRequest对象
func NewTaobaoTanxAuditCreativeModifyRequest() *TaobaoTanxAuditCreativeModifyAPIRequest {
	return &TaobaoTanxAuditCreativeModifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoTanxAuditCreativeModifyAPIRequest) GetApiMethodName() string {
	return "taobao.tanx.audit.creative.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoTanxAuditCreativeModifyAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetMemberId is MemberId Setter
// DSP用户ID
func (r *TaobaoTanxAuditCreativeModifyAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// GetMemberId MemberId Getter
func (r TaobaoTanxAuditCreativeModifyAPIRequest) GetMemberId() int64 {
	return r._memberId
}

// SetToken is Token Setter
// dsp用户身份认证的TOKEN
func (r *TaobaoTanxAuditCreativeModifyAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaoTanxAuditCreativeModifyAPIRequest) GetToken() string {
	return r._token
}

// SetSignTime is SignTime Setter
// 当前时间戳，1970-01-01后的秒数
func (r *TaobaoTanxAuditCreativeModifyAPIRequest) SetSignTime(_signTime int64) error {
	r._signTime = _signTime
	r.Set("sign_time", _signTime)
	return nil
}

// GetSignTime SignTime Getter
func (r TaobaoTanxAuditCreativeModifyAPIRequest) GetSignTime() int64 {
	return r._signTime
}
