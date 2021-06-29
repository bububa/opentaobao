package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过企业名得到唯一标识ref_ent_id及企业ent_id APIRequest
alibaba.alihealth.drugtrace.top.yljg.query.getentinfo

根据企业名称查询ID
*/
type AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoRequest struct {
    model.Params

    // 公司名称(全称)
    entName   string 

}

func NewAlibabaAlihealthDrugtraceTopYljgQueryGetentinfoRequest() *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoRequest{
    return &AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugtrace.top.yljg.query.getentinfo"
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoRequest) SetEntName(entName string) error {
    r.entName = entName
    r.Set("ent_name", entName)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryGetentinfoRequest) GetEntName() string {
    return r.entName
}

