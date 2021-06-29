package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询上游出库单明细(带追溯码信息) API请求
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

// 初始化AlibabaAlihealthDrugBillUpbillDetailWithcodeRequest对象
func NewAlibabaAlihealthDrugBillUpbillDetailWithcodeRequest() *AlibabaAlihealthDrugBillUpbillDetailWithcodeRequest{
    return &AlibabaAlihealthDrugBillUpbillDetailWithcodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugBillUpbillDetailWithcodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.bill.upbill.detail.withcode"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugBillUpbillDetailWithcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业id
func (r *AlibabaAlihealthDrugBillUpbillDetailWithcodeRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugBillUpbillDetailWithcodeRequest) GetRefEntId() string {
    return r.refEntId
}
// BillCode Setter
// 单据编码
func (r *AlibabaAlihealthDrugBillUpbillDetailWithcodeRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugBillUpbillDetailWithcodeRequest) GetBillCode() string {
    return r.billCode
}
// FromRefUserId Setter
// 发货企业renEntId
func (r *AlibabaAlihealthDrugBillUpbillDetailWithcodeRequest) SetFromRefUserId(fromRefUserId string) error {
    r.fromRefUserId = fromRefUserId
    r.Set("from_ref_user_id", fromRefUserId)
    return nil
}

// FromRefUserId Getter
func (r AlibabaAlihealthDrugBillUpbillDetailWithcodeRequest) GetFromRefUserId() string {
    return r.fromRefUserId
}
// ToRefUserId Setter
// 收货企业refEntId
func (r *AlibabaAlihealthDrugBillUpbillDetailWithcodeRequest) SetToRefUserId(toRefUserId string) error {
    r.toRefUserId = toRefUserId
    r.Set("to_ref_user_id", toRefUserId)
    return nil
}

// ToRefUserId Getter
func (r AlibabaAlihealthDrugBillUpbillDetailWithcodeRequest) GetToRefUserId() string {
    return r.toRefUserId
}
