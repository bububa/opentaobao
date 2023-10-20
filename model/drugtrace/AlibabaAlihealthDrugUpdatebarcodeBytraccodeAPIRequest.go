package drugtrace

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaalihealthdrugupdatebarcodebytraccodeAPIRequest 根据追溯码修改69码 API请求
// alibaba.alihealth.drug.updatebarcode.bytraccode
//
// 根据追溯码修改69码
type AlibabaalihealthdrugupdatebarcodebytraccodeAPIRequest struct {
	model.Params
	// 追溯码
	_traceCode string
	// 69码
	_barcode string
}

// NewAlibabaalihealthdrugupdatebarcodebytraccodeRequest 初始化AlibabaalihealthdrugupdatebarcodebytraccodeAPIRequest对象
func NewAlibabaalihealthdrugupdatebarcodebytraccodeRequest() *AlibabaalihealthdrugupdatebarcodebytraccodeAPIRequest {
	return &AlibabaalihealthdrugupdatebarcodebytraccodeAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaalihealthdrugupdatebarcodebytraccodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.updatebarcode.bytraccode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaalihealthdrugupdatebarcodebytraccodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaalihealthdrugupdatebarcodebytraccodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTraceCode is TraceCode Setter
// 追溯码
func (r *AlibabaalihealthdrugupdatebarcodebytraccodeAPIRequest) SetTraceCode(_traceCode string) error {
	r._traceCode = _traceCode
	r.Set("trace_code", _traceCode)
	return nil
}

// GetTraceCode TraceCode Getter
func (r AlibabaalihealthdrugupdatebarcodebytraccodeAPIRequest) GetTraceCode() string {
	return r._traceCode
}

// SetBarcode is Barcode Setter
// 69码
func (r *AlibabaalihealthdrugupdatebarcodebytraccodeAPIRequest) SetBarcode(_barcode string) error {
	r._barcode = _barcode
	r.Set("barcode", _barcode)
	return nil
}

// GetBarcode Barcode Getter
func (r AlibabaalihealthdrugupdatebarcodebytraccodeAPIRequest) GetBarcode() string {
	return r._barcode
}
