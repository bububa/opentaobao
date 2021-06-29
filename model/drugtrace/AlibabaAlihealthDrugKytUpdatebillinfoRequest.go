package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售端平台单据更新 APIRequest
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

func NewAlibabaAlihealthDrugKytUpdatebillinfoRequest() *AlibabaAlihealthDrugKytUpdatebillinfoRequest{
    return &AlibabaAlihealthDrugKytUpdatebillinfoRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytUpdatebillinfoRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.updatebillinfo"
}

func (r AlibabaAlihealthDrugKytUpdatebillinfoRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytUpdatebillinfoRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytUpdatebillinfoRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytUpdatebillinfoRequest) SetEntId(entId string) error {
    r.entId = entId
    r.Set("ent_id", entId)
    return nil
}

func (r AlibabaAlihealthDrugKytUpdatebillinfoRequest) GetEntId() string {
    return r.entId
}

func (r *AlibabaAlihealthDrugKytUpdatebillinfoRequest) SetIcCode(icCode string) error {
    r.icCode = icCode
    r.Set("ic_code", icCode)
    return nil
}

func (r AlibabaAlihealthDrugKytUpdatebillinfoRequest) GetIcCode() string {
    return r.icCode
}

func (r *AlibabaAlihealthDrugKytUpdatebillinfoRequest) SetBillId(billId string) error {
    r.billId = billId
    r.Set("bill_id", billId)
    return nil
}

func (r AlibabaAlihealthDrugKytUpdatebillinfoRequest) GetBillId() string {
    return r.billId
}

func (r *AlibabaAlihealthDrugKytUpdatebillinfoRequest) SetBillType(billType string) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

func (r AlibabaAlihealthDrugKytUpdatebillinfoRequest) GetBillType() string {
    return r.billType
}

func (r *AlibabaAlihealthDrugKytUpdatebillinfoRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

func (r AlibabaAlihealthDrugKytUpdatebillinfoRequest) GetBillCode() string {
    return r.billCode
}

func (r *AlibabaAlihealthDrugKytUpdatebillinfoRequest) SetPartnerIdSend(partnerIdSend string) error {
    r.partnerIdSend = partnerIdSend
    r.Set("partner_id_send", partnerIdSend)
    return nil
}

func (r AlibabaAlihealthDrugKytUpdatebillinfoRequest) GetPartnerIdSend() string {
    return r.partnerIdSend
}

func (r *AlibabaAlihealthDrugKytUpdatebillinfoRequest) SetPartnerIdRecv(partnerIdRecv string) error {
    r.partnerIdRecv = partnerIdRecv
    r.Set("partner_id_recv", partnerIdRecv)
    return nil
}

func (r AlibabaAlihealthDrugKytUpdatebillinfoRequest) GetPartnerIdRecv() string {
    return r.partnerIdRecv
}

func (r *AlibabaAlihealthDrugKytUpdatebillinfoRequest) SetNote(note string) error {
    r.note = note
    r.Set("note", note)
    return nil
}

func (r AlibabaAlihealthDrugKytUpdatebillinfoRequest) GetNote() string {
    return r.note
}

