package aeusergrowth

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressUsergrowthSearchItemsGetAPIRequest 第三方平台搜索AE商品 API请求
// aliexpress.usergrowth.search.items.get
//
// 第三方平台的搜索服务   获取AE商品list
type AliexpressUsergrowthSearchItemsGetAPIRequest struct {
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

// NewAliexpressUsergrowthSearchItemsGetRequest 初始化AliexpressUsergrowthSearchItemsGetAPIRequest对象
func NewAliexpressUsergrowthSearchItemsGetRequest() *AliexpressUsergrowthSearchItemsGetAPIRequest {
	return &AliexpressUsergrowthSearchItemsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressUsergrowthSearchItemsGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.usergrowth.search.items.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressUsergrowthSearchItemsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Keywords Setter
// user input keypods
func (r *AliexpressUsergrowthSearchItemsGetAPIRequest) SetKeywords(_keywords string) error {
	r._keywords = _keywords
	r.Set("keywords", _keywords)
	return nil
}

// Get Keywords Getter
func (r AliexpressUsergrowthSearchItemsGetAPIRequest) GetKeywords() string {
	return r._keywords
}

// Set is TrackingId Setter
// Third party tracking_id
func (r *AliexpressUsergrowthSearchItemsGetAPIRequest) SetTrackingId(_trackingId string) error {
	r._trackingId = _trackingId
	r.Set("tracking_id", _trackingId)
	return nil
}

// Get TrackingId Getter
func (r AliexpressUsergrowthSearchItemsGetAPIRequest) GetTrackingId() string {
	return r._trackingId
}

// Set is CurrencyCode Setter
// currency code
func (r *AliexpressUsergrowthSearchItemsGetAPIRequest) SetCurrencyCode(_currencyCode string) error {
	r._currencyCode = _currencyCode
	r.Set("currency_code", _currencyCode)
	return nil
}

// Get CurrencyCode Getter
func (r AliexpressUsergrowthSearchItemsGetAPIRequest) GetCurrencyCode() string {
	return r._currencyCode
}

// Set is Language Setter
// language
func (r *AliexpressUsergrowthSearchItemsGetAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// Get Language Getter
func (r AliexpressUsergrowthSearchItemsGetAPIRequest) GetLanguage() string {
	return r._language
}

// Set is PageSize Setter
// page size
func (r *AliexpressUsergrowthSearchItemsGetAPIRequest) SetPageSize(_pageSize string) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AliexpressUsergrowthSearchItemsGetAPIRequest) GetPageSize() string {
	return r._pageSize
}

// Set is PageIndex Setter
// page index
func (r *AliexpressUsergrowthSearchItemsGetAPIRequest) SetPageIndex(_pageIndex string) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// Get PageIndex Getter
func (r AliexpressUsergrowthSearchItemsGetAPIRequest) GetPageIndex() string {
	return r._pageIndex
}

// Set is CountryCode Setter
// ship to country
func (r *AliexpressUsergrowthSearchItemsGetAPIRequest) SetCountryCode(_countryCode string) error {
	r._countryCode = _countryCode
	r.Set("country_code", _countryCode)
	return nil
}

// Get CountryCode Getter
func (r AliexpressUsergrowthSearchItemsGetAPIRequest) GetCountryCode() string {
	return r._countryCode
}
