package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询一个码的所有子码 API请求
alibaba.alihealth.drug.kyt.yy.querysubcodes

单码的了码查询
*/
type AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest struct {
    model.Params
    // 接口调用企业的唯一标识（接口调用者）
    _refEntId   string
    // 码
    _codes   []string
}

// 初始化AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest对象
func NewAlibabaAlihealthDrugKytYyQuerysubcodesRequest() *AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest{
    return &AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.yy.querysubcodes"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest) GetRefEntId() string {
    return r._refEntId
}
// Codes Setter
// 码
func (r *AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest) SetCodes(_codes []string) error {
    r._codes = _codes
    r.Set("codes", _codes)
    return nil
}

// Codes Getter
func (r AlibabaAlihealthDrugKytYyQuerysubcodesAPIRequest) GetCodes() []string {
    return r._codes
}
