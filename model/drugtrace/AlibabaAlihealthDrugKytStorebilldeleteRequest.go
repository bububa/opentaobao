package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售端单据删除 API请求
alibaba.alihealth.drug.kyt.storebilldelete

零售端单据删除
*/
type AlibabaAlihealthDrugKytStorebilldeleteRequest struct {
    model.Params
    // 企业ID
    _refEntId   string
    // 操作人编码
    _icCode   string
    // 单据ID
    _billId   string
    // 单据类型
    _billType   string
}

// 初始化AlibabaAlihealthDrugKytStorebilldeleteRequest对象
func NewAlibabaAlihealthDrugKytStorebilldeleteRequest() *AlibabaAlihealthDrugKytStorebilldeleteRequest{
    return &AlibabaAlihealthDrugKytStorebilldeleteRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytStorebilldeleteRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.storebilldelete"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytStorebilldeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytStorebilldeleteRequest) SetRefEntId(_refEntId string) error {
    r._refEntId = _refEntId
    r.Set("ref_ent_id", _refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytStorebilldeleteRequest) GetRefEntId() string {
    return r._refEntId
}
// IcCode Setter
// 操作人编码
func (r *AlibabaAlihealthDrugKytStorebilldeleteRequest) SetIcCode(_icCode string) error {
    r._icCode = _icCode
    r.Set("ic_code", _icCode)
    return nil
}

// IcCode Getter
func (r AlibabaAlihealthDrugKytStorebilldeleteRequest) GetIcCode() string {
    return r._icCode
}
// BillId Setter
// 单据ID
func (r *AlibabaAlihealthDrugKytStorebilldeleteRequest) SetBillId(_billId string) error {
    r._billId = _billId
    r.Set("bill_id", _billId)
    return nil
}

// BillId Getter
func (r AlibabaAlihealthDrugKytStorebilldeleteRequest) GetBillId() string {
    return r._billId
}
// BillType Setter
// 单据类型
func (r *AlibabaAlihealthDrugKytStorebilldeleteRequest) SetBillType(_billType string) error {
    r._billType = _billType
    r.Set("bill_type", _billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytStorebilldeleteRequest) GetBillType() string {
    return r._billType
}
