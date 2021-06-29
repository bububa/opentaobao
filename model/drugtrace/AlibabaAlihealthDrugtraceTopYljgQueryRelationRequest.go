package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
单码关联关系查询 API请求
alibaba.alihealth.drugtrace.top.yljg.query.relation

单码关联关系查询
*/
type AlibabaAlihealthDrugtraceTopYljgQueryRelationRequest struct {
    model.Params
    // 追溯码
    _code   string
    // 接口调用企业的唯一标识（接口调用者）
    _refEntId   string
    // 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
    _desRefEntId   string
}

// 初始化AlibabaAlihealthDrugtraceTopYljgQueryRelationRequest对象
func NewAlibabaAlihealthDrugtraceTopYljgQueryRelationRequest() *AlibabaAlihealthDrugtraceTopYljgQueryRelationRequest{
    return &AlibabaAlihealthDrugtraceTopYljgQueryRelationRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugtraceTopYljgQueryRelationRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugtrace.top.yljg.query.relation"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugtraceTopYljgQueryRelationRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Code Setter
// 追溯码
func (r *AlibabaAlihealthDrugtraceTopYljgQueryRelationRequest) SetCode(_code string) error {
    r._code = _code
    r.Set("code", _code)
    return nil
}

// Code Getter
func (r AlibabaAlihealthDrugtraceTopYljgQueryRelationRequest) GetCode() string {
    return r._code
}
// RefEntId Setter
// 接口调用企业的唯一标识（接口调用者）
func (r *AlibabaAlihealthDrugtraceTopYljgQueryRelationRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugtraceTopYljgQueryRelationRequest) GetRefEntId() string {
    return r._refEntId
}
// DesRefEntId Setter
// 目标企业唯一标识（为哪个企业查询，一般与入参ref_ent_id一样）
func (r *AlibabaAlihealthDrugtraceTopYljgQueryRelationRequest) SetDesRefEntId(_desRefEntId string) error {
    r._desRefEntId = _desRefEntId
    r.Set("des_ref_ent_id", _desRefEntId)
    return nil
}

// DesRefEntId Getter
func (r AlibabaAlihealthDrugtraceTopYljgQueryRelationRequest) GetDesRefEntId() string {
    return r._desRefEntId
}
