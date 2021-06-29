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
    _refEntId   string
    // 单据号
    _billCode   string
    // 审核状态，'Y'审批通过 'N' 审批不通过
    _checkType   string
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
func (r *AlibabaAlihealthDrugKytDestbillCheckRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytDestbillCheckRequest) GetRefEntId() string {
    return r._refEntId
}
// BillCode Setter
// 单据号
func (r *AlibabaAlihealthDrugKytDestbillCheckRequest) SetBillCode(_billCode string) error {
    r._billCode = _billCode
    r.Set("bill_code", _billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytDestbillCheckRequest) GetBillCode() string {
    return r._billCode
}
// CheckType Setter
// 审核状态，'Y'审批通过 'N' 审批不通过
func (r *AlibabaAlihealthDrugKytDestbillCheckRequest) SetCheckType(_checkType string) error {
    r._checkType = _checkType
    r.Set("check_type", _checkType)
    return nil
}

// CheckType Getter
func (r AlibabaAlihealthDrugKytDestbillCheckRequest) GetCheckType() string {
    return r._checkType
}