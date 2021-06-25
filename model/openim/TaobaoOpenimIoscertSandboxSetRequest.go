package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设置开发环境证书 APIRequest
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

func NewTaobaoOpenimIoscertSandboxSetRequest() *TaobaoOpenimIoscertSandboxSetRequest{
    return &TaobaoOpenimIoscertSandboxSetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenimIoscertSandboxSetRequest) GetApiMethodName() string {
    return "taobao.openim.ioscert.sandbox.set"
}

func (r TaobaoOpenimIoscertSandboxSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenimIoscertSandboxSetRequest) SetCert(cert string) error {
    r.cert = cert
    r.Set("cert", cert)
    return nil
}

func (r TaobaoOpenimIoscertSandboxSetRequest) GetCert() string {
    return r.cert
}

func (r *TaobaoOpenimIoscertSandboxSetRequest) SetPassword(password string) error {
    r.password = password
    r.Set("password", password)
    return nil
}

func (r TaobaoOpenimIoscertSandboxSetRequest) GetPassword() string {
    return r.password
}

