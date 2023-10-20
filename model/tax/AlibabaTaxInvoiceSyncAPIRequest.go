package tax

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabataxinvoicesyncAPIRequest 第三方开票回调API API请求
// alibaba.tax.invoice.sync
//
// 该接口只提供给俄罗斯供应商开具发票使用，请勿申请。
type AlibabataxinvoicesyncAPIRequest struct {
	model.Params
	// 回调发票信息，请求参数详情见链接：https://yuque.antfin-inc.com/tax/biw99l/uestb7
	_paramJson string
}

// NewAlibabataxinvoicesyncRequest 初始化AlibabataxinvoicesyncAPIRequest对象
func NewAlibabataxinvoicesyncRequest() *AlibabataxinvoicesyncAPIRequest {
	return &AlibabataxinvoicesyncAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabataxinvoicesyncAPIRequest) GetApiMethodName() string {
	return "alibaba.tax.invoice.sync"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabataxinvoicesyncAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabataxinvoicesyncAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetParamJson is ParamJson Setter
// 回调发票信息，请求参数详情见链接：https://yuque.antfin-inc.com/tax/biw99l/uestb7
func (r *AlibabataxinvoicesyncAPIRequest) SetParamJson(_paramJson string) error {
	r._paramJson = _paramJson
	r.Set("param_json", _paramJson)
	return nil
}

// GetParamJson ParamJson Getter
func (r AlibabataxinvoicesyncAPIRequest) GetParamJson() string {
	return r._paramJson
}
