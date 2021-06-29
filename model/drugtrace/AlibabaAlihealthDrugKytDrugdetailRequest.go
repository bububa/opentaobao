package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询药品详细信息 APIRequest
alibaba.alihealth.drug.kyt.drugdetail

查询药品详细信息
*/
type AlibabaAlihealthDrugKytDrugdetailRequest struct {
    model.Params

    // 企业ID
    refEntId   string 

    // 药品ID
    drugEntBaseInfoId   string 

}

func NewAlibabaAlihealthDrugKytDrugdetailRequest() *AlibabaAlihealthDrugKytDrugdetailRequest{
    return &AlibabaAlihealthDrugKytDrugdetailRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytDrugdetailRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.drugdetail"
}

func (r AlibabaAlihealthDrugKytDrugdetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytDrugdetailRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrugdetailRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytDrugdetailRequest) SetDrugEntBaseInfoId(drugEntBaseInfoId string) error {
    r.drugEntBaseInfoId = drugEntBaseInfoId
    r.Set("drug_ent_base_info_id", drugEntBaseInfoId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrugdetailRequest) GetDrugEntBaseInfoId() string {
    return r.drugEntBaseInfoId
}

