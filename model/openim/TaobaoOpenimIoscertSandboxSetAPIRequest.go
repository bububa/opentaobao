package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设置开发环境证书 API请求
taobao.openim.ioscert.sandbox.set

设置开发环境证书
*/
type TaobaoOpenimIoscertSandboxSetAPIRequest struct {
    model.Params
    // 证书内容,base64编码
    _cert   string
    // 系统自动生成
    _password   string
}

// 初始化TaobaoOpenimIoscertSandboxSetAPIRequest对象
func NewTaobaoOpenimIoscertSandboxSetRequest() *TaobaoOpenimIoscertSandboxSetAPIRequest{
    return &TaobaoOpenimIoscertSandboxSetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimIoscertSandboxSetAPIRequest) GetApiMethodName() string {
    return "taobao.openim.ioscert.sandbox.set"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimIoscertSandboxSetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Cert Setter
// 证书内容,base64编码
func (r *TaobaoOpenimIoscertSandboxSetAPIRequest) SetCert(_cert string) error {
    r._cert = _cert
    r.Set("cert", _cert)
    return nil
}

// Cert Getter
func (r TaobaoOpenimIoscertSandboxSetAPIRequest) GetCert() string {
    return r._cert
}
// Password Setter
// 系统自动生成
func (r *TaobaoOpenimIoscertSandboxSetAPIRequest) SetPassword(_password string) error {
    r._password = _password
    r.Set("password", _password)
    return nil
}

// Password Getter
func (r TaobaoOpenimIoscertSandboxSetAPIRequest) GetPassword() string {
    return r._password
}
