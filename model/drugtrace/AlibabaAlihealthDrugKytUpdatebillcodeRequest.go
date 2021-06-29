package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售修改出入库单追溯码 API请求
alibaba.alihealth.drug.kyt.updatebillcode

零售修改出入库单追溯码
*/
type AlibabaAlihealthDrugKytUpdatebillcodeRequest struct {
    model.Params
    // 企业ID
    refEntId   string
    // 操作人ID
    icCode   string
    // 单据ID
    billId   string
    // 单据类型
    billType   string
    // 追溯码
    codeList   []string
}

// 初始化AlibabaAlihealthDrugKytUpdatebillcodeRequest对象
func NewAlibabaAlihealthDrugKytUpdatebillcodeRequest() *AlibabaAlihealthDrugKytUpdatebillcodeRequest{
    return &AlibabaAlihealthDrugKytUpdatebillcodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugKytUpdatebillcodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.updatebillcode"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugKytUpdatebillcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// RefEntId Setter
// 企业ID
func (r *AlibabaAlihealthDrugKytUpdatebillcodeRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

// RefEntId Getter
func (r AlibabaAlihealthDrugKytUpdatebillcodeRequest) GetRefEntId() string {
    return r.refEntId
}
// IcCode Setter
// 操作人ID
func (r *AlibabaAlihealthDrugKytUpdatebillcodeRequest) SetIcCode(icCode string) error {
    r.icCode = icCode
    r.Set("ic_code", icCode)
    return nil
}

// IcCode Getter
func (r AlibabaAlihealthDrugKytUpdatebillcodeRequest) GetIcCode() string {
    return r.icCode
}
// BillId Setter
// 单据ID
func (r *AlibabaAlihealthDrugKytUpdatebillcodeRequest) SetBillId(billId string) error {
    r.billId = billId
    r.Set("bill_id", billId)
    return nil
}

// BillId Getter
func (r AlibabaAlihealthDrugKytUpdatebillcodeRequest) GetBillId() string {
    return r.billId
}
// BillType Setter
// 单据类型
func (r *AlibabaAlihealthDrugKytUpdatebillcodeRequest) SetBillType(billType string) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

// BillType Getter
func (r AlibabaAlihealthDrugKytUpdatebillcodeRequest) GetBillType() string {
    return r.billType
}
// CodeList Setter
// 追溯码
func (r *AlibabaAlihealthDrugKytUpdatebillcodeRequest) SetCodeList(codeList []string) error {
    r.codeList = codeList
    r.Set("code_list", codeList)
    return nil
}

// CodeList Getter
func (r AlibabaAlihealthDrugKytUpdatebillcodeRequest) GetCodeList() []string {
    return r.codeList
}
