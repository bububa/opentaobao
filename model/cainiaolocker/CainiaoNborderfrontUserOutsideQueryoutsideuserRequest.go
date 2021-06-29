package cainiaolocker

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询外部小件员休息 APIRequest
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

func NewCainiaoNborderfrontUserOutsideQueryoutsideuserRequest() *CainiaoNborderfrontUserOutsideQueryoutsideuserRequest{
    return &CainiaoNborderfrontUserOutsideQueryoutsideuserRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoNborderfrontUserOutsideQueryoutsideuserRequest) GetApiMethodName() string {
    return "cainiao.nborderfront.user.outside.queryoutsideuser"
}

func (r CainiaoNborderfrontUserOutsideQueryoutsideuserRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoNborderfrontUserOutsideQueryoutsideuserRequest) SetCpCode(cpCode string) error {
    r.cpCode = cpCode
    r.Set("cp_code", cpCode)
    return nil
}

func (r CainiaoNborderfrontUserOutsideQueryoutsideuserRequest) GetCpCode() string {
    return r.cpCode
}

func (r *CainiaoNborderfrontUserOutsideQueryoutsideuserRequest) SetCpUserId(cpUserId string) error {
    r.cpUserId = cpUserId
    r.Set("cp_user_id", cpUserId)
    return nil
}

func (r CainiaoNborderfrontUserOutsideQueryoutsideuserRequest) GetCpUserId() string {
    return r.cpUserId
}

