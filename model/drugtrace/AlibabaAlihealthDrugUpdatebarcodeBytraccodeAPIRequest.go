package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest 根据追溯码修改69码 API请求
// alibaba.alihealth.drug.updatebarcode.bytraccode
//
// 根据追溯码修改69码
type AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest struct {
	model.Params
	// 追溯码
	_traceCode string
	// 69码
	_barcode string
}

// NewAlibabaAlihealthDrugUpdatebarcodeBytraccodeRequest 初始化AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest对象
func NewAlibabaAlihealthDrugUpdatebarcodeBytraccodeRequest() *AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest {
	return &AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.updatebarcode.bytraccode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTraceCode is TraceCode Setter
// 追溯码
func (r *AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest) SetTraceCode(_traceCode string) error {
	r._traceCode = _traceCode
	r.Set("trace_code", _traceCode)
	return nil
}

// GetTraceCode TraceCode Getter
func (r AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest) GetTraceCode() string {
	return r._traceCode
}

// SetBarcode is Barcode Setter
// 69码
func (r *AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest) SetBarcode(_barcode string) error {
	r._barcode = _barcode
	r.Set("barcode", _barcode)
	return nil
}

// GetBarcode Barcode Getter
func (r AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest) GetBarcode() string {
	return r._barcode
}
