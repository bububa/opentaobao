package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗单据与设备绑定 APIRequest
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

func NewAlibabaAlihealthDrugKytDrAssociateequiRequest() *AlibabaAlihealthDrugKytDrAssociateequiRequest{
    return &AlibabaAlihealthDrugKytDrAssociateequiRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytDrAssociateequiRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.associateequi"
}

func (r AlibabaAlihealthDrugKytDrAssociateequiRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytDrAssociateequiRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrAssociateequiRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytDrAssociateequiRequest) SetBillCodes(billCodes string) error {
    r.billCodes = billCodes
    r.Set("bill_codes", billCodes)
    return nil
}

func (r AlibabaAlihealthDrugKytDrAssociateequiRequest) GetBillCodes() string {
    return r.billCodes
}

func (r *AlibabaAlihealthDrugKytDrAssociateequiRequest) SetVaEquipmentId(vaEquipmentId string) error {
    r.vaEquipmentId = vaEquipmentId
    r.Set("va_equipment_id", vaEquipmentId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrAssociateequiRequest) GetVaEquipmentId() string {
    return r.vaEquipmentId
}

