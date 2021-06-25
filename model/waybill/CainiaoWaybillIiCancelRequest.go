package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家取消获取的电子面单号 APIRequest
cainiao.waybill.ii.cancel

面单号有误需要取消的时候，调用该接口取消获取的电子面单。
*/
type CainiaoWaybillIiCancelRequest struct {
    model.Params

    // 快递公司code
    cpCode   string 

    // 电子面单号
    waybillCode   string 

}

func NewCainiaoWaybillIiCancelRequest() *CainiaoWaybillIiCancelRequest{
    return &CainiaoWaybillIiCancelRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoWaybillIiCancelRequest) GetApiMethodName() string {
    return "cainiao.waybill.ii.cancel"
}

func (r CainiaoWaybillIiCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoWaybillIiCancelRequest) SetCpCode(cpCode string) error {
    r.cpCode = cpCode
    r.Set("cp_code", cpCode)
    return nil
}

func (r CainiaoWaybillIiCancelRequest) GetCpCode() string {
    return r.cpCode
}

func (r *CainiaoWaybillIiCancelRequest) SetWaybillCode(waybillCode string) error {
    r.waybillCode = waybillCode
    r.Set("waybill_code", waybillCode)
    return nil
}

func (r CainiaoWaybillIiCancelRequest) GetWaybillCode() string {
    return r.waybillCode
}

