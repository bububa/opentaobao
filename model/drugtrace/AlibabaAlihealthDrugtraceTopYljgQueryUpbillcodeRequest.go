package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
通过一个码，查询这个码对应的上游企业出库单的单据号 APIRequest
alibaba.alihealth.drugtrace.top.yljg.query.upbillcode

一个查询上游出库单号的接口。企业在扫码入库时，接口通过扫到的码判定这个码对应的上游企业所属的出库单据号
*/
type AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest struct {
    model.Params

    // 追溯码
    code   string 

    // 企业ID （一般为要查询单据的收货企业）
    refEntId   string 

}

func NewAlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest() *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest{
    return &AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugtrace.top.yljg.query.upbillcode"
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest) GetCode() string {
    return r.code
}

func (r *AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryUpbillcodeRequest) GetRefEntId() string {
    return r.refEntId
}

