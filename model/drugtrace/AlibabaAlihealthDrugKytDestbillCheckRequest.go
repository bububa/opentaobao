package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
直调审批 APIRequest
alibaba.alihealth.drug.kyt.destbill.check

为药企提供直调单据审批操作
*/
type AlibabaAlihealthDrugKytDestbillCheckRequest struct {
    model.Params

    // 企业ID
    refEntId   string 

    // 单据号
    billCode   string 

    // 审核状态，'Y'审批通过 'N' 审批不通过
    checkType   string 

}

func NewAlibabaAlihealthDrugKytDestbillCheckRequest() *AlibabaAlihealthDrugKytDestbillCheckRequest{
    return &AlibabaAlihealthDrugKytDestbillCheckRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytDestbillCheckRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.destbill.check"
}

func (r AlibabaAlihealthDrugKytDestbillCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytDestbillCheckRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytDestbillCheckRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytDestbillCheckRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

func (r AlibabaAlihealthDrugKytDestbillCheckRequest) GetBillCode() string {
    return r.billCode
}

func (r *AlibabaAlihealthDrugKytDestbillCheckRequest) SetCheckType(checkType string) error {
    r.checkType = checkType
    r.Set("check_type", checkType)
    return nil
}

func (r AlibabaAlihealthDrugKytDestbillCheckRequest) GetCheckType() string {
    return r.checkType
}

