package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单码关联关系查询 API请求
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

// 初始化AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest对象
func NewAlibabaAlihealthDrugtraceTopLsydQueryRelationRequest() *AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest{
    return &AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugtrace.top.lsyd.query.relation"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest) GetCode() string {
    return r.code
}
// RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest) GetRefEntId() string {
    return r.refEntId
}
// DesRefEntId Setter
// 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
func (r *AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest) SetDesRefEntId(desRefEntId string) error {
    r.desRefEntId = desRefEntId
    r.Set("des_ref_ent_id", desRefEntId)
    return nil
}

// DesRefEntId Getter
func (r AlibabaAlihealthDrugtraceTopLsydQueryRelationRequest) GetDesRefEntId() string {
    return r.desRefEntId
}
