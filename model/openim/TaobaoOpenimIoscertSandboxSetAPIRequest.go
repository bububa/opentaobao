package openim

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TaobaoopenimioscertsandboxsetAPIRequest 设置开发环境证书 API请求
// taobao.openim.ioscert.sandbox.set
//
// 设置开发环境证书
type TaobaoopenimioscertsandboxsetAPIRequest struct {
	model.Params
	// 证书内容,base64编码
	_cert string
	// 系统自动生成
	_password string
}

// NewTaobaoopenimioscertsandboxsetRequest 初始化TaobaoopenimioscertsandboxsetAPIRequest对象
func NewTaobaoopenimioscertsandboxsetRequest() *TaobaoopenimioscertsandboxsetAPIRequest {
	return &TaobaoopenimioscertsandboxsetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TaobaoopenimioscertsandboxsetAPIRequest) GetApiMethodName() string {
	return "taobao.openim.ioscert.sandbox.set"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TaobaoopenimioscertsandboxsetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TaobaoopenimioscertsandboxsetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCert is Cert Setter
// 证书内容,base64编码
func (r *TaobaoopenimioscertsandboxsetAPIRequest) SetCert(_cert string) error {
	r._cert = _cert
	r.Set("cert", _cert)
	return nil
}

// GetCert Cert Getter
func (r TaobaoopenimioscertsandboxsetAPIRequest) GetCert() string {
	return r._cert
}

// SetPassword is Password Setter
// 系统自动生成
func (r *TaobaoopenimioscertsandboxsetAPIRequest) SetPassword(_password string) error {
	r._password = _password
	r.Set("password", _password)
	return nil
}

// GetPassword Password Getter
func (r TaobaoopenimioscertsandboxsetAPIRequest) GetPassword() string {
	return r._password
}
