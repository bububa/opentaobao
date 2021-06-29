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
    _make   string
    // 设备类型
    _model   string
    // 序列号
    _sn   string
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
func (r *YunosDmSysGetDomainRequest) SetMake(_make string) error {
    r._make = _make
    r.Set("make", _make)
    return nil
}

// Make Getter
func (r YunosDmSysGetDomainRequest) GetMake() string {
    return r._make
}
// Model Setter
// 设备类型
func (r *YunosDmSysGetDomainRequest) SetModel(_model string) error {
    r._model = _model
    r.Set("model", _model)
    return nil
}

// Model Getter
func (r YunosDmSysGetDomainRequest) GetModel() string {
    return r._model
}
// Sn Setter
// 序列号
func (r *YunosDmSysGetDomainRequest) SetSn(_sn string) error {
    r._sn = _sn
    r.Set("sn", _sn)
    return nil
}

// Sn Getter
func (r YunosDmSysGetDomainRequest) GetSn() string {
    return r._sn
}
