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
type AlibabaAlihealthDrugKytDrAssociateequiAPIRequest struct {
    model.Params
    // 企业refentid
    _refEntId   string
    // 单据编号，多个用逗号分隔
    _billCodes   string
    // 设备ID
    _vaEquipmentId   string
}

// 初始化AlibabaAlihealthDrugKytDrAssociateequiAPIRequest对象
func NewAlibabaAlihealthDrugKytDrAssociateequiRequest() *AlibabaAlihealthDrugKytDrAssociateequiAPIRequest{
    return &AlibabaAlihealthDrugKytDrAssociateequiAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrAssociateequiAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.associateequi"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrAssociateequiAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业refentid
func (r *AlibabaAlihealthDrugKytDrAssociateequiAPIRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytDrAssociateequiAPIRequest) GetRefEntId() string {
    return r._refEntId
}
// BillCodes Setter
// 单据编号，多个用逗号分隔
func (r *AlibabaAlihealthDrugKytDrAssociateequiAPIRequest) SetBillCodes(_billCodes string) error {
    r._billCodes = _billCodes
    r.Set("bill_codes", _billCodes)
    return nil
}

// BillCodes Getter
func (r AlibabaAlihealthDrugKytDrAssociateequiAPIRequest) GetBillCodes() string {
    return r._billCodes
}
// VaEquipmentId Setter
// 设备ID
func (r *AlibabaAlihealthDrugKytDrAssociateequiAPIRequest) SetVaEquipmentId(_vaEquipmentId string) error {
    r._vaEquipmentId = _vaEquipmentId
    r.Set("va_equipment_id", _vaEquipmentId)
    return nil
}

// VaEquipmentId Getter
func (r AlibabaAlihealthDrugKytDrAssociateequiAPIRequest) GetVaEquipmentId() string {
    return r._vaEquipmentId
}
