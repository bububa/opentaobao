package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询追溯码对应的药品信息（药店） APIRequest
alibaba.alihealth.drug.code.kyt.yd.querycode

此接口针对有码药品，提供可通过追溯码获取该药品的基础信息和生产信息；
核查平台优先过滤非8开头的，长度非20位数字的码信息。
*/
type AlibabaAlihealthDrugCodeKytYdQuerycodeRequest struct {
    model.Params

    // 企业唯一标识
    refEntId   string 

    // 码列表
    codes   []string 

}

func NewAlibabaAlihealthDrugCodeKytYdQuerycodeRequest() *AlibabaAlihealthDrugCodeKytYdQuerycodeRequest{
    return &AlibabaAlihealthDrugCodeKytYdQuerycodeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugCodeKytYdQuerycodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.code.kyt.yd.querycode"
}

func (r AlibabaAlihealthDrugCodeKytYdQuerycodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugCodeKytYdQuerycodeRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugCodeKytYdQuerycodeRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugCodeKytYdQuerycodeRequest) SetCodes(codes []string) error {
    r.codes = codes
    r.Set("codes", codes)
    return nil
}

func (r AlibabaAlihealthDrugCodeKytYdQuerycodeRequest) GetCodes() []string {
    return r.codes
}

