package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售修改出入库单追溯码 APIRequest
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

func NewAlibabaAlihealthDrugKytUpdatebillcodeRequest() *AlibabaAlihealthDrugKytUpdatebillcodeRequest{
    return &AlibabaAlihealthDrugKytUpdatebillcodeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytUpdatebillcodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.updatebillcode"
}

func (r AlibabaAlihealthDrugKytUpdatebillcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytUpdatebillcodeRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytUpdatebillcodeRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytUpdatebillcodeRequest) SetIcCode(icCode string) error {
    r.icCode = icCode
    r.Set("ic_code", icCode)
    return nil
}

func (r AlibabaAlihealthDrugKytUpdatebillcodeRequest) GetIcCode() string {
    return r.icCode
}

func (r *AlibabaAlihealthDrugKytUpdatebillcodeRequest) SetBillId(billId string) error {
    r.billId = billId
    r.Set("bill_id", billId)
    return nil
}

func (r AlibabaAlihealthDrugKytUpdatebillcodeRequest) GetBillId() string {
    return r.billId
}

func (r *AlibabaAlihealthDrugKytUpdatebillcodeRequest) SetBillType(billType string) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

func (r AlibabaAlihealthDrugKytUpdatebillcodeRequest) GetBillType() string {
    return r.billType
}

func (r *AlibabaAlihealthDrugKytUpdatebillcodeRequest) SetCodeList(codeList []string) error {
    r.codeList = codeList
    r.Set("code_list", codeList)
    return nil
}

func (r AlibabaAlihealthDrugKytUpdatebillcodeRequest) GetCodeList() []string {
    return r.codeList
}

