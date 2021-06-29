package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据码查询码信息 APIRequest
alibaba.alihealth.drugtrace.top.lsyd.query.codedetail

服务描述
此接口，针对有码药品，提供可通过追溯码获取该药品的基础信息和生产状况信息。
核查平台优先过滤非8开头的，长度非20位数字的码信息。
*/
type AlibabaAlihealthDrugtraceTopLsydQueryCodedetailRequest struct {
    model.Params

    // 企业唯一标识（或appkey）
    refEntId   string 

    // 码列表
    codes   []string 

}

func NewAlibabaAlihealthDrugtraceTopLsydQueryCodedetailRequest() *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailRequest{
    return &AlibabaAlihealthDrugtraceTopLsydQueryCodedetailRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugtraceTopLsydQueryCodedetailRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugtrace.top.lsyd.query.codedetail"
}

func (r AlibabaAlihealthDrugtraceTopLsydQueryCodedetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopLsydQueryCodedetailRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugtraceTopLsydQueryCodedetailRequest) SetCodes(codes []string) error {
    r.codes = codes
    r.Set("codes", codes)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopLsydQueryCodedetailRequest) GetCodes() []string {
    return r.codes
}

