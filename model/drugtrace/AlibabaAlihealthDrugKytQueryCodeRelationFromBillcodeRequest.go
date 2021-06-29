package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据单据号码查询码单据详情和码信息 API请求
alibaba.alihealth.drug.kyt.query.code.relation.from.billcode

根据单据号码查询码单据详情和码信息
*/
type AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeRequest struct {
    model.Params
    // 单据号码
    _billCode   string
    // 企业refEntId
    _refEntId   string
}

// 初始化AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeRequest对象
func NewAlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeRequest() *AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeRequest{
    return &AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.query.code.relation.from.billcode"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// BillCode Setter
// 单据号码
func (r *AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeRequest) SetBillCode(_billCode string) error {
    r._billCode = _billCode
    r.Set("bill_code", _billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeRequest) GetBillCode() string {
    return r._billCode
}
// RefEntId Setter
// 企业refEntId
func (r *AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeRequest) GetRefEntId() string {
    return r._refEntId
}
