package cainiaohandover

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// CainiaogloballogisticscarrierquerylistAPIRequest 实际承运商查询 API请求
// cainiao.global.logistics.carrier.querylist
//
// 查询出所有的实际承运商
type CainiaogloballogisticscarrierquerylistAPIRequest struct {
	model.Params
	// 多语言(暂不支持，保留入参)
	_locale string
}

// NewCainiaogloballogisticscarrierquerylistRequest 初始化CainiaogloballogisticscarrierquerylistAPIRequest对象
func NewCainiaogloballogisticscarrierquerylistRequest() *CainiaogloballogisticscarrierquerylistAPIRequest {
	return &CainiaogloballogisticscarrierquerylistAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r CainiaogloballogisticscarrierquerylistAPIRequest) GetApiMethodName() string {
	return "cainiao.global.logistics.carrier.querylist"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r CainiaogloballogisticscarrierquerylistAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r CainiaogloballogisticscarrierquerylistAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetLocale is Locale Setter
// 多语言(暂不支持，保留入参)
func (r *CainiaogloballogisticscarrierquerylistAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// GetLocale Locale Getter
func (r CainiaogloballogisticscarrierquerylistAPIRequest) GetLocale() string {
	return r._locale
}
