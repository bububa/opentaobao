package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询一个码的所有子码 APIRequest
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

func NewAlibabaAlihealthDrugKytYyQuerysubcodesRequest() *AlibabaAlihealthDrugKytYyQuerysubcodesRequest{
    return &AlibabaAlihealthDrugKytYyQuerysubcodesRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytYyQuerysubcodesRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.yy.querysubcodes"
}

func (r AlibabaAlihealthDrugKytYyQuerysubcodesRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytYyQuerysubcodesRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytYyQuerysubcodesRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytYyQuerysubcodesRequest) SetCodes(codes []string) error {
    r.codes = codes
    r.Set("codes", codes)
    return nil
}

func (r AlibabaAlihealthDrugKytYyQuerysubcodesRequest) GetCodes() []string {
    return r.codes
}

