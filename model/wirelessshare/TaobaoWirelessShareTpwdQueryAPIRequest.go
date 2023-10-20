package wirelessshare

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaowirelesssharetpwdqueryAPIRequest 查询解析淘口令 API请求
// taobao.wireless.share.tpwd.query
//
// 查询解析淘口令
type TaobaowirelesssharetpwdqueryAPIRequest struct {
	model.Params
	// 淘口令
	_passwordContent string
}

// NewTaobaowirelesssharetpwdqueryRequest 初始化TaobaowirelesssharetpwdqueryAPIRequest对象
func NewTaobaowirelesssharetpwdqueryRequest() *TaobaowirelesssharetpwdqueryAPIRequest {
	return &TaobaowirelesssharetpwdqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaowirelesssharetpwdqueryAPIRequest) GetApiMethodName() string {
	return "taobao.wireless.share.tpwd.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaowirelesssharetpwdqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaowirelesssharetpwdqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPasswordContent is PasswordContent Setter
// 淘口令
func (r *TaobaowirelesssharetpwdqueryAPIRequest) SetPasswordContent(_passwordContent string) error {
	r._passwordContent = _passwordContent
	r.Set("password_content", _passwordContent)
	return nil
}

// GetPasswordContent PasswordContent Getter
func (r TaobaowirelesssharetpwdqueryAPIRequest) GetPasswordContent() string {
	return r._passwordContent
}
