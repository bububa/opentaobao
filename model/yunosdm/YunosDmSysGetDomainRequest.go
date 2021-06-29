package yunosdm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取动态域名 APIRequest
yunos.dm.sys.get.domain

返回alios ucp后端域名
*/
type YunosDmSysGetDomainRequest struct {
    model.Params

    // 制造商
    make   string 

    // 设备类型
    model   string 

    // 序列号
    sn   string 

}

func NewYunosDmSysGetDomainRequest() *YunosDmSysGetDomainRequest{
    return &YunosDmSysGetDomainRequest{
        Params: model.NewParams(),
    }
}

func (r YunosDmSysGetDomainRequest) GetApiMethodName() string {
    return "yunos.dm.sys.get.domain"
}

func (r YunosDmSysGetDomainRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *YunosDmSysGetDomainRequest) SetMake(make string) error {
    r.make = make
    r.Set("make", make)
    return nil
}

func (r YunosDmSysGetDomainRequest) GetMake() string {
    return r.make
}

func (r *YunosDmSysGetDomainRequest) SetModel(model string) error {
    r.model = model
    r.Set("model", model)
    return nil
}

func (r YunosDmSysGetDomainRequest) GetModel() string {
    return r.model
}

func (r *YunosDmSysGetDomainRequest) SetSn(sn string) error {
    r.sn = sn
    r.Set("sn", sn)
    return nil
}

func (r YunosDmSysGetDomainRequest) GetSn() string {
    return r.sn
}

