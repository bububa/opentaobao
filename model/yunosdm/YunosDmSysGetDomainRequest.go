package yunosdm

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取动态域名 API请求
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

// 初始化YunosDmSysGetDomainRequest对象
func NewYunosDmSysGetDomainRequest() *YunosDmSysGetDomainRequest{
    return &YunosDmSysGetDomainRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r YunosDmSysGetDomainRequest) GetApiMethodName() string {
    return "yunos.dm.sys.get.domain"
}

// IRequest interface 方法, 获取API参数
func (r YunosDmSysGetDomainRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Make Setter
// 制造商
func (r *YunosDmSysGetDomainRequest) SetMake(make string) error {
    r.make = make
    r.Set("make", make)
    return nil
}

// Make Getter
func (r YunosDmSysGetDomainRequest) GetMake() string {
    return r.make
}
// Model Setter
// 设备类型
func (r *YunosDmSysGetDomainRequest) SetModel(model string) error {
    r.model = model
    r.Set("model", model)
    return nil
}

// Model Getter
func (r YunosDmSysGetDomainRequest) GetModel() string {
    return r.model
}
// Sn Setter
// 序列号
func (r *YunosDmSysGetDomainRequest) SetSn(sn string) error {
    r.sn = sn
    r.Set("sn", sn)
    return nil
}

// Sn Getter
func (r YunosDmSysGetDomainRequest) GetSn() string {
    return r.sn
}
