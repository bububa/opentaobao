package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据单据编号查询单据明细 APIRequest
alibaba.alihealth.drug.kyt.query.druginfo.from.billcode

根据单据编号查询单据明细
*/
type AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeRequest struct {
    model.Params

    // 单据号
    billCode   string 

    // 企业id
    refEntId   string 

}

func NewAlibabaAlihealthDrugKytQueryDruginfoFromBillcodeRequest() *AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeRequest{
    return &AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.query.druginfo.from.billcode"
}

func (r AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeRequest) SetBillCode(billCode string) error {
    r.billCode = billCode
    r.Set("bill_code", billCode)
    return nil
}

func (r AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeRequest) GetBillCode() string {
    return r.billCode
}

func (r *AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytQueryDruginfoFromBillcodeRequest) GetRefEntId() string {
    return r.refEntId
}

