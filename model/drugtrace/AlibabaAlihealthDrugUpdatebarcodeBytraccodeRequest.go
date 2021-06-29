package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据追溯码修改69码 APIRequest
alibaba.alihealth.drug.updatebarcode.bytraccode

根据追溯码修改69码
*/
type AlibabaAlihealthDrugUpdatebarcodeBytraccodeRequest struct {
    model.Params

    // 追溯码
    traceCode   string 

    // 69码
    barcode   string 

}

func NewAlibabaAlihealthDrugUpdatebarcodeBytraccodeRequest() *AlibabaAlihealthDrugUpdatebarcodeBytraccodeRequest{
    return &AlibabaAlihealthDrugUpdatebarcodeBytraccodeRequest{
        Params: model.NewParams(),
    }
}

func (r AlibabaAlihealthDrugUpdatebarcodeBytraccodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.updatebarcode.bytraccode"
}

func (r AlibabaAlihealthDrugUpdatebarcodeBytraccodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AlibabaAlihealthDrugUpdatebarcodeBytraccodeRequest) SetTraceCode(traceCode string) error {
    r.traceCode = traceCode
    r.Set("trace_code", traceCode)
    return nil
}

func (r AlibabaAlihealthDrugUpdatebarcodeBytraccodeRequest) GetTraceCode() string {
    return r.traceCode
}

func (r *AlibabaAlihealthDrugUpdatebarcodeBytraccodeRequest) SetBarcode(barcode string) error {
    r.barcode = barcode
    r.Set("barcode", barcode)
    return nil
}

func (r AlibabaAlihealthDrugUpdatebarcodeBytraccodeRequest) GetBarcode() string {
    return r.barcode
}

