package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过企业名得到唯一标识（ref_ent_id）及企业ID(ent_id) APIRequest
alibaba.alihealth.drug.kyt.dr.getentinfo

根据企业名称查询ID
*/
type AlibabaAlihealthDrugKytDrGetentinfoRequest struct {
    model.Params

    // 公司名称(全称)
    entName   string 

}

func NewAlibabaAlihealthDrugKytDrGetentinfoRequest() *AlibabaAlihealthDrugKytDrGetentinfoRequest{
    return &AlibabaAlihealthDrugKytDrGetentinfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytDrGetentinfoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.getentinfo"
}

func (r AlibabaAlihealthDrugKytDrGetentinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytDrGetentinfoRequest) SetEntName(entName string) error {
    r.entName = entName
    r.Set("ent_name", entName)
    return nil
}

func (r AlibabaAlihealthDrugKytDrGetentinfoRequest) GetEntName() string {
    return r.entName
}

