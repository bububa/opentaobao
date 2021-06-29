package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
零售端单据删除 APIRequest
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

func NewAlibabaAlihealthDrugKytStorebilldeleteRequest() *AlibabaAlihealthDrugKytStorebilldeleteRequest{
    return &AlibabaAlihealthDrugKytStorebilldeleteRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytStorebilldeleteRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.storebilldelete"
}

func (r AlibabaAlihealthDrugKytStorebilldeleteRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytStorebilldeleteRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytStorebilldeleteRequest) GetRefEntId() string {
    return r.refEntId
}

func (r *AlibabaAlihealthDrugKytStorebilldeleteRequest) SetIcCode(icCode string) error {
    r.icCode = icCode
    r.Set("ic_code", icCode)
    return nil
}

func (r AlibabaAlihealthDrugKytStorebilldeleteRequest) GetIcCode() string {
    return r.icCode
}

func (r *AlibabaAlihealthDrugKytStorebilldeleteRequest) SetBillId(billId string) error {
    r.billId = billId
    r.Set("bill_id", billId)
    return nil
}

func (r AlibabaAlihealthDrugKytStorebilldeleteRequest) GetBillId() string {
    return r.billId
}

func (r *AlibabaAlihealthDrugKytStorebilldeleteRequest) SetBillType(billType string) error {
    r.billType = billType
    r.Set("bill_type", billType)
    return nil
}

func (r AlibabaAlihealthDrugKytStorebilldeleteRequest) GetBillType() string {
    return r.billType
}

