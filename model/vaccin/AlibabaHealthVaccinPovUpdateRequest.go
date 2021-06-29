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
type AlibabaHealthVaccinPovUpdateRequest struct {
    model.Params
    // 接种点联系电话
    telephone   string
    // 接种点具体地址
    address   string
    // 接种点介绍
    description   string
    // 接种点编码
    povNo   string
    // 接种点名称
    povName   string
    // 服务时间
    businessTime   string
}

// 初始化AlibabaHealthVaccinPovUpdateRequest对象
func NewAlibabaHealthVaccinPovUpdateRequest() *AlibabaHealthVaccinPovUpdateRequest{
    return &AlibabaHealthVaccinPovUpdateRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaHealthVaccinPovUpdateRequest) GetApiMethodName() string {
    return "alibaba.health.vaccin.pov.update"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaHealthVaccinPovUpdateRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Telephone Setter
// 接种点联系电话
func (r *AlibabaHealthVaccinPovUpdateRequest) SetTelephone(telephone string) error {
    r.telephone = telephone
    r.Set("telephone", telephone)
    return nil
}

// Telephone Getter
func (r AlibabaHealthVaccinPovUpdateRequest) GetTelephone() string {
    return r.telephone
}
// Address Setter
// 接种点具体地址
func (r *AlibabaHealthVaccinPovUpdateRequest) SetAddress(address string) error {
    r.address = address
    r.Set("address", address)
    return nil
}

// Address Getter
func (r AlibabaHealthVaccinPovUpdateRequest) GetAddress() string {
    return r.address
}
// Description Setter
// 接种点介绍
func (r *AlibabaHealthVaccinPovUpdateRequest) SetDescription(description string) error {
    r.description = description
    r.Set("description", description)
    return nil
}

// Description Getter
func (r AlibabaHealthVaccinPovUpdateRequest) GetDescription() string {
    return r.description
}
// PovNo Setter
// 接种点编码
func (r *AlibabaHealthVaccinPovUpdateRequest) SetPovNo(povNo string) error {
    r.povNo = povNo
    r.Set("pov_no", povNo)
    return nil
}

// PovNo Getter
func (r AlibabaHealthVaccinPovUpdateRequest) GetPovNo() string {
    return r.povNo
}
// PovName Setter
// 接种点名称
func (r *AlibabaHealthVaccinPovUpdateRequest) SetPovName(povName string) error {
    r.povName = povName
    r.Set("pov_name", povName)
    return nil
}

// PovName Getter
func (r AlibabaHealthVaccinPovUpdateRequest) GetPovName() string {
    return r.povName
}
// BusinessTime Setter
// 服务时间
func (r *AlibabaHealthVaccinPovUpdateRequest) SetBusinessTime(businessTime string) error {
    r.businessTime = businessTime
    r.Set("business_time", businessTime)
    return nil
}

// BusinessTime Getter
func (r AlibabaHealthVaccinPovUpdateRequest) GetBusinessTime() string {
    return r.businessTime
}
