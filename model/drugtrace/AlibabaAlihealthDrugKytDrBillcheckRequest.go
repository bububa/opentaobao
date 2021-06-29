package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
疫苗追溯验证 APIRequest
alibaba.alihealth.drug.kyt.dr.billcheck

各级疾控在入库完成后，需要做追溯信息验证
*/
type AlibabaAlihealthDrugKytDrBillcheckRequest struct {
    model.Params

    // 调用企业ID
    refEntId   string 

    // 单据编号
    billCode   string 

    // 单据类型
    billType   string 

    // 单据企业refEntId
    owerRefEntId   string 

}

func NewAlibabaAlihealthDrugKytDrBillcheckRequest() *AlibabaAlihealthDrugKytDrBillcheckRequest{
    return &AlibabaAlihealthDrugKytDrBillcheckRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytDrBillcheckRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.billcheck"
}

func (r AlibabaAlihealthDrugKytDrBillcheckRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytDrBillcheckRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrBillcheckRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytDrBillcheckRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

func (r AlibabaAlihealthDrugKytDrBillcheckRequest) GetBillCode() string {
    return r.billCode
}

func (r *AlibabaAlihealthDrugKytDrBillcheckRequest) SetBillType(billType string) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

func (r AlibabaAlihealthDrugKytDrBillcheckRequest) GetBillType() string {
    return r.billType
}

func (r *AlibabaAlihealthDrugKytDrBillcheckRequest) SetOwerRefEntId(owerRefEntId string) error {
    r.owerRefEntId = owerRefEntId
    r.Set("ower_ref_ent_id", owerRefEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrBillcheckRequest) GetOwerRefEntId() string {
    return r.owerRefEntId
}

