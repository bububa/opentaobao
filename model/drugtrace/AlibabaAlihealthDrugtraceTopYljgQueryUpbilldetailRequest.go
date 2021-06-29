package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据单据号查询单据的详情信息【注意：查询的是本企业的单据】 APIRequest
alibaba.alihealth.drugtrace.top.yljg.query.upbilldetail

根据单据号查询单据的详情信息【注意：这个接口查询的是本企业的单据，如果是查询上游的单据明细信息，使用xxxxxxx.listupout.detail接口】。
*/
type AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest struct {
    model.Params

    // 单据号码
    billCode   string 

    // 本企业refEntId
    refEntId   string 

}

func NewAlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest() *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest{
    return &AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugtrace.top.yljg.query.upbilldetail"
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest) GetBillCode() string {
    return r.billCode
}

func (r *AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgQueryUpbilldetailRequest) GetRefEntId() string {
    return r.refEntId
}

