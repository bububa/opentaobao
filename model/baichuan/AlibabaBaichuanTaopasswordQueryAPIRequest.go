package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaBaichuanTaopasswordQueryAPIRequest 查询解析淘口令 API请求
// alibaba.baichuan.taopassword.query
//
// 查询，解析淘口令
type AlibabaBaichuanTaopasswordQueryAPIRequest struct {
	model.Params
	// 淘口令
	_passwordContent string
}

// NewAlibabaBaichuanTaopasswordQueryRequest 初始化AlibabaBaichuanTaopasswordQueryAPIRequest对象
func NewAlibabaBaichuanTaopasswordQueryRequest() *AlibabaBaichuanTaopasswordQueryAPIRequest {
	return &AlibabaBaichuanTaopasswordQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaBaichuanTaopasswordQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.baichuan.taopassword.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaBaichuanTaopasswordQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetPasswordContent is PasswordContent Setter
// 淘口令
func (r *AlibabaBaichuanTaopasswordQueryAPIRequest) SetPasswordContent(_passwordContent string) error {
	r._passwordContent = _passwordContent
	r.Set("password_content", _passwordContent)
	return nil
}

// GetPasswordContent PasswordContent Getter
func (r AlibabaBaichuanTaopasswordQueryAPIRequest) GetPasswordContent() string {
	return r._passwordContent
}
