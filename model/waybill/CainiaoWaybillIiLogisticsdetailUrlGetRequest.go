package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子面单物流详情授权url获取 APIRequest
cainiao.waybill.ii.logisticsdetail.url.get

获取电子面单物流详情授权访问的H5 url
*/
type CainiaoWaybillIiLogisticsdetailUrlGetRequest struct {
    model.Params

    // 快递公司编码
    cpCode   string 

    // 电子面单单号
    waybillCode   string 

}

func NewCainiaoWaybillIiLogisticsdetailUrlGetRequest() *CainiaoWaybillIiLogisticsdetailUrlGetRequest{
    return &CainiaoWaybillIiLogisticsdetailUrlGetRequest{
        Params: model.NewParams(),
    }
}

func (r CainiaoWaybillIiLogisticsdetailUrlGetRequest) GetApiMethodName() string {
    return "cainiao.waybill.ii.logisticsdetail.url.get"
}

func (r CainiaoWaybillIiLogisticsdetailUrlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *CainiaoWaybillIiLogisticsdetailUrlGetRequest) SetCpCode(cpCode string) error {
    r.cpCode = cpCode
    r.Set("cp_code", cpCode)
    return nil
}

func (r CainiaoWaybillIiLogisticsdetailUrlGetRequest) GetCpCode() string {
    return r.cpCode
}

func (r *CainiaoWaybillIiLogisticsdetailUrlGetRequest) SetWaybillCode(waybillCode string) error {
    r.waybillCode = waybillCode
    r.Set("waybill_code", waybillCode)
    return nil
}

func (r CainiaoWaybillIiLogisticsdetailUrlGetRequest) GetWaybillCode() string {
    return r.waybillCode
}

