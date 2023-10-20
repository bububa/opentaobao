package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotanxauditcreativeaddAPIRequest 创意预审新增接口 API请求
// taobao.tanx.audit.creative.add
//
// 创意预审新增接口
type TaobaotanxauditcreativeaddAPIRequest struct {
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

// NewTaobaotanxauditcreativeaddRequest 初始化TaobaotanxauditcreativeaddAPIRequest对象
func NewTaobaotanxauditcreativeaddRequest() *TaobaotanxauditcreativeaddAPIRequest {
	return &TaobaotanxauditcreativeaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotanxauditcreativeaddAPIRequest) GetApiMethodName() string {
	return "taobao.tanx.audit.creative.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotanxauditcreativeaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotanxauditcreativeaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// dsp用户身份认证的TOKEN
func (r *TaobaotanxauditcreativeaddAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaotanxauditcreativeaddAPIRequest) GetToken() string {
	return r._token
}

// SetMemberId is MemberId Setter
// DSP的memberId
func (r *TaobaotanxauditcreativeaddAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// GetMemberId MemberId Getter
func (r TaobaotanxauditcreativeaddAPIRequest) GetMemberId() int64 {
	return r._memberId
}

// SetSignTime is SignTime Setter
// 当前时间戳，1970-01-01后的秒数
func (r *TaobaotanxauditcreativeaddAPIRequest) SetSignTime(_signTime int64) error {
	r._signTime = _signTime
	r.Set("sign_time", _signTime)
	return nil
}

// GetSignTime SignTime Getter
func (r TaobaotanxauditcreativeaddAPIRequest) GetSignTime() int64 {
	return r._signTime
}

// SetCreative is Creative Setter
// 预审核创意对象
func (r *TaobaotanxauditcreativeaddAPIRequest) SetCreative(_creative *CreativeParamDto) error {
	r._creative = _creative
	r.Set("creative", _creative)
	return nil
}

// GetCreative Creative Getter
func (r TaobaotanxauditcreativeaddAPIRequest) GetCreative() *CreativeParamDto {
	return r._creative
}
