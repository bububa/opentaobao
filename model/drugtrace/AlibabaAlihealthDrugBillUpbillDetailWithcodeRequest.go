package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询上游出库单明细(带追溯码信息) APIRequest
alibaba.alihealth.drug.bill.upbill.detail.withcode

查询上游出库单明细(带追溯码信息)
*/
type AlibabaAlihealthDrugBillUpbillDetailWithcodeRequest struct {
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

func NewAlibabaAlihealthDrugBillUpbillDetailWithcodeRequest() *AlibabaAlihealthDrugBillUpbillDetailWithcodeRequest{
    return &AlibabaAlihealthDrugBillUpbillDetailWithcodeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugBillUpbillDetailWithcodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.bill.upbill.detail.withcode"
}

func (r AlibabaAlihealthDrugBillUpbillDetailWithcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugBillUpbillDetailWithcodeRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugBillUpbillDetailWithcodeRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugBillUpbillDetailWithcodeRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

func (r AlibabaAlihealthDrugBillUpbillDetailWithcodeRequest) GetBillCode() string {
    return r.billCode
}

func (r *AlibabaAlihealthDrugBillUpbillDetailWithcodeRequest) SetFromRefUserId(fromRefUserId string) error {
    r.fromRefUserId = fromRefUserId
    r.Set("from_ref_user_id", fromRefUserId)
    return nil
}

func (r AlibabaAlihealthDrugBillUpbillDetailWithcodeRequest) GetFromRefUserId() string {
    return r.fromRefUserId
}

func (r *AlibabaAlihealthDrugBillUpbillDetailWithcodeRequest) SetToRefUserId(toRefUserId string) error {
    r.toRefUserId = toRefUserId
    r.Set("to_ref_user_id", toRefUserId)
    return nil
}

func (r AlibabaAlihealthDrugBillUpbillDetailWithcodeRequest) GetToRefUserId() string {
    return r.toRefUserId
}

