package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaailabstmallgenieauthdevicevalidauthcodeAPIRequest 根据authcode查询绑定结果 API请求
// alibaba.ailabs.tmallgenie.auth.device.validauthcode
//
// 根据authcode查询绑定结果
type AlibabaailabstmallgenieauthdevicevalidauthcodeAPIRequest struct {
	model.Params
	// authcode
	_authcode string
}

// NewAlibabaailabstmallgenieauthdevicevalidauthcodeRequest 初始化AlibabaailabstmallgenieauthdevicevalidauthcodeAPIRequest对象
func NewAlibabaailabstmallgenieauthdevicevalidauthcodeRequest() *AlibabaailabstmallgenieauthdevicevalidauthcodeAPIRequest {
	return &AlibabaailabstmallgenieauthdevicevalidauthcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaailabstmallgenieauthdevicevalidauthcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.validauthcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaailabstmallgenieauthdevicevalidauthcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaailabstmallgenieauthdevicevalidauthcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAuthcode is Authcode Setter
// authcode
func (r *AlibabaailabstmallgenieauthdevicevalidauthcodeAPIRequest) SetAuthcode(_authcode string) error {
	r._authcode = _authcode
	r.Set("authcode", _authcode)
	return nil
}

// GetAuthcode Authcode Getter
func (r AlibabaailabstmallgenieauthdevicevalidauthcodeAPIRequest) GetAuthcode() string {
	return r._authcode
}
