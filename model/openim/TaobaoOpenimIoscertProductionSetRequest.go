package openim

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
设置ios证书 APIRequest
taobao.openim.ioscert.production.set

设置ios证书
*/
type TaobaoOpenimIoscertProductionSetRequest struct {
    model.Params

    // 证书密码
    password   string 

    // 证书文件内容,base64编码
    cert   string 

}

func NewTaobaoOpenimIoscertProductionSetRequest() *TaobaoOpenimIoscertProductionSetRequest{
    return &TaobaoOpenimIoscertProductionSetRequest{
        Params: model.NewParams(),
    }
}

func (r TaobaoOpenimIoscertProductionSetRequest) GetApiMethodName() string {
    return "taobao.openim.ioscert.production.set"
}

func (r TaobaoOpenimIoscertProductionSetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *TaobaoOpenimIoscertProductionSetRequest) SetPassword(password string) error {
    r.password = password
    r.Set("password", password)
    return nil
}

func (r TaobaoOpenimIoscertProductionSetRequest) GetPassword() string {
    return r.password
}

func (r *TaobaoOpenimIoscertProductionSetRequest) SetCert(cert string) error {
    r.cert = cert
    r.Set("cert", cert)
    return nil
}

func (r TaobaoOpenimIoscertProductionSetRequest) GetCert() string {
    return r.cert
}

