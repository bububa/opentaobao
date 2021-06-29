package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售端平台单据更新 API请求
alibaba.alihealth.drug.kyt.updatebillinfo

零售端平台单据更新
*/
type AlibabaAlihealthDrugKytUpdatebillinfoRequest struct {
    model.Params
    // 企业ID
    refEntId   string
    // 企业ID
    entId   string
    // 操作人编码
    icCode   string
    // 单据ID
    billId   string
    // 单据类型
    billType   string
    // 单据编码
    billCode   string
    // 发货单位ID
    partnerIdSend   string
    // 收货单信ID
    partnerIdRecv   string
    // 详情
    note   string
}

// 初始化AlibabaAlihealthDrugKytUpdatebillinfoRequest对象
func NewAlibabaAlihealthDrugKytUpdatebillinfoRequest() *AlibabaAlihealthDrugKytUpdatebillinfoRequest{
    return &AlibabaAlihealthDrugKytUpdatebillinfoRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytUpdatebillinfoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.updatebillinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytUpdatebillinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytUpdatebillinfoRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytUpdatebillinfoRequest) GetRefEntId() string {
    return r.refEntId
}
// EntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytUpdatebillinfoRequest) SetEntId(entId string) error {
    r.entId = entId
    r.Set("ent_id", entId)
    return nil
}

// EntId Getter
func (r AlibabaAlihealthDrugKytUpdatebillinfoRequest) GetEntId() string {
    return r.entId
}
// IcCode Setter
// 操作人编码
func (r *AlibabaAlihealthDrugKytUpdatebillinfoRequest) SetIcCode(icCode string) error {
    r.icCode = icCode
    r.Set("ic_code", icCode)
    return nil
}

// IcCode Getter
func (r AlibabaAlihealthDrugKytUpdatebillinfoRequest) GetIcCode() string {
    return r.icCode
}
// BillId Setter
// 单据ID
func (r *AlibabaAlihealthDrugKytUpdatebillinfoRequest) SetBillId(billId string) error {
    r.billId = billId
    r.Set("bill_id", billId)
    return nil
}

// BillId Getter
func (r AlibabaAlihealthDrugKytUpdatebillinfoRequest) GetBillId() string {
    return r.billId
}
// BillType Setter
// 单据类型
func (r *AlibabaAlihealthDrugKytUpdatebillinfoRequest) SetBillType(billType string) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytUpdatebillinfoRequest) GetBillType() string {
    return r.billType
}
// BillCode Setter
// 单据编码
func (r *AlibabaAlihealthDrugKytUpdatebillinfoRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytUpdatebillinfoRequest) GetBillCode() string {
    return r.billCode
}
// PartnerIdSend Setter
// 发货单位ID
func (r *AlibabaAlihealthDrugKytUpdatebillinfoRequest) SetPartnerIdSend(partnerIdSend string) error {
    r.partnerIdSend = partnerIdSend
    r.Set("partner_id_send", partnerIdSend)
    return nil
}

// PartnerIdSend Getter
func (r AlibabaAlihealthDrugKytUpdatebillinfoRequest) GetPartnerIdSend() string {
    return r.partnerIdSend
}
// PartnerIdRecv Setter
// 收货单信ID
func (r *AlibabaAlihealthDrugKytUpdatebillinfoRequest) SetPartnerIdRecv(partnerIdRecv string) error {
    r.partnerIdRecv = partnerIdRecv
    r.Set("partner_id_recv", partnerIdRecv)
    return nil
}

// PartnerIdRecv Getter
func (r AlibabaAlihealthDrugKytUpdatebillinfoRequest) GetPartnerIdRecv() string {
    return r.partnerIdRecv
}
// Note Setter
// 详情
func (r *AlibabaAlihealthDrugKytUpdatebillinfoRequest) SetNote(note string) error {
    r.note = note
    r.Set("note", note)
    return nil
}

// Note Getter
func (r AlibabaAlihealthDrugKytUpdatebillinfoRequest) GetNote() string {
    return r.note
}
