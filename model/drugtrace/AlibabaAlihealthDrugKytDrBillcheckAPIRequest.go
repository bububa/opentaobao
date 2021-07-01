package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗追溯验证 API请求
alibaba.alihealth.drug.kyt.dr.billcheck

各级疾控在入库完成后，需要做追溯信息验证
*/
type AlibabaAlihealthDrugKytDrBillcheckAPIRequest struct {
    model.Params
    // 调用企业ID
    _refEntId   string
    // 单据编号
    _billCode   string
    // 单据类型
    _billType   string
    // 单据企业refEntId
    _owerRefEntId   string
}

// 初始化AlibabaAlihealthDrugKytDrBillcheckAPIRequest对象
func NewAlibabaAlihealthDrugKytDrBillcheckRequest() *AlibabaAlihealthDrugKytDrBillcheckAPIRequest{
    return &AlibabaAlihealthDrugKytDrBillcheckAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytDrBillcheckAPIRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.billcheck"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytDrBillcheckAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 调用企业ID
func (r *AlibabaAlihealthDrugKytDrBillcheckAPIRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytDrBillcheckAPIRequest) GetRefEntId() string {
    return r._refEntId
}
// BillCode Setter
// 单据编号
func (r *AlibabaAlihealthDrugKytDrBillcheckAPIRequest) SetBillCode(_billCode string) error {
    r._billCode = _billCode
    r.Set("bill_code", _billCode)
    return nil
}

// BillCode Getter
func (r AlibabaAlihealthDrugKytDrBillcheckAPIRequest) GetBillCode() string {
    return r._billCode
}
// BillType Setter
// 单据类型
func (r *AlibabaAlihealthDrugKytDrBillcheckAPIRequest) SetBillType(_billType string) error {
    r._billType = _billType
    r.Set("bill_type", _billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytDrBillcheckAPIRequest) GetBillType() string {
    return r._billType
}
// OwerRefEntId Setter
// 单据企业refEntId
func (r *AlibabaAlihealthDrugKytDrBillcheckAPIRequest) SetOwerRefEntId(_owerRefEntId string) error {
    r._owerRefEntId = _owerRefEntId
    r.Set("ower_ref_ent_id", _owerRefEntId)
    return nil
}

// OwerRefEntId Getter
func (r AlibabaAlihealthDrugKytDrBillcheckAPIRequest) GetOwerRefEntId() string {
    return r._owerRefEntId
}
