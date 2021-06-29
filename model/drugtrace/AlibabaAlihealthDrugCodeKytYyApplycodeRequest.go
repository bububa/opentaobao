package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
医院药品子码申请接口 APIRequest
alibaba.alihealth.drug.code.kyt.yy.applycode

根据父码及所属企业ID生成子码信息
*/
type AlibabaAlihealthDrugCodeKytYyApplycodeRequest struct {
    model.Params

    // 企业ID（ref_ent_id)
    refEntId   string 

    // 父码
    code   string 

    // 申请数量
    amount   int64 

}

func NewAlibabaAlihealthDrugCodeKytYyApplycodeRequest() *AlibabaAlihealthDrugCodeKytYyApplycodeRequest{
    return &AlibabaAlihealthDrugCodeKytYyApplycodeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugCodeKytYyApplycodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.code.kyt.yy.applycode"
}

func (r AlibabaAlihealthDrugCodeKytYyApplycodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugCodeKytYyApplycodeRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugCodeKytYyApplycodeRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugCodeKytYyApplycodeRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaAlihealthDrugCodeKytYyApplycodeRequest) GetCode() string {
    return r.code
}

func (r *AlibabaAlihealthDrugCodeKytYyApplycodeRequest) SetAmount(amount int64) error {
    r.amount = amount
    r.Set("amount", amount)
    return nil
}

func (r AlibabaAlihealthDrugCodeKytYyApplycodeRequest) GetAmount() int64 {
    return r.amount
}

