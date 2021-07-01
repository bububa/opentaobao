package aeusergrowth

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressUsergrowthRecommendItemsGetAPIRequest
第三方平台推荐商品 API请求
aliexpress.usergrowth.recommend.items.get

第三方平台的推荐AE商品   场景：skin 、底部推荐等 */
type AliexpressUsergrowthRecommendItemsGetAPIRequest struct {
	model.Params
	// third party trackingId
	_trackingId string
	// currency Code
	_currencyCode string
	// language
	_language string
	// user type
	_userTypeCode string
	// page index,start from 1
	_pageIndex string
	// page size
	_pageSize string
	// country code
	_countryCode string
}

// NewAliexpressUsergrowthRecommendItemsGetRequest 初始化AliexpressUsergrowthRecommendItemsGetAPIRequest对象
func NewAliexpressUsergrowthRecommendItemsGetRequest() *AliexpressUsergrowthRecommendItemsGetAPIRequest {
	return &AliexpressUsergrowthRecommendItemsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressUsergrowthRecommendItemsGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.usergrowth.recommend.items.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressUsergrowthRecommendItemsGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is TrackingId Setter
// third party trackingId
func (r *AliexpressUsergrowthRecommendItemsGetAPIRequest) SetTrackingId(_trackingId string) error {
	r._trackingId = _trackingId
	r.Set("tracking_id", _trackingId)
	return nil
}

// Get TrackingId Getter
func (r AliexpressUsergrowthRecommendItemsGetAPIRequest) GetTrackingId() string {
	return r._trackingId
}

// Set is CurrencyCode Setter
// currency Code
func (r *AliexpressUsergrowthRecommendItemsGetAPIRequest) SetCurrencyCode(_currencyCode string) error {
	r._currencyCode = _currencyCode
	r.Set("currency_code", _currencyCode)
	return nil
}

// Get CurrencyCode Getter
func (r AliexpressUsergrowthRecommendItemsGetAPIRequest) GetCurrencyCode() string {
	return r._currencyCode
}

// Set is Language Setter
// language
func (r *AliexpressUsergrowthRecommendItemsGetAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// Get Language Getter
func (r AliexpressUsergrowthRecommendItemsGetAPIRequest) GetLanguage() string {
	return r._language
}

// Set is UserTypeCode Setter
// user type
func (r *AliexpressUsergrowthRecommendItemsGetAPIRequest) SetUserTypeCode(_userTypeCode string) error {
	r._userTypeCode = _userTypeCode
	r.Set("user_type_code", _userTypeCode)
	return nil
}

// Get UserTypeCode Getter
func (r AliexpressUsergrowthRecommendItemsGetAPIRequest) GetUserTypeCode() string {
	return r._userTypeCode
}

// Set is PageIndex Setter
// page index,start from 1
func (r *AliexpressUsergrowthRecommendItemsGetAPIRequest) SetPageIndex(_pageIndex string) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// Get PageIndex Getter
func (r AliexpressUsergrowthRecommendItemsGetAPIRequest) GetPageIndex() string {
	return r._pageIndex
}

// Set is PageSize Setter
// page size
func (r *AliexpressUsergrowthRecommendItemsGetAPIRequest) SetPageSize(_pageSize string) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AliexpressUsergrowthRecommendItemsGetAPIRequest) GetPageSize() string {
	return r._pageSize
}

// Set is CountryCode Setter
// country code
func (r *AliexpressUsergrowthRecommendItemsGetAPIRequest) SetCountryCode(_countryCode string) error {
	r._countryCode = _countryCode
	r.Set("country_code", _countryCode)
	return nil
}

// Get CountryCode Getter
func (r AliexpressUsergrowthRecommendItemsGetAPIRequest) GetCountryCode() string {
	return r._countryCode
}
