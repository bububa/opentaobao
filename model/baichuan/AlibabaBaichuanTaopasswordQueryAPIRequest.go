package baichuan

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibababaichuantaopasswordqueryAPIRequest 查询解析淘口令 API请求
// alibaba.baichuan.taopassword.query
//
// 查询，解析淘口令
type AlibababaichuantaopasswordqueryAPIRequest struct {
	model.Params
	// 淘口令
	_passwordContent string
}

// NewAlibababaichuantaopasswordqueryRequest 初始化AlibababaichuantaopasswordqueryAPIRequest对象
func NewAlibababaichuantaopasswordqueryRequest() *AlibababaichuantaopasswordqueryAPIRequest {
	return &AlibababaichuantaopasswordqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibababaichuantaopasswordqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.baichuan.taopassword.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibababaichuantaopasswordqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibababaichuantaopasswordqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPasswordContent is PasswordContent Setter
// 淘口令
func (r *AlibababaichuantaopasswordqueryAPIRequest) SetPasswordContent(_passwordContent string) error {
	r._passwordContent = _passwordContent
	r.Set("password_content", _passwordContent)
	return nil
}

// GetPasswordContent PasswordContent Getter
func (r AlibababaichuantaopasswordqueryAPIRequest) GetPasswordContent() string {
	return r._passwordContent
}
