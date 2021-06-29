package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取企业资质 APIRequest
alibaba.alihealth.drug.kyt.getentlicense

获取企业的资质信息。
*/
type AlibabaAlihealthDrugKytGetentlicenseRequest struct {
    model.Params

    // 企业ID
    refEntId   string 

}

func NewAlibabaAlihealthDrugKytGetentlicenseRequest() *AlibabaAlihealthDrugKytGetentlicenseRequest{
    return &AlibabaAlihealthDrugKytGetentlicenseRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytGetentlicenseRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.getentlicense"
}

func (r AlibabaAlihealthDrugKytGetentlicenseRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytGetentlicenseRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytGetentlicenseRequest) GetRefEntId() string {
    return r.refEntId
}

