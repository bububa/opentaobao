package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
多融根据码查询码信息 APIRequest
alibaba.alihealth.drug.code.kyt.dr.querycode

服务描述
此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息。
核查平台优先过滤非8开头的，长度非20位数字的码信息。
*/
type AlibabaAlihealthDrugCodeKytDrQuerycodeRequest struct {
    model.Params

    // 企业唯一标识（或appkey）
    refEntId   string 

    // 码列表
    codes   []string 

}

func NewAlibabaAlihealthDrugCodeKytDrQuerycodeRequest() *AlibabaAlihealthDrugCodeKytDrQuerycodeRequest{
    return &AlibabaAlihealthDrugCodeKytDrQuerycodeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugCodeKytDrQuerycodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.code.kyt.dr.querycode"
}

func (r AlibabaAlihealthDrugCodeKytDrQuerycodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugCodeKytDrQuerycodeRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugCodeKytDrQuerycodeRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugCodeKytDrQuerycodeRequest) SetCodes(codes []string) error {
    r.codes = codes
    r.Set("codes", codes)
    return nil
}

func (r AlibabaAlihealthDrugCodeKytDrQuerycodeRequest) GetCodes() []string {
    return r.codes
}

