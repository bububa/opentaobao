package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单码关联关系查询 APIRequest
alibaba.alihealth.drugtrace.top.lsyd.query.relation

单码关联关系查询
*/
type AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest struct {
    model.Params

    // 追溯码
    code   string 

    // 接口调用企业的唯一标识（接口调用者）
    refEntId   string 

    // 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
    desRefEntId   string 

}

func NewAlibabaAlihealthDrugtraceTopLsydQueryRelationRequest() *AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest{
    return &AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugtrace.top.lsyd.query.relation"
}

func (r AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest) GetCode() string {
    return r.code
}

func (r *AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest) SetDesRefEntId(desRefEntId string) error {
    r.desRefEntId = desRefEntId
    r.Set("des_ref_ent_id", desRefEntId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest) GetDesRefEntId() string {
    return r.desRefEntId
}

