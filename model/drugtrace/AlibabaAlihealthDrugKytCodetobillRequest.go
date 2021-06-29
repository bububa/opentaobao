package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过追溯码查单据 APIRequest
alibaba.alihealth.drug.kyt.codetobill

通过追溯码查单据
*/
type AlibabaAlihealthDrugKytCodetobillRequest struct {
    model.Params

    // 企业ID
    refEntId   string 

    // 追溯码
    code   string 

}

func NewAlibabaAlihealthDrugKytCodetobillRequest() *AlibabaAlihealthDrugKytCodetobillRequest{
    return &AlibabaAlihealthDrugKytCodetobillRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytCodetobillRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.codetobill"
}

func (r AlibabaAlihealthDrugKytCodetobillRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytCodetobillRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytCodetobillRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytCodetobillRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaAlihealthDrugKytCodetobillRequest) GetCode() string {
    return r.code
}

