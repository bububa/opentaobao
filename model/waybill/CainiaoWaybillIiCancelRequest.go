package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家取消获取的电子面单号 API请求
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

// 初始化CainiaoWaybillIiCancelRequest对象
func NewCainiaoWaybillIiCancelRequest() *CainiaoWaybillIiCancelRequest{
    return &CainiaoWaybillIiCancelRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoWaybillIiCancelRequest) GetApiMethodName() string {
    return "cainiao.waybill.ii.cancel"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoWaybillIiCancelRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CpCode Setter
// 快递公司code
func (r *CainiaoWaybillIiCancelRequest) SetCpCode(cpCode string) error {
    r.cpCode = cpCode
    r.Set("cp_code", cpCode)
    return nil
}

// CpCode Getter
func (r CainiaoWaybillIiCancelRequest) GetCpCode() string {
    return r.cpCode
}
// WaybillCode Setter
// 电子面单号
func (r *CainiaoWaybillIiCancelRequest) SetWaybillCode(waybillCode string) error {
    r.waybillCode = waybillCode
    r.Set("waybill_code", waybillCode)
    return nil
}

// WaybillCode Getter
func (r CainiaoWaybillIiCancelRequest) GetWaybillCode() string {
    return r.waybillCode
}
