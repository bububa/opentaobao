package aeusergrowth

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressusergrowthsearchitemsgetAPIRequest 第三方平台搜索AE商品 API请求
// aliexpress.usergrowth.search.items.get
//
// 第三方平台的搜索服务   获取AE商品list
type AliexpressusergrowthsearchitemsgetAPIRequest struct {
	model.Params
	// user input keypods
	_keywords string
	// Third party tracking_id
	_trackingId string
	// currency code
	_currencyCode string
	// language
	_language string
	// page size
	_pageSize string
	// page index
	_pageIndex string
	// ship to country
	_countryCode string
}

// NewAliexpressusergrowthsearchitemsgetRequest 初始化AliexpressusergrowthsearchitemsgetAPIRequest对象
func NewAliexpressusergrowthsearchitemsgetRequest() *AliexpressusergrowthsearchitemsgetAPIRequest {
	return &AliexpressusergrowthsearchitemsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressusergrowthsearchitemsgetAPIRequest) GetApiMethodName() string {
	return "aliexpress.usergrowth.search.items.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressusergrowthsearchitemsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressusergrowthsearchitemsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetKeywords is Keywords Setter
// user input keypods
func (r *AliexpressusergrowthsearchitemsgetAPIRequest) SetKeywords(_keywords string) error {
	r._keywords = _keywords
	r.Set("keywords", _keywords)
	return nil
}

// GetKeywords Keywords Getter
func (r AliexpressusergrowthsearchitemsgetAPIRequest) GetKeywords() string {
	return r._keywords
}

// SetTrackingId is TrackingId Setter
// Third party tracking_id
func (r *AliexpressusergrowthsearchitemsgetAPIRequest) SetTrackingId(_trackingId string) error {
	r._trackingId = _trackingId
	r.Set("tracking_id", _trackingId)
	return nil
}

// GetTrackingId TrackingId Getter
func (r AliexpressusergrowthsearchitemsgetAPIRequest) GetTrackingId() string {
	return r._trackingId
}

// SetCurrencyCode is CurrencyCode Setter
// currency code
func (r *AliexpressusergrowthsearchitemsgetAPIRequest) SetCurrencyCode(_currencyCode string) error {
	r._currencyCode = _currencyCode
	r.Set("currency_code", _currencyCode)
	return nil
}

// GetCurrencyCode CurrencyCode Getter
func (r AliexpressusergrowthsearchitemsgetAPIRequest) GetCurrencyCode() string {
	return r._currencyCode
}

// SetLanguage is Language Setter
// language
func (r *AliexpressusergrowthsearchitemsgetAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// GetLanguage Language Getter
func (r AliexpressusergrowthsearchitemsgetAPIRequest) GetLanguage() string {
	return r._language
}

// SetPageSize is PageSize Setter
// page size
func (r *AliexpressusergrowthsearchitemsgetAPIRequest) SetPageSize(_pageSize string) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AliexpressusergrowthsearchitemsgetAPIRequest) GetPageSize() string {
	return r._pageSize
}

// SetPageIndex is PageIndex Setter
// page index
func (r *AliexpressusergrowthsearchitemsgetAPIRequest) SetPageIndex(_pageIndex string) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r AliexpressusergrowthsearchitemsgetAPIRequest) GetPageIndex() string {
	return r._pageIndex
}

// SetCountryCode is CountryCode Setter
// ship to country
func (r *AliexpressusergrowthsearchitemsgetAPIRequest) SetCountryCode(_countryCode string) error {
	r._countryCode = _countryCode
	r.Set("country_code", _countryCode)
	return nil
}

// GetCountryCode CountryCode Getter
func (r AliexpressusergrowthsearchitemsgetAPIRequest) GetCountryCode() string {
	return r._countryCode
}
