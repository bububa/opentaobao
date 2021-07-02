package tax

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaTaxInvoiceSyncAPIRequest) GetApiMethodName() string {
	return "alibaba.tax.invoice.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaTaxInvoiceSyncAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is ParamJson Setter
// 回调发票信息，请求参数详情见链接：https://yuque.antfin-inc.com/tax/biw99l/uestb7
func (r *AlibabaTaxInvoiceSyncAPIRequest) SetParamJson(_paramJson string) error {
	r._paramJson = _paramJson
	r.Set("param_json", _paramJson)
	return nil
}

// Get ParamJson Getter
func (r AlibabaTaxInvoiceSyncAPIRequest) GetParamJson() string {
	return r._paramJson
}
