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
    trackingId   string
    // currency Code
    currencyCode   string
    // language
    language   string
    // user type
    userTypeCode   string
    // page index,start from 1
    pageIndex   string
    // page size
    pageSize   string
    // country code
    countryCode   string
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
func (r *AliexpressUsergrowthRecommendItemsGetRequest) SetTrackingId(trackingId string) error {
    r.trackingId = trackingId
    r.Set("tracking_id", trackingId)
    return nil
}

// TrackingId Getter
func (r AliexpressUsergrowthRecommendItemsGetRequest) GetTrackingId() string {
    return r.trackingId
}
// CurrencyCode Setter
// currency Code
func (r *AliexpressUsergrowthRecommendItemsGetRequest) SetCurrencyCode(currencyCode string) error {
    r.currencyCode = currencyCode
    r.Set("currency_code", currencyCode)
    return nil
}

// CurrencyCode Getter
func (r AliexpressUsergrowthRecommendItemsGetRequest) GetCurrencyCode() string {
    return r.currencyCode
}
// Language Setter
// language
func (r *AliexpressUsergrowthRecommendItemsGetRequest) SetLanguage(language string) error {
    r.language = language
    r.Set("language", language)
    return nil
}

// Language Getter
func (r AliexpressUsergrowthRecommendItemsGetRequest) GetLanguage() string {
    return r.language
}
// UserTypeCode Setter
// user type
func (r *AliexpressUsergrowthRecommendItemsGetRequest) SetUserTypeCode(userTypeCode string) error {
    r.userTypeCode = userTypeCode
    r.Set("user_type_code", userTypeCode)
    return nil
}

// UserTypeCode Getter
func (r AliexpressUsergrowthRecommendItemsGetRequest) GetUserTypeCode() string {
    return r.userTypeCode
}
// PageIndex Setter
// page index,start from 1
func (r *AliexpressUsergrowthRecommendItemsGetRequest) SetPageIndex(pageIndex string) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

// PageIndex Getter
func (r AliexpressUsergrowthRecommendItemsGetRequest) GetPageIndex() string {
    return r.pageIndex
}
// PageSize Setter
// page size
func (r *AliexpressUsergrowthRecommendItemsGetRequest) SetPageSize(pageSize string) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AliexpressUsergrowthRecommendItemsGetRequest) GetPageSize() string {
    return r.pageSize
}
// CountryCode Setter
// country code
func (r *AliexpressUsergrowthRecommendItemsGetRequest) SetCountryCode(countryCode string) error {
    r.countryCode = countryCode
    r.Set("country_code", countryCode)
    return nil
}

// CountryCode Getter
func (r AliexpressUsergrowthRecommendItemsGetRequest) GetCountryCode() string {
    return r.countryCode
}
