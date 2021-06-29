package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】 APIRequest
alibaba.alihealth.drug.kyt.getentinfo

根据企业名称查询企业唯一标识【ref_ent_id】和企业ID【ent_id】
*/
type AlibabaAlihealthDrugKytGetentinfoRequest struct {
    model.Params

    // 公司名称
    entName   string 

}

func NewAlibabaAlihealthDrugKytGetentinfoRequest() *AlibabaAlihealthDrugKytGetentinfoRequest{
    return &AlibabaAlihealthDrugKytGetentinfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytGetentinfoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.getentinfo"
}

func (r AlibabaAlihealthDrugKytGetentinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytGetentinfoRequest) SetEntName(entName string) error {
    r.entName = entName
    r.Set("ent_name", entName)
    return nil
}

func (r AlibabaAlihealthDrugKytGetentinfoRequest) GetEntName() string {
    return r.entName
}

