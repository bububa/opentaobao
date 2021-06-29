package waybill

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
商家查询物流商产品类型接口 API请求
cainiao.waybill.ii.product

商家可以查询物流商的产品类型和服务能力。
*/
type CainiaoWaybillIiProductRequest struct {
    model.Params
    // 快递公司code
    cpCode   string
}

// 初始化CainiaoWaybillIiProductRequest对象
func NewCainiaoWaybillIiProductRequest() *CainiaoWaybillIiProductRequest{
    return &CainiaoWaybillIiProductRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r CainiaoWaybillIiProductRequest) GetApiMethodName() string {
    return "cainiao.waybill.ii.product"
}

// IRequest interface 方法, 获取API参数
func (r CainiaoWaybillIiProductRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// CpCode Setter
// 快递公司code
func (r *CainiaoWaybillIiProductRequest) SetCpCode(cpCode string) error {
    r.cpCode = cpCode
    r.Set("cp_code", cpCode)
    return nil
}

// CpCode Getter
func (r CainiaoWaybillIiProductRequest) GetCpCode() string {
    return r.cpCode
}
