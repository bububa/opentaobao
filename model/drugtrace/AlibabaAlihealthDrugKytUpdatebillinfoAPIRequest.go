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
type AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest struct {
    model.Params
    // 企业ID
    _refEntId   string
    // 企业ID
    _entId   string
    // 操作人编码
    _icCode   string
    // 单据ID
    _billId   string
    // 单据类型
    _billType   string
    // 单据编码
    _billCode   string
    // 发货单位ID
    _partnerIdSend   string
    // 收货单信ID
    _partnerIdRecv   string
    // 详情
    _note   string
}

// 初始化AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest对象
func NewAlibabaAlihealthDrugKytUpdatebillinfoRequest() *AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest{
    return &AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.updatebillinfo"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest) GetRefEntId() string {
    return r._refEntId
}
// EntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest) SetEntId(_entId string) error {
    r._entId = _entId
    r.Set("ent_id", _entId)
    return nil
}

// EntId Getter
func (r AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest) GetEntId() string {
    return r._entId
}
// IcCode Setter
// 操作人编码
func (r *AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest) SetIcCode(_icCode string) error {
    r._icCode = _icCode
    r.Set("ic_code", _icCode)
    return nil
}

// IcCode Getter
func (r AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest) GetIcCode() string {
    return r._icCode
}
// BillId Setter
// 单据ID
func (r *AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest) SetBillId(_billId string) error {
    r._billId = _billId
    r.Set("bill_id", _billId)
    return nil
}

// BillId Getter
func (r AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest) GetBillId() string {
    return r._billId
}
// BillType Setter
// 单据类型
func (r *AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest) SetBillType(_billType string) error {
    r._billType = _billType
    r.Set("bill_type", _billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest) GetBillType() string {
    return r._billType
}
// BillCode Setter
// 单据编码
func (r *AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest) SetBillCode(_billCode string) error {
    r._billCode = _billCode
    r.Set("bill_code", _billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest) GetBillCode() string {
    return r._billCode
}
// PartnerIdSend Setter
// 发货单位ID
func (r *AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest) SetPartnerIdSend(_partnerIdSend string) error {
    r._partnerIdSend = _partnerIdSend
    r.Set("partner_id_send", _partnerIdSend)
    return nil
}

// PartnerIdSend Getter
func (r AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest) GetPartnerIdSend() string {
    return r._partnerIdSend
}
// PartnerIdRecv Setter
// 收货单信ID
func (r *AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest) SetPartnerIdRecv(_partnerIdRecv string) error {
    r._partnerIdRecv = _partnerIdRecv
    r.Set("partner_id_recv", _partnerIdRecv)
    return nil
}

// PartnerIdRecv Getter
func (r AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest) GetPartnerIdRecv() string {
    return r._partnerIdRecv
}
// Note Setter
// 详情
func (r *AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest) SetNote(_note string) error {
    r._note = _note
    r.Set("note", _note)
    return nil
}

// Note Getter
func (r AlibabaAlihealthDrugKytUpdatebillinfoAPIRequest) GetNote() string {
    return r._note
}
