package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询面单服务订购及面单使用情况 APIRequest
cainiao.waybill.ii.search

获取发货地&CP开通状态&账户的使用情况
*/
type CainiaoWaybillIiSearchRequest struct {
    model.Params

    // 物流公司code
    cpCode   string 

}

func NewCainiaoWaybillIiSearchRequest() *CainiaoWaybillIiSearchRequest{
    return &CainiaoWaybillIiSearchRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoWaybillIiSearchRequest) GetApiMethodName() string {
    return "cainiao.waybill.ii.search"
}

func (r CainiaoWaybillIiSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoWaybillIiSearchRequest) SetCpCode(cpCode string) error {
    r.cpCode = cpCode
    r.Set("cp_code", cpCode)
    return nil
}

func (r CainiaoWaybillIiSearchRequest) GetCpCode() string {
    return r.cpCode
}

