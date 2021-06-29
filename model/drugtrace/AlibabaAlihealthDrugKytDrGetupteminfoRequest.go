package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
获取上游温度信息（疫苗） APIRequest
alibaba.alihealth.drug.kyt.dr.getupteminfo

根据追溯码及企业ID获取上游运输及存储温度信息（疫苗）
*/
type AlibabaAlihealthDrugKytDrGetupteminfoRequest struct {
    model.Params

    // 追溯码
    code   string 

    // 企业ID
    refEntId   string 

}

func NewAlibabaAlihealthDrugKytDrGetupteminfoRequest() *AlibabaAlihealthDrugKytDrGetupteminfoRequest{
    return &AlibabaAlihealthDrugKytDrGetupteminfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytDrGetupteminfoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.getupteminfo"
}

func (r AlibabaAlihealthDrugKytDrGetupteminfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytDrGetupteminfoRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaAlihealthDrugKytDrGetupteminfoRequest) GetCode() string {
    return r.code
}

func (r *AlibabaAlihealthDrugKytDrGetupteminfoRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrGetupteminfoRequest) GetRefEntId() string {
    return r.refEntId
}

