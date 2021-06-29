package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据码获取基本和单据信息 APIRequest
alibaba.alihealth.drug.kyt.getcodebillinfo

根据码信息获取基本信息和单据信息
*/
type AlibabaAlihealthDrugKytGetcodebillinfoRequest struct {
    model.Params

    // 企业ID
    refEntId   string 

    // 码
    code   string 

}

func NewAlibabaAlihealthDrugKytGetcodebillinfoRequest() *AlibabaAlihealthDrugKytGetcodebillinfoRequest{
    return &AlibabaAlihealthDrugKytGetcodebillinfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytGetcodebillinfoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.getcodebillinfo"
}

func (r AlibabaAlihealthDrugKytGetcodebillinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytGetcodebillinfoRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytGetcodebillinfoRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytGetcodebillinfoRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaAlihealthDrugKytGetcodebillinfoRequest) GetCode() string {
    return r.code
}

