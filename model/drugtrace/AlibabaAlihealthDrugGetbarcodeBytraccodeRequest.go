package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据追溯码获取69码 APIRequest
alibaba.alihealth.drug.getbarcode.bytraccode

根据追溯码获取69码
*/
type AlibabaAlihealthDrugGetbarcodeBytraccodeRequest struct {
    model.Params

    // 追溯码
    traceCode   string 

}

func NewAlibabaAlihealthDrugGetbarcodeBytraccodeRequest() *AlibabaAlihealthDrugGetbarcodeBytraccodeRequest{
    return &AlibabaAlihealthDrugGetbarcodeBytraccodeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugGetbarcodeBytraccodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.getbarcode.bytraccode"
}

func (r AlibabaAlihealthDrugGetbarcodeBytraccodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugGetbarcodeBytraccodeRequest) SetTraceCode(traceCode string) error {
    r.traceCode = traceCode
    r.Set("trace_code", traceCode)
    return nil
}

func (r AlibabaAlihealthDrugGetbarcodeBytraccodeRequest) GetTraceCode() string {
    return r.traceCode
}

