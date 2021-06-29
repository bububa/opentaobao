package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询上游企业出库单据号 APIRequest
alibaba.alihealth.drug.kyt.dr.queryupbillcode

疫苗温度合规补充需求-增加一个查询上游出库单号的接口。疾控在扫码入库时，接口通过扫到的码判定这个码对应所属的出库单据号
*/
type AlibabaAlihealthDrugKytDrQueryupbillcodeRequest struct {
    model.Params

    // 追溯码
    code   string 

    // 企业ID （一般为要查询单据的收货企业）
    refEntId   string 

}

func NewAlibabaAlihealthDrugKytDrQueryupbillcodeRequest() *AlibabaAlihealthDrugKytDrQueryupbillcodeRequest{
    return &AlibabaAlihealthDrugKytDrQueryupbillcodeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugKytDrQueryupbillcodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.kyt.dr.queryupbillcode"
}

func (r AlibabaAlihealthDrugKytDrQueryupbillcodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugKytDrQueryupbillcodeRequest) SetCode(code string) error {
    r.code = code
    r.Set("code", code)
    return nil
}

func (r AlibabaAlihealthDrugKytDrQueryupbillcodeRequest) GetCode() string {
    return r.code
}

func (r *AlibabaAlihealthDrugKytDrQueryupbillcodeRequest) SetRefEntId(refEntId string) error {
    r.refEntId = refEntId
    r.Set("ref_ent_id", refEntId)
    return nil
}

func (r AlibabaAlihealthDrugKytDrQueryupbillcodeRequest) GetRefEntId() string {
    return r.refEntId
}

