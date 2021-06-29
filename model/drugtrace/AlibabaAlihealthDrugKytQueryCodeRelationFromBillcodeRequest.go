package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据单据号码查询码单据详情和码信息 APIRequest
alibaba.alihealth.drug.kyt.query.code.relation.from.billcode

根据单据号码查询码单据详情和码信息
*/
type AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeRequest struct {
    model.Params

    // 单据号码
    billCode   string 

    // 企业refEntId
    refEntId   string 

}

func NewAlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeRequest() *AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeRequest{
    return &AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.query.code.relation.from.billcode"
}

func (r AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

func (r AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeRequest) GetBillCode() string {
    return r.billCode
}

func (r *AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytQueryCodeRelationFromBillcodeRequest) GetRefEntId() string {
    return r.refEntId
}

