package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过企业名得到唯一标识（ref_ent_id）及企业ID(ent_id) APIRequest
alibaba.alihealth.drug.kyt.va.getentinfo

根据企业名称查询企业唯一标识（ref_ent_id）及企业ID(ent_id)
*/
type AlibabaAlihealthDrugKytVaGetentinfoRequest struct {
    model.Params

    // 公司名称(全称)
    entName   string 

}

func NewAlibabaAlihealthDrugKytVaGetentinfoRequest() *AlibabaAlihealthDrugKytVaGetentinfoRequest{
    return &AlibabaAlihealthDrugKytVaGetentinfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytVaGetentinfoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.va.getentinfo"
}

func (r AlibabaAlihealthDrugKytVaGetentinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytVaGetentinfoRequest) SetEntName(entName string) error {
    r.entName = entName
    r.Set("ent_name", entName)
    return nil
}

func (r AlibabaAlihealthDrugKytVaGetentinfoRequest) GetEntName() string {
    return r.entName
}

