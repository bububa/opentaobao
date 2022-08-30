package tmallgenie

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIRequest 根据authcode查询绑定结果 API请求
// alibaba.ailabs.tmallgenie.auth.device.validauthcode
//
// 根据authcode查询绑定结果
type AlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIRequest struct {
	model.Params
	// authcode
	_authcode string
}

// NewAlibabaAilabsTmallgenieAuthDeviceValidauthcodeRequest 初始化AlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIRequest对象
func NewAlibabaAilabsTmallgenieAuthDeviceValidauthcodeRequest() *AlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIRequest {
	return &AlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.validauthcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAuthcode is Authcode Setter
// authcode
func (r *AlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIRequest) SetAuthcode(_authcode string) error {
	r._authcode = _authcode
	r.Set("authcode", _authcode)
	return nil
}

// GetAuthcode Authcode Getter
func (r AlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIRequest) GetAuthcode() string {
	return r._authcode
}
