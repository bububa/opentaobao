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
type AlibabaAlihealthDrugKytYyQuerysubcodesRequest struct {
    model.Params
    // 接口调用企业的唯一标识（接口调用者）
    refEntId   string
    // 码
    codes   []string
}

// 初始化AlibabaAlihealthDrugKytYyQuerysubcodesRequest对象
func NewAlibabaAlihealthDrugKytYyQuerysubcodesRequest() *AlibabaAlihealthDrugKytYyQuerysubcodesRequest{
    return &AlibabaAlihealthDrugKytYyQuerysubcodesRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytYyQuerysubcodesRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.yy.querysubcodes"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytYyQuerysubcodesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugKytYyQuerysubcodesRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytYyQuerysubcodesRequest) GetRefEntId() string {
    return r.refEntId
}
// Codes Setter
// 码
func (r *AlibabaAlihealthDrugKytYyQuerysubcodesRequest) SetCodes(codes []string) error {
    r.codes = codes
    r.Set("codes", codes)
    return nil
}

// Codes Getter
func (r AlibabaAlihealthDrugKytYyQuerysubcodesRequest) GetCodes() []string {
    return r.codes
}
