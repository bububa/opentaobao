package aeusergrowth

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
第三方平台推荐商品 API请求
aliexpress.usergrowth.recommend.items.get

第三方平台的推荐AE商品   场景：skin 、底部推荐等
*/
type AliexpressUsergrowthRecommendItemsGetRequest struct {
    model.Params
    // third party trackingId
    _trackingId   string
    // currency Code
    _currencyCode   string
    // language
    _language   string
    // user type
    _userTypeCode   string
    // page index,start from 1
    _pageIndex   string
    // page size
    _pageSize   string
    // country code
    _countryCode   string
}

// 初始化AliexpressUsergrowthRecommendItemsGetRequest对象
func NewAliexpressUsergrowthRecommendItemsGetRequest() *AliexpressUsergrowthRecommendItemsGetRequest{
    return &AliexpressUsergrowthRecommendItemsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressUsergrowthRecommendItemsGetRequest) GetApiMethodName() string {
    return "aliexpress.usergrowth.recommend.items.get"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressUsergrowthRecommendItemsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TrackingId Setter
// third party trackingId
func (r *AliexpressUsergrowthRecommendItemsGetRequest) SetTrackingId(_trackingId string) error {
    r._trackingId = _trackingId
    r.Set("tracking_id", _trackingId)
    return nil
}

// TrackingId Getter
func (r AliexpressUsergrowthRecommendItemsGetRequest) GetTrackingId() string {
    return r._trackingId
}
// CurrencyCode Setter
// currency Code
func (r *AliexpressUsergrowthRecommendItemsGetRequest) SetCurrencyCode(_currencyCode string) error {
    r._currencyCode = _currencyCode
    r.Set("currency_code", _currencyCode)
    return nil
}

// CurrencyCode Getter
func (r AliexpressUsergrowthRecommendItemsGetRequest) GetCurrencyCode() string {
    return r._currencyCode
}
// Language Setter
// language
func (r *AliexpressUsergrowthRecommendItemsGetRequest) SetLanguage(_language string) error {
    r._language = _language
    r.Set("language", _language)
    return nil
}

// Language Getter
func (r AliexpressUsergrowthRecommendItemsGetRequest) GetLanguage() string {
    return r._language
}
// UserTypeCode Setter
// user type
func (r *AliexpressUsergrowthRecommendItemsGetRequest) SetUserTypeCode(_userTypeCode string) error {
    r._userTypeCode = _userTypeCode
    r.Set("user_type_code", _userTypeCode)
    return nil
}

// UserTypeCode Getter
func (r AliexpressUsergrowthRecommendItemsGetRequest) GetUserTypeCode() string {
    return r._userTypeCode
}
// PageIndex Setter
// page index,start from 1
func (r *AliexpressUsergrowthRecommendItemsGetRequest) SetPageIndex(_pageIndex string) error {
    r._pageIndex = _pageIndex
    r.Set("page_index", _pageIndex)
    return nil
}

// PageIndex Getter
func (r AliexpressUsergrowthRecommendItemsGetRequest) GetPageIndex() string {
    return r._pageIndex
}
// PageSize Setter
// page size
func (r *AliexpressUsergrowthRecommendItemsGetRequest) SetPageSize(_pageSize string) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AliexpressUsergrowthRecommendItemsGetRequest) GetPageSize() string {
    return r._pageSize
}
// CountryCode Setter
// country code
func (r *AliexpressUsergrowthRecommendItemsGetRequest) SetCountryCode(_countryCode string) error {
    r._countryCode = _countryCode
    r.Set("country_code", _countryCode)
    return nil
}

// CountryCode Getter
func (r AliexpressUsergrowthRecommendItemsGetRequest) GetCountryCode() string {
    return r._countryCode
}
