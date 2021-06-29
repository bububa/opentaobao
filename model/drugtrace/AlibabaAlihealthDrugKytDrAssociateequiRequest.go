package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗单据与设备绑定 API请求
alibaba.alihealth.drug.kyt.dr.associateequi

疫苗单据与设备绑定
*/
type AlibabaAlihealthDrugKytDrAssociateequiRequest struct {
    model.Params
    // 企业refentid
    refEntId   string
    // 单据编号，多个用逗号分隔
    billCodes   string
    // 设备ID
    vaEquipmentId   string
}

// 初始化AlibabaAlihealthDrugKytDrAssociateequiRequest对象
func NewAlibabaAlihealthDrugKytDrAssociateequiRequest() *AlibabaAlihealthDrugKytDrAssociateequiRequest{
    return &AlibabaAlihealthDrugKytDrAssociateequiRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrAssociateequiRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.associateequi"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrAssociateequiRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业refentid
func (r *AlibabaAlihealthDrugKytDrAssociateequiRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytDrAssociateequiRequest) GetRefEntId() string {
    return r.refEntId
}
// BillCodes Setter
// 单据编号，多个用逗号分隔
func (r *AlibabaAlihealthDrugKytDrAssociateequiRequest) SetBillCodes(billCodes string) error {
    r.billCodes = billCodes
    r.Set("bill_codes", billCodes)
    return nil
}

// BillCodes Getter
func (r AlibabaAlihealthDrugKytDrAssociateequiRequest) GetBillCodes() string {
    return r.billCodes
}
// VaEquipmentId Setter
// 设备ID
func (r *AlibabaAlihealthDrugKytDrAssociateequiRequest) SetVaEquipmentId(vaEquipmentId string) error {
    r.vaEquipmentId = vaEquipmentId
    r.Set("va_equipment_id", vaEquipmentId)
    return nil
}

// VaEquipmentId Getter
func (r AlibabaAlihealthDrugKytDrAssociateequiRequest) GetVaEquipmentId() string {
    return r.vaEquipmentId
}
