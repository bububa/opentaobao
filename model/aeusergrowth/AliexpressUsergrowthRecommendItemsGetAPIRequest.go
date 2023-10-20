package aeusergrowth

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressUsergrowthRecommendItemsGetAPIRequest 第三方平台推荐商品 API请求
// aliexpress.usergrowth.recommend.items.get
//
// 第三方平台的推荐AE商品   场景：skin 、底部推荐等
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
func (r AliexpressUsergrowthRecommendItemsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressUsergrowthRecommendItemsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTrackingId is TrackingId Setter
// third party trackingId
func (r *AliexpressUsergrowthRecommendItemsGetAPIRequest) SetTrackingId(_trackingId string) error {
	r._trackingId = _trackingId
	r.Set("tracking_id", _trackingId)
	return nil
}

// GetTrackingId TrackingId Getter
func (r AliexpressUsergrowthRecommendItemsGetAPIRequest) GetTrackingId() string {
	return r._trackingId
}

// SetCurrencyCode is CurrencyCode Setter
// currency Code
func (r *AliexpressUsergrowthRecommendItemsGetAPIRequest) SetCurrencyCode(_currencyCode string) error {
	r._currencyCode = _currencyCode
	r.Set("currency_code", _currencyCode)
	return nil
}

// GetCurrencyCode CurrencyCode Getter
func (r AliexpressUsergrowthRecommendItemsGetAPIRequest) GetCurrencyCode() string {
	return r._currencyCode
}

// SetLanguage is Language Setter
// language
func (r *AliexpressUsergrowthRecommendItemsGetAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// GetLanguage Language Getter
func (r AliexpressUsergrowthRecommendItemsGetAPIRequest) GetLanguage() string {
	return r._language
}

// SetUserTypeCode is UserTypeCode Setter
// user type
func (r *AliexpressUsergrowthRecommendItemsGetAPIRequest) SetUserTypeCode(_userTypeCode string) error {
	r._userTypeCode = _userTypeCode
	r.Set("user_type_code", _userTypeCode)
	return nil
}

// GetUserTypeCode UserTypeCode Getter
func (r AliexpressUsergrowthRecommendItemsGetAPIRequest) GetUserTypeCode() string {
	return r._userTypeCode
}

// SetPageIndex is PageIndex Setter
// page index,start from 1
func (r *AliexpressUsergrowthRecommendItemsGetAPIRequest) SetPageIndex(_pageIndex string) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r AliexpressUsergrowthRecommendItemsGetAPIRequest) GetPageIndex() string {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// page size
func (r *AliexpressUsergrowthRecommendItemsGetAPIRequest) SetPageSize(_pageSize string) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AliexpressUsergrowthRecommendItemsGetAPIRequest) GetPageSize() string {
	return r._pageSize
}

// SetCountryCode is CountryCode Setter
// country code
func (r *AliexpressUsergrowthRecommendItemsGetAPIRequest) SetCountryCode(_countryCode string) error {
	r._countryCode = _countryCode
	r.Set("country_code", _countryCode)
	return nil
}

// GetCountryCode CountryCode Getter
func (r AliexpressUsergrowthRecommendItemsGetAPIRequest) GetCountryCode() string {
	return r._countryCode
}
