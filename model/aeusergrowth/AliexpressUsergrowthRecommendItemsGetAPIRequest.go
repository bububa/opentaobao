package aeusergrowth

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressusergrowthrecommenditemsgetAPIRequest 第三方平台推荐商品 API请求
// aliexpress.usergrowth.recommend.items.get
//
// 第三方平台的推荐AE商品   场景：skin 、底部推荐等
type AliexpressusergrowthrecommenditemsgetAPIRequest struct {
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

// NewAliexpressusergrowthrecommenditemsgetRequest 初始化AliexpressusergrowthrecommenditemsgetAPIRequest对象
func NewAliexpressusergrowthrecommenditemsgetRequest() *AliexpressusergrowthrecommenditemsgetAPIRequest {
	return &AliexpressusergrowthrecommenditemsgetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressusergrowthrecommenditemsgetAPIRequest) GetApiMethodName() string {
	return "aliexpress.usergrowth.recommend.items.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressusergrowthrecommenditemsgetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressusergrowthrecommenditemsgetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTrackingId is TrackingId Setter
// third party trackingId
func (r *AliexpressusergrowthrecommenditemsgetAPIRequest) SetTrackingId(_trackingId string) error {
	r._trackingId = _trackingId
	r.Set("tracking_id", _trackingId)
	return nil
}

// GetTrackingId TrackingId Getter
func (r AliexpressusergrowthrecommenditemsgetAPIRequest) GetTrackingId() string {
	return r._trackingId
}

// SetCurrencyCode is CurrencyCode Setter
// currency Code
func (r *AliexpressusergrowthrecommenditemsgetAPIRequest) SetCurrencyCode(_currencyCode string) error {
	r._currencyCode = _currencyCode
	r.Set("currency_code", _currencyCode)
	return nil
}

// GetCurrencyCode CurrencyCode Getter
func (r AliexpressusergrowthrecommenditemsgetAPIRequest) GetCurrencyCode() string {
	return r._currencyCode
}

// SetLanguage is Language Setter
// language
func (r *AliexpressusergrowthrecommenditemsgetAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// GetLanguage Language Getter
func (r AliexpressusergrowthrecommenditemsgetAPIRequest) GetLanguage() string {
	return r._language
}

// SetUserTypeCode is UserTypeCode Setter
// user type
func (r *AliexpressusergrowthrecommenditemsgetAPIRequest) SetUserTypeCode(_userTypeCode string) error {
	r._userTypeCode = _userTypeCode
	r.Set("user_type_code", _userTypeCode)
	return nil
}

// GetUserTypeCode UserTypeCode Getter
func (r AliexpressusergrowthrecommenditemsgetAPIRequest) GetUserTypeCode() string {
	return r._userTypeCode
}

// SetPageIndex is PageIndex Setter
// page index,start from 1
func (r *AliexpressusergrowthrecommenditemsgetAPIRequest) SetPageIndex(_pageIndex string) error {
	r._pageIndex = _pageIndex
	r.Set("page_index", _pageIndex)
	return nil
}

// GetPageIndex PageIndex Getter
func (r AliexpressusergrowthrecommenditemsgetAPIRequest) GetPageIndex() string {
	return r._pageIndex
}

// SetPageSize is PageSize Setter
// page size
func (r *AliexpressusergrowthrecommenditemsgetAPIRequest) SetPageSize(_pageSize string) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AliexpressusergrowthrecommenditemsgetAPIRequest) GetPageSize() string {
	return r._pageSize
}

// SetCountryCode is CountryCode Setter
// country code
func (r *AliexpressusergrowthrecommenditemsgetAPIRequest) SetCountryCode(_countryCode string) error {
	r._countryCode = _countryCode
	r.Set("country_code", _countryCode)
	return nil
}

// GetCountryCode CountryCode Getter
func (r AliexpressusergrowthrecommenditemsgetAPIRequest) GetCountryCode() string {
	return r._countryCode
}
