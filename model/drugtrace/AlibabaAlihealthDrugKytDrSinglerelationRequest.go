package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
多融单码关联关系查询 APIRequest
alibaba.alihealth.drug.kyt.dr.singlerelation

单码关联关系查询
*/
type AlibabaAlihealthDrugKytDrSinglerelationRequest struct {
    model.Params

    // 追溯码
    code   string 

    // 接口调用企业的唯一标识（接口调用者）
    refEntId   string 

    // 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
    desRefEntId   string 

}

func NewAlibabaAlihealthDrugKytDrSinglerelationRequest() *AlibabaAlihealthDrugKytDrSinglerelationRequest{
    return &AlibabaAlihealthDrugKytDrSinglerelationRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytDrSinglerelationRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.singlerelation"
}

func (r AlibabaAlihealthDrugKytDrSinglerelationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytDrSinglerelationRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaAlihealthDrugKytDrSinglerelationRequest) GetCode() string {
    return r.code
}

func (r *AlibabaAlihealthDrugKytDrSinglerelationRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrSinglerelationRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytDrSinglerelationRequest) SetDesRefEntId(desRefEntId string) error {
    r.desRefEntId = desRefEntId
    r.Set("des_ref_ent_id", desRefEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrSinglerelationRequest) GetDesRefEntId() string {
    return r.desRefEntId
}

