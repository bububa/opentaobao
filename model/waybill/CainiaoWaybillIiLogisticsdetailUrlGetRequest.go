package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
电子面单物流详情授权url获取 API请求
cainiao.waybill.ii.logisticsdetail.url.get

获取电子面单物流详情授权访问的H5 url
*/
type CainiaoWaybillIiLogisticsdetailUrlGetRequest struct {
    model.Params
    // 快递公司编码
    _cpCode   string
    // 电子面单单号
    _waybillCode   string
}

// 初始化CainiaoWaybillIiLogisticsdetailUrlGetRequest对象
func NewCainiaoWaybillIiLogisticsdetailUrlGetRequest() *CainiaoWaybillIiLogisticsdetailUrlGetRequest{
    return &CainiaoWaybillIiLogisticsdetailUrlGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoWaybillIiLogisticsdetailUrlGetRequest) GetApiMethodName() string {
    return "cainiao.waybill.ii.logisticsdetail.url.get"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoWaybillIiLogisticsdetailUrlGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CpCode Setter
// 快递公司编码
func (r *CainiaoWaybillIiLogisticsdetailUrlGetRequest) SetCpCode(_cpCode string) error {
    r._cpCode = _cpCode
    r.Set("cp_code", _cpCode)
    return nil
}

// CpCode Getter
func (r CainiaoWaybillIiLogisticsdetailUrlGetRequest) GetCpCode() string {
    return r._cpCode
}
// WaybillCode Setter
// 电子面单单号
func (r *CainiaoWaybillIiLogisticsdetailUrlGetRequest) SetWaybillCode(_waybillCode string) error {
    r._waybillCode = _waybillCode
    r.Set("waybill_code", _waybillCode)
    return nil
}

// WaybillCode Getter
func (r CainiaoWaybillIiLogisticsdetailUrlGetRequest) GetWaybillCode() string {
    return r._waybillCode
}
