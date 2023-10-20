package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotanxauditdepositcreativeaddAPIRequest dsp托管创意新增接口 API请求
// taobao.tanx.audit.depositcreative.add
//
// dsp托管创意新增接口
type TaobaotanxauditdepositcreativeaddAPIRequest struct {
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

// NewTaobaotanxauditdepositcreativeaddRequest 初始化TaobaotanxauditdepositcreativeaddAPIRequest对象
func NewTaobaotanxauditdepositcreativeaddRequest() *TaobaotanxauditdepositcreativeaddAPIRequest {
	return &TaobaotanxauditdepositcreativeaddAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotanxauditdepositcreativeaddAPIRequest) GetApiMethodName() string {
	return "taobao.tanx.audit.depositcreative.add"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotanxauditdepositcreativeaddAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotanxauditdepositcreativeaddAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// dsp用户身份认证的TOKEN
func (r *TaobaotanxauditdepositcreativeaddAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaotanxauditdepositcreativeaddAPIRequest) GetToken() string {
	return r._token
}

// SetMemberId is MemberId Setter
// DSP的memberId
func (r *TaobaotanxauditdepositcreativeaddAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// GetMemberId MemberId Getter
func (r TaobaotanxauditdepositcreativeaddAPIRequest) GetMemberId() int64 {
	return r._memberId
}

// SetSignTime is SignTime Setter
// 当前时间戳，1970-01-01后的秒数
func (r *TaobaotanxauditdepositcreativeaddAPIRequest) SetSignTime(_signTime int64) error {
	r._signTime = _signTime
	r.Set("sign_time", _signTime)
	return nil
}

// GetSignTime SignTime Getter
func (r TaobaotanxauditdepositcreativeaddAPIRequest) GetSignTime() int64 {
	return r._signTime
}

// SetCreative is Creative Setter
// 托管创意信息
func (r *TaobaotanxauditdepositcreativeaddAPIRequest) SetCreative(_creative *CreativeInfoDto) error {
	r._creative = _creative
	r.Set("creative", _creative)
	return nil
}

// GetCreative Creative Getter
func (r TaobaotanxauditdepositcreativeaddAPIRequest) GetCreative() *CreativeInfoDto {
	return r._creative
}
