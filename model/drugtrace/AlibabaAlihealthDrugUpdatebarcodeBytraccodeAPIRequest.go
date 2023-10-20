package drugtrace

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest) Reset() {
	r._traceCode = ""
	r._barcode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.updatebarcode.bytraccode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest) GetRawParams() model.Params {
	return r.Params
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

var poolAlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugUpdatebarcodeBytraccodeRequest()
	},
}

// GetAlibabaAlihealthDrugUpdatebarcodeBytraccodeRequest 从 sync.Pool 获取 AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest
func GetAlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest() *AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest {
	return poolAlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest.Get().(*AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest)
}

// ReleaseAlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest 将 AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest(v *AlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugUpdatebarcodeBytraccodeAPIRequest.Put(v)
}
