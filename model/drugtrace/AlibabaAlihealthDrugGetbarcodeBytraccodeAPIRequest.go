package drugtrace

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest 根据追溯码获取69码 API请求
// alibaba.alihealth.drug.getbarcode.bytraccode
//
// 根据追溯码获取69码
type AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest struct {
	model.Params
	// 追溯码
	_traceCode string
}

// NewAlibabaAlihealthDrugGetbarcodeBytraccodeRequest 初始化AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest对象
func NewAlibabaAlihealthDrugGetbarcodeBytraccodeRequest() *AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest {
	return &AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest) Reset() {
	r._traceCode = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest) GetApiMethodName() string {
	return "alibaba.alihealth.drug.getbarcode.bytraccode"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTraceCode is TraceCode Setter
// 追溯码
func (r *AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest) SetTraceCode(_traceCode string) error {
	r._traceCode = _traceCode
	r.Set("trace_code", _traceCode)
	return nil
}

// GetTraceCode TraceCode Getter
func (r AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest) GetTraceCode() string {
	return r._traceCode
}

var poolAlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaAlihealthDrugGetbarcodeBytraccodeRequest()
	},
}

// GetAlibabaAlihealthDrugGetbarcodeBytraccodeRequest 从 sync.Pool 获取 AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest
func GetAlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest() *AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest {
	return poolAlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest.Get().(*AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest)
}

// ReleaseAlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest 将 AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest 放入 sync.Pool
func ReleaseAlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest(v *AlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest) {
	v.Reset()
	poolAlibabaAlihealthDrugGetbarcodeBytraccodeAPIRequest.Put(v)
}
