package tmallgenie

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIRequest) Reset() {
	r._authcode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIRequest) GetApiMethodName() string {
	return "alibaba.ailabs.tmallgenie.auth.device.validauthcode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAilabsTmallgenieAuthDeviceValidauthcodeRequest()
	},
}

// GetAlibabaAilabsTmallgenieAuthDeviceValidauthcodeRequest 从 sync.Pool 获取 AlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIRequest
func GetAlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIRequest() *AlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIRequest {
	return poolAlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIRequest.Get().(*AlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIRequest)
}

// ReleaseAlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIRequest 将 AlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIRequest(v *AlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIRequest) {
	v.Reset()
	poolAlibabaAilabsTmallgenieAuthDeviceValidauthcodeAPIRequest.Put(v)
}
