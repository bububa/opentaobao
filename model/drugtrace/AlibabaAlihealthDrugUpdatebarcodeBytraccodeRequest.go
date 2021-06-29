package drugtrace

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
根据追溯码修改69码 API请求
alibaba.alihealth.drug.updatebarcode.bytraccode

根据追溯码修改69码
*/
type AlibabaAlihealthDrugUpdatebarcodeBytraccodeRequest struct {
    model.Params
    // 追溯码
    _traceCode   string
    // 69码
    _barcode   string
}

// 初始化AlibabaAlihealthDrugUpdatebarcodeBytraccodeRequest对象
func NewAlibabaAlihealthDrugUpdatebarcodeBytraccodeRequest() *AlibabaAlihealthDrugUpdatebarcodeBytraccodeRequest{
    return &AlibabaAlihealthDrugUpdatebarcodeBytraccodeRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugUpdatebarcodeBytraccodeRequest) GetApiMethodName() string {
    return "alibaba.alihealth.drug.updatebarcode.bytraccode"
}

// IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugUpdatebarcodeBytraccodeRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TraceCode Setter
// 追溯码
func (r *AlibabaAlihealthDrugUpdatebarcodeBytraccodeRequest) SetTraceCode(_traceCode string) error {
    r._traceCode = _traceCode
    r.Set("trace_code", _traceCode)
    return nil
}

// TraceCode Getter
func (r AlibabaAlihealthDrugUpdatebarcodeBytraccodeRequest) GetTraceCode() string {
    return r._traceCode
}
// Barcode Setter
// 69码
func (r *AlibabaAlihealthDrugUpdatebarcodeBytraccodeRequest) SetBarcode(_barcode string) error {
    r._barcode = _barcode
    r.Set("barcode", _barcode)
    return nil
}

// Barcode Getter
func (r AlibabaAlihealthDrugUpdatebarcodeBytraccodeRequest) GetBarcode() string {
    return r._barcode
}
