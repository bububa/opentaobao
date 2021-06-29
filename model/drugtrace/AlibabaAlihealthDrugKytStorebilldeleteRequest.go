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
    refEntId   string
    // 操作人编码
    icCode   string
    // 单据ID
    billId   string
    // 单据类型
    billType   string
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
func (r *AlibabaAlihealthDrugKytStorebilldeleteRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytStorebilldeleteRequest) GetRefEntId() string {
    return r.refEntId
}
// IcCode Setter
// 操作人编码
func (r *AlibabaAlihealthDrugKytStorebilldeleteRequest) SetIcCode(icCode string) error {
    r.icCode = icCode
    r.Set("ic_code", icCode)
    return nil
}

// IcCode Getter
func (r AlibabaAlihealthDrugKytStorebilldeleteRequest) GetIcCode() string {
    return r.icCode
}
// BillId Setter
// 单据ID
func (r *AlibabaAlihealthDrugKytStorebilldeleteRequest) SetBillId(billId string) error {
    r.billId = billId
    r.Set("bill_id", billId)
    return nil
}

// BillId Getter
func (r AlibabaAlihealthDrugKytStorebilldeleteRequest) GetBillId() string {
    return r.billId
}
// BillType Setter
// 单据类型
func (r *AlibabaAlihealthDrugKytStorebilldeleteRequest) SetBillType(billType string) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytStorebilldeleteRequest) GetBillType() string {
    return r.billType
}
