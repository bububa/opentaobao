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
type TaobaoOpenimIoscertSandboxSetRequest struct {
    model.Params
    // 证书内容,base64编码
    cert   string
    // 系统自动生成
    password   string
}

// 初始化TaobaoOpenimIoscertSandboxSetRequest对象
func NewTaobaoOpenimIoscertSandboxSetRequest() *TaobaoOpenimIoscertSandboxSetRequest{
    return &TaobaoOpenimIoscertSandboxSetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimIoscertSandboxSetRequest) GetApiMethodName() string {
    return "taobao.openim.ioscert.sandbox.set"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimIoscertSandboxSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Cert Setter
// 证书内容,base64编码
func (r *TaobaoOpenimIoscertSandboxSetRequest) SetCert(cert string) error {
    r.cert = cert
    r.Set("cert", cert)
    return nil
}

// Cert Getter
func (r TaobaoOpenimIoscertSandboxSetRequest) GetCert() string {
    return r.cert
}
// Password Setter
// 系统自动生成
func (r *TaobaoOpenimIoscertSandboxSetRequest) SetPassword(password string) error {
    r.password = password
    r.Set("password", password)
    return nil
}

// Password Getter
func (r TaobaoOpenimIoscertSandboxSetRequest) GetPassword() string {
    return r.password
}
