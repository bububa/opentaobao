package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取药品资质信息 APIRequest
alibaba.alihealth.drug.kyt.getdruglicense

获取药品的资质信息。
*/
type AlibabaAlihealthDrugKytGetdruglicenseRequest struct {
    model.Params

    // 企业ID
    refEntId   string 

    // 药品ID
    drugId   string 

}

func NewAlibabaAlihealthDrugKytGetdruglicenseRequest() *AlibabaAlihealthDrugKytGetdruglicenseRequest{
    return &AlibabaAlihealthDrugKytGetdruglicenseRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytGetdruglicenseRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.getdruglicense"
}

func (r AlibabaAlihealthDrugKytGetdruglicenseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytGetdruglicenseRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytGetdruglicenseRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytGetdruglicenseRequest) SetDrugId(drugId string) error {
    r.drugId = drugId
    r.Set("drug_id", drugId)
    return nil
}

func (r AlibabaAlihealthDrugKytGetdruglicenseRequest) GetDrugId() string {
    return r.drugId
}

