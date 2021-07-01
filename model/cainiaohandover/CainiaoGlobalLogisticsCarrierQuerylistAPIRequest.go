package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* CainiaoGlobalLogisticsCarrierQuerylistAPIRequest
实际承运商查询 API请求
cainiao.global.logistics.carrier.querylist

查询出所有的实际承运商 */
type CainiaoGlobalLogisticsCarrierQuerylistAPIRequest struct {
	model.Params
	// 多语言(暂不支持，保留入参)
	_locale string
}

// NewCainiaoGlobalLogisticsCarrierQuerylistRequest 初始化CainiaoGlobalLogisticsCarrierQuerylistAPIRequest对象
func NewCainiaoGlobalLogisticsCarrierQuerylistRequest() *CainiaoGlobalLogisticsCarrierQuerylistAPIRequest {
	return &CainiaoGlobalLogisticsCarrierQuerylistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaoGlobalLogisticsCarrierQuerylistAPIRequest) GetApiMethodName() string {
	return "cainiao.global.logistics.carrier.querylist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaoGlobalLogisticsCarrierQuerylistAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Locale Setter
// 多语言(暂不支持，保留入参)
func (r *CainiaoGlobalLogisticsCarrierQuerylistAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// Get Locale Getter
func (r CainiaoGlobalLogisticsCarrierQuerylistAPIRequest) GetLocale() string {
	return r._locale
}
