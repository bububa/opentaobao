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
    _cpCode   string
    // 电子面单号
    _waybillCode   string
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
func (r *CainiaoWaybillIiCancelRequest) SetCpCode(_cpCode string) error {
    r._cpCode = _cpCode
    r.Set("cp_code", _cpCode)
    return nil
}

// CpCode Getter
func (r CainiaoWaybillIiCancelRequest) GetCpCode() string {
    return r._cpCode
}
// WaybillCode Setter
// 电子面单号
func (r *CainiaoWaybillIiCancelRequest) SetWaybillCode(_waybillCode string) error {
    r._waybillCode = _waybillCode
    r.Set("waybill_code", _waybillCode)
    return nil
}

// WaybillCode Getter
func (r CainiaoWaybillIiCancelRequest) GetWaybillCode() string {
    return r._waybillCode
}
