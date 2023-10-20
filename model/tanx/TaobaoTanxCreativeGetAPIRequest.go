package tanx

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaotanxcreativegetAPIRequest 获取单个审核创意状态 API请求
// taobao.tanx.creative.get
//
// 获取单个审核创意状态
type TaobaotanxcreativegetAPIRequest struct {
	model.Params
	// dsp用户身份认证的TOKEN
	_token string
	// 创意ID
	_creativeId string
	// DSP的memberId
	_memberId int64
	// 当前时间戳，1970-01-01后的秒数
	_signTime int64
}

// NewTaobaotanxcreativegetRequest 初始化TaobaotanxcreativegetAPIRequest对象
func NewTaobaotanxcreativegetRequest() *TaobaotanxcreativegetAPIRequest {
	return &TaobaotanxcreativegetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaotanxcreativegetAPIRequest) GetApiMethodName() string {
	return "taobao.tanx.creative.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaotanxcreativegetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaotanxcreativegetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetToken is Token Setter
// dsp用户身份认证的TOKEN
func (r *TaobaotanxcreativegetAPIRequest) SetToken(_token string) error {
	r._token = _token
	r.Set("token", _token)
	return nil
}

// GetToken Token Getter
func (r TaobaotanxcreativegetAPIRequest) GetToken() string {
	return r._token
}

// SetCreativeId is CreativeId Setter
// 创意ID
func (r *TaobaotanxcreativegetAPIRequest) SetCreativeId(_creativeId string) error {
	r._creativeId = _creativeId
	r.Set("creative_id", _creativeId)
	return nil
}

// GetCreativeId CreativeId Getter
func (r TaobaotanxcreativegetAPIRequest) GetCreativeId() string {
	return r._creativeId
}

// SetMemberId is MemberId Setter
// DSP的memberId
func (r *TaobaotanxcreativegetAPIRequest) SetMemberId(_memberId int64) error {
	r._memberId = _memberId
	r.Set("member_id", _memberId)
	return nil
}

// GetMemberId MemberId Getter
func (r TaobaotanxcreativegetAPIRequest) GetMemberId() int64 {
	return r._memberId
}

// SetSignTime is SignTime Setter
// 当前时间戳，1970-01-01后的秒数
func (r *TaobaotanxcreativegetAPIRequest) SetSignTime(_signTime int64) error {
	r._signTime = _signTime
	r.Set("sign_time", _signTime)
	return nil
}

// GetSignTime SignTime Getter
func (r TaobaotanxcreativegetAPIRequest) GetSignTime() int64 {
	return r._signTime
}
