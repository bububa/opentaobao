package aeusergrowth

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
第三方平台搜索AE商品 APIRequest
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

func NewAliexpressUsergrowthSearchItemsGetRequest() *AliexpressUsergrowthSearchItemsGetRequest{
    return &AliexpressUsergrowthSearchItemsGetRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressUsergrowthSearchItemsGetRequest) GetApiMethodName() string {
    return "aliexpress.usergrowth.search.items.get"
}

func (r AliexpressUsergrowthSearchItemsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressUsergrowthSearchItemsGetRequest) SetKeywords(keywords string) error {
    r.keywords = keywords
    r.Set("keywords", keywords)
    return nil
}

func (r AliexpressUsergrowthSearchItemsGetRequest) GetKeywords() string {
    return r.keywords
}

func (r *AliexpressUsergrowthSearchItemsGetRequest) SetTrackingId(trackingId string) error {
    r.trackingId = trackingId
    r.Set("tracking_id", trackingId)
    return nil
}

func (r AliexpressUsergrowthSearchItemsGetRequest) GetTrackingId() string {
    return r.trackingId
}

func (r *AliexpressUsergrowthSearchItemsGetRequest) SetCurrencyCode(currencyCode string) error {
    r.currencyCode = currencyCode
    r.Set("currency_code", currencyCode)
    return nil
}

func (r AliexpressUsergrowthSearchItemsGetRequest) GetCurrencyCode() string {
    return r.currencyCode
}

func (r *AliexpressUsergrowthSearchItemsGetRequest) SetLanguage(language string) error {
    r.language = language
    r.Set("language", language)
    return nil
}

func (r AliexpressUsergrowthSearchItemsGetRequest) GetLanguage() string {
    return r.language
}

func (r *AliexpressUsergrowthSearchItemsGetRequest) SetPageSize(pageSize string) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AliexpressUsergrowthSearchItemsGetRequest) GetPageSize() string {
    return r.pageSize
}

func (r *AliexpressUsergrowthSearchItemsGetRequest) SetPageIndex(pageIndex string) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

func (r AliexpressUsergrowthSearchItemsGetRequest) GetPageIndex() string {
    return r.pageIndex
}

func (r *AliexpressUsergrowthSearchItemsGetRequest) SetCountryCode(countryCode string) error {
    r.countryCode = countryCode
    r.Set("country_code", countryCode)
    return nil
}

func (r AliexpressUsergrowthSearchItemsGetRequest) GetCountryCode() string {
    return r.countryCode
}

