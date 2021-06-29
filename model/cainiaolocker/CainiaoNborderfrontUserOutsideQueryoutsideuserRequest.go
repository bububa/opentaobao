package cainiaolocker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询外部小件员休息 API请求
cainiao.nborderfront.user.outside.queryoutsideuser

采用SPI方式查询外部公司的小件员信息
*/
type CainiaoNborderfrontUserOutsideQueryoutsideuserRequest struct {
    model.Params
    // cpcode
    cpCode   string
    // cp小件员ID
    cpUserId   string
}

// 初始化CainiaoNborderfrontUserOutsideQueryoutsideuserRequest对象
func NewCainiaoNborderfrontUserOutsideQueryoutsideuserRequest() *CainiaoNborderfrontUserOutsideQueryoutsideuserRequest{
    return &CainiaoNborderfrontUserOutsideQueryoutsideuserRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoNborderfrontUserOutsideQueryoutsideuserRequest) GetApiMethodName() string {
    return "cainiao.nborderfront.user.outside.queryoutsideuser"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoNborderfrontUserOutsideQueryoutsideuserRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CpCode Setter
// cpcode
func (r *CainiaoNborderfrontUserOutsideQueryoutsideuserRequest) SetCpCode(cpCode string) error {
    r.cpCode = cpCode
    r.Set("cp_code", cpCode)
    return nil
}

// CpCode Getter
func (r CainiaoNborderfrontUserOutsideQueryoutsideuserRequest) GetCpCode() string {
    return r.cpCode
}
// CpUserId Setter
// cp小件员ID
func (r *CainiaoNborderfrontUserOutsideQueryoutsideuserRequest) SetCpUserId(cpUserId string) error {
    r.cpUserId = cpUserId
    r.Set("cp_user_id", cpUserId)
    return nil
}

// CpUserId Getter
func (r CainiaoNborderfrontUserOutsideQueryoutsideuserRequest) GetCpUserId() string {
    return r.cpUserId
}
