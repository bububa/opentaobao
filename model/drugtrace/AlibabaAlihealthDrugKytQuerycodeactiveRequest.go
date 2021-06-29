package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询码是否激活 APIRequest
alibaba.alihealth.drug.kyt.querycodeactive

查询码是否激活
*/
type AlibabaAlihealthDrugKytQuerycodeactiveRequest struct {
    model.Params

    // 企业
    refEntId   string 

    // 码
    codes   []string 

}

func NewAlibabaAlihealthDrugKytQuerycodeactiveRequest() *AlibabaAlihealthDrugKytQuerycodeactiveRequest{
    return &AlibabaAlihealthDrugKytQuerycodeactiveRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytQuerycodeactiveRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.querycodeactive"
}

func (r AlibabaAlihealthDrugKytQuerycodeactiveRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytQuerycodeactiveRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytQuerycodeactiveRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytQuerycodeactiveRequest) SetCodes(codes []string) error {
    r.codes = codes
    r.Set("codes", codes)
    return nil
}

func (r AlibabaAlihealthDrugKytQuerycodeactiveRequest) GetCodes() []string {
    return r.codes
}

