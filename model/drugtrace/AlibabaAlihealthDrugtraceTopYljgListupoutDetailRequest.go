package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
上游出库单单据明细查询 APIRequest
alibaba.alihealth.drugtrace.top.yljg.listupout.detail

查询上游出库单明细(带追溯码信息)
*/
type AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest struct {
    model.Params

    // 企业id
    refEntId   string 

    // 单据编码
    billCode   string 

    // 发货企业renEntId
    fromRefUserId   string 

    // 收货企业refEntId
    toRefUserId   string 

}

func NewAlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest() *AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest{
    return &AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drugtrace.top.yljg.listupout.detail"
}

func (r AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) GetBillCode() string {
    return r.billCode
}

func (r *AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) SetFromRefUserId(fromRefUserId string) error {
    r.fromRefUserId = fromRefUserId
    r.Set("from_ref_user_id", fromRefUserId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) GetFromRefUserId() string {
    return r.fromRefUserId
}

func (r *AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) SetToRefUserId(toRefUserId string) error {
    r.toRefUserId = toRefUserId
    r.Set("to_ref_user_id", toRefUserId)
    return nil
}

func (r AlibabaAlihealthDrugtraceTopYljgListupoutDetailRequest) GetToRefUserId() string {
    return r.toRefUserId
}

