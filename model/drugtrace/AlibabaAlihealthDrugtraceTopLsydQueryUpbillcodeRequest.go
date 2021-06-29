package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过一个码，查询这个码对应的上游企业出库单的单据号 APIRequest
alibaba.alihealth.drugtrace.top.lsyd.query.upbillcode

一个查询上游出库单号的接口。企业在扫码入库时，接口通过扫到的码判定这个码对应的上游企业所属的出库单据号
*/
type AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeRequest struct {
    model.Params

    // 追溯码
    code   string 

    // 企业REF_ENT_ID （当前企业的唯一标识）
    refEntId   string 

}

func NewAlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeRequest() *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeRequest{
    return &AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugtrace.top.lsyd.query.upbillcode"
}

func (r AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeRequest) GetCode() string {
    return r.code
}

func (r *AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopLsydQueryUpbillcodeRequest) GetRefEntId() string {
    return r.refEntId
}

