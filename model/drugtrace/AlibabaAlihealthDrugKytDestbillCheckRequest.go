package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
直调审批 API请求
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

// 初始化AlibabaAlihealthDrugKytDestbillCheckRequest对象
func NewAlibabaAlihealthDrugKytDestbillCheckRequest() *AlibabaAlihealthDrugKytDestbillCheckRequest{
    return &AlibabaAlihealthDrugKytDestbillCheckRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDestbillCheckRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.destbill.check"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDestbillCheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytDestbillCheckRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytDestbillCheckRequest) GetRefEntId() string {
    return r.refEntId
}
// BillCode Setter
// 单据号
func (r *AlibabaAlihealthDrugKytDestbillCheckRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytDestbillCheckRequest) GetBillCode() string {
    return r.billCode
}
// CheckType Setter
// 审核状态，'Y'审批通过 'N' 审批不通过
func (r *AlibabaAlihealthDrugKytDestbillCheckRequest) SetCheckType(checkType string) error {
    r.checkType = checkType
    r.Set("check_type", checkType)
    return nil
}

// CheckType Getter
func (r AlibabaAlihealthDrugKytDestbillCheckRequest) GetCheckType() string {
    return r.checkType
}
