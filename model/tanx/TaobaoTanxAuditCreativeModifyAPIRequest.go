package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotanxauditcreativemodifyAPIRequest 创意修改接口 API请求
// taobao.tanx.audit.creative.modify
//
// 创意修改接口
type TaobaotanxauditcreativemodifyAPIRequest struct {
	model.Params
	// dsp用户身份认证的TOKEN
	_token string
	// DSP用户ID
	_memberId int64
	// 当前时间戳，1970-01-01后的秒数
	_signTime int64
}

// NewTaobaotanxauditcreativemodifyRequest 初始化TaobaotanxauditcreativemodifyAPIRequest对象
func NewTaobaotanxauditcreativemodifyRequest() *TaobaotanxauditcreativemodifyAPIRequest {
	return &TaobaotanxauditcreativemodifyAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotanxauditcreativemodifyAPIRequest) GetApiMethodName() string {
	return "taobao.tanx.audit.creative.modify"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotanxauditcreativemodifyAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotanxauditcreativemodifyAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// dsp用户身份认证的TOKEN
func (r *TaobaotanxauditcreativemodifyAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaotanxauditcreativemodifyAPIRequest) GetToken() string {
	return r._token
}

// SetMemberId is MemberId Setter
// DSP用户ID
func (r *TaobaotanxauditcreativemodifyAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// GetMemberId MemberId Getter
func (r TaobaotanxauditcreativemodifyAPIRequest) GetMemberId() int64 {
	return r._memberId
}

// SetSignTime is SignTime Setter
// 当前时间戳，1970-01-01后的秒数
func (r *TaobaotanxauditcreativemodifyAPIRequest) SetSignTime(_signTime int64) error {
	r._signTime = _signTime
	r.Set("sign_time", _signTime)
	return nil
}

// GetSignTime SignTime Getter
func (r TaobaotanxauditcreativemodifyAPIRequest) GetSignTime() int64 {
	return r._signTime
}
