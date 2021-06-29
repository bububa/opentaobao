package aeusergrowth

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
第三方平台搜索AE商品 API请求
aliexpress.usergrowth.search.items.get

第三方平台的搜索服务   获取AE商品list
*/
type AliexpressUsergrowthSearchItemsGetRequest struct {
    model.Params
    // user input keypods
    keywords   string
    // Third party tracking_id
    trackingId   string
    // currency code
    currencyCode   string
    // language
    language   string
    // page size
    pageSize   string
    // page index
    pageIndex   string
    // ship to country
    countryCode   string
}

// 初始化AliexpressUsergrowthSearchItemsGetRequest对象
func NewAliexpressUsergrowthSearchItemsGetRequest() *AliexpressUsergrowthSearchItemsGetRequest{
    return &AliexpressUsergrowthSearchItemsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressUsergrowthSearchItemsGetRequest) GetApiMethodName() string {
    return "aliexpress.usergrowth.search.items.get"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressUsergrowthSearchItemsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Keywords Setter
// user input keypods
func (r *AliexpressUsergrowthSearchItemsGetRequest) SetKeywords(keywords string) error {
    r.keywords = keywords
    r.Set("keywords", keywords)
    return nil
}

// Keywords Getter
func (r AliexpressUsergrowthSearchItemsGetRequest) GetKeywords() string {
    return r.keywords
}
// TrackingId Setter
// Third party tracking_id
func (r *AliexpressUsergrowthSearchItemsGetRequest) SetTrackingId(trackingId string) error {
    r.trackingId = trackingId
    r.Set("tracking_id", trackingId)
    return nil
}

// TrackingId Getter
func (r AliexpressUsergrowthSearchItemsGetRequest) GetTrackingId() string {
    return r.trackingId
}
// CurrencyCode Setter
// currency code
func (r *AliexpressUsergrowthSearchItemsGetRequest) SetCurrencyCode(currencyCode string) error {
    r.currencyCode = currencyCode
    r.Set("currency_code", currencyCode)
    return nil
}

// CurrencyCode Getter
func (r AliexpressUsergrowthSearchItemsGetRequest) GetCurrencyCode() string {
    return r.currencyCode
}
// Language Setter
// language
func (r *AliexpressUsergrowthSearchItemsGetRequest) SetLanguage(language string) error {
    r.language = language
    r.Set("language", language)
    return nil
}

// Language Getter
func (r AliexpressUsergrowthSearchItemsGetRequest) GetLanguage() string {
    return r.language
}
// PageSize Setter
// page size
func (r *AliexpressUsergrowthSearchItemsGetRequest) SetPageSize(pageSize string) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AliexpressUsergrowthSearchItemsGetRequest) GetPageSize() string {
    return r.pageSize
}
// PageIndex Setter
// page index
func (r *AliexpressUsergrowthSearchItemsGetRequest) SetPageIndex(pageIndex string) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

// PageIndex Getter
func (r AliexpressUsergrowthSearchItemsGetRequest) GetPageIndex() string {
    return r.pageIndex
}
// CountryCode Setter
// ship to country
func (r *AliexpressUsergrowthSearchItemsGetRequest) SetCountryCode(countryCode string) error {
    r.countryCode = countryCode
    r.Set("country_code", countryCode)
    return nil
}

// CountryCode Getter
func (r AliexpressUsergrowthSearchItemsGetRequest) GetCountryCode() string {
    return r.countryCode
}
