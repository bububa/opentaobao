package vaccin

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
新增/变更接种点信息 API请求
alibaba.health.vaccin.pov.update

ISV 将疫苗的接种点信息同步到免疫规划中心，提醒用户接种时可提供接种点详情。
*/
type AlibabaHealthVaccinPovUpdateAPIRequest struct {
    model.Params
    // 接种点联系电话
    _telephone   string
    // 接种点具体地址
    _address   string
    // 接种点介绍
    _description   string
    // 接种点编码
    _povNo   string
    // 接种点名称
    _povName   string
    // 服务时间
    _businessTime   string
}

// 初始化AlibabaHealthVaccinPovUpdateAPIRequest对象
func NewAlibabaHealthVaccinPovUpdateRequest() *AlibabaHealthVaccinPovUpdateAPIRequest{
    return &AlibabaHealthVaccinPovUpdateAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinPovUpdateAPIRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.pov.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinPovUpdateAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Telephone Setter
// 接种点联系电话
func (r *AlibabaHealthVaccinPovUpdateAPIRequest) SetTelephone(_telephone string) error {
    r._telephone = _telephone
    r.Set("telephone", _telephone)
    return nil
}

// Telephone Getter
func (r AlibabaHealthVaccinPovUpdateAPIRequest) GetTelephone() string {
    return r._telephone
}
// Address Setter
// 接种点具体地址
func (r *AlibabaHealthVaccinPovUpdateAPIRequest) SetAddress(_address string) error {
    r._address = _address
    r.Set("address", _address)
    return nil
}

// Address Getter
func (r AlibabaHealthVaccinPovUpdateAPIRequest) GetAddress() string {
    return r._address
}
// Description Setter
// 接种点介绍
func (r *AlibabaHealthVaccinPovUpdateAPIRequest) SetDescription(_description string) error {
    r._description = _description
    r.Set("description", _description)
    return nil
}

// Description Getter
func (r AlibabaHealthVaccinPovUpdateAPIRequest) GetDescription() string {
    return r._description
}
// PovNo Setter
// 接种点编码
func (r *AlibabaHealthVaccinPovUpdateAPIRequest) SetPovNo(_povNo string) error {
    r._povNo = _povNo
    r.Set("pov_no", _povNo)
    return nil
}

// PovNo Getter
func (r AlibabaHealthVaccinPovUpdateAPIRequest) GetPovNo() string {
    return r._povNo
}
// PovName Setter
// 接种点名称
func (r *AlibabaHealthVaccinPovUpdateAPIRequest) SetPovName(_povName string) error {
    r._povName = _povName
    r.Set("pov_name", _povName)
    return nil
}

// PovName Getter
func (r AlibabaHealthVaccinPovUpdateAPIRequest) GetPovName() string {
    return r._povName
}
// BusinessTime Setter
// 服务时间
func (r *AlibabaHealthVaccinPovUpdateAPIRequest) SetBusinessTime(_businessTime string) error {
    r._businessTime = _businessTime
    r.Set("business_time", _businessTime)
    return nil
}

// BusinessTime Getter
func (r AlibabaHealthVaccinPovUpdateAPIRequest) GetBusinessTime() string {
    return r._businessTime
}
