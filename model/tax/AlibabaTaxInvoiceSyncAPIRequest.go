package tax

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaTaxInvoiceSyncAPIRequest 第三方开票回调API API请求
// alibaba.tax.invoice.sync
//
// 该接口只提供给俄罗斯供应商开具发票使用，请勿申请。
type AlibabaTaxInvoiceSyncAPIRequest struct {
	model.Params
	// 回调发票信息，请求参数详情见链接：https://yuque.antfin-inc.com/tax/biw99l/uestb7
	_paramJson string
}

// NewAlibabaTaxInvoiceSyncRequest 初始化AlibabaTaxInvoiceSyncAPIRequest对象
func NewAlibabaTaxInvoiceSyncRequest() *AlibabaTaxInvoiceSyncAPIRequest {
	return &AlibabaTaxInvoiceSyncAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaTaxInvoiceSyncAPIRequest) Reset() {
	r._paramJson = ""
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTaxInvoiceSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.tax.invoice.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTaxInvoiceSyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaTaxInvoiceSyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamJson is ParamJson Setter
// 回调发票信息，请求参数详情见链接：https://yuque.antfin-inc.com/tax/biw99l/uestb7
func (r *AlibabaTaxInvoiceSyncAPIRequest) SetParamJson(_paramJson string) error {
	r._paramJson = _paramJson
	r.Set("param_json", _paramJson)
	return nil
}

// GetParamJson ParamJson Getter
func (r AlibabaTaxInvoiceSyncAPIRequest) GetParamJson() string {
	return r._paramJson
}

var poolAlibabaTaxInvoiceSyncAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaTaxInvoiceSyncRequest()
	},
}

// GetAlibabaTaxInvoiceSyncRequest 从 sync.Pool 获取 AlibabaTaxInvoiceSyncAPIRequest
func GetAlibabaTaxInvoiceSyncAPIRequest() *AlibabaTaxInvoiceSyncAPIRequest {
	return poolAlibabaTaxInvoiceSyncAPIRequest.Get().(*AlibabaTaxInvoiceSyncAPIRequest)
}

// ReleaseAlibabaTaxInvoiceSyncAPIRequest 将 AlibabaTaxInvoiceSyncAPIRequest 放入 sync.Pool
func ReleaseAlibabaTaxInvoiceSyncAPIRequest(v *AlibabaTaxInvoiceSyncAPIRequest) {
	v.Reset()
	poolAlibabaTaxInvoiceSyncAPIRequest.Put(v)
}
