package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设置ios证书 API请求
taobao.openim.ioscert.production.set

设置ios证书
*/
type TaobaoOpenimIoscertProductionSetRequest struct {
    model.Params
    // 证书密码
    _password   string
    // 证书文件内容,base64编码
    _cert   string
}

// 初始化TaobaoOpenimIoscertProductionSetRequest对象
func NewTaobaoOpenimIoscertProductionSetRequest() *TaobaoOpenimIoscertProductionSetRequest{
    return &TaobaoOpenimIoscertProductionSetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r TaobaoOpenimIoscertProductionSetRequest) GetApiMethodName() string {
    return "taobao.openim.ioscert.production.set"
}

// IRequest interface 方法, 获取API参数
func (r TaobaoOpenimIoscertProductionSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Password Setter
// 证书密码
func (r *TaobaoOpenimIoscertProductionSetRequest) SetPassword(_password string) error {
    r._password = _password
    r.Set("password", _password)
    return nil
}

// Password Getter
func (r TaobaoOpenimIoscertProductionSetRequest) GetPassword() string {
    return r._password
}
// Cert Setter
// 证书文件内容,base64编码
func (r *TaobaoOpenimIoscertProductionSetRequest) SetCert(_cert string) error {
    r._cert = _cert
    r.Set("cert", _cert)
    return nil
}

// Cert Getter
func (r TaobaoOpenimIoscertProductionSetRequest) GetCert() string {
    return r._cert
}
