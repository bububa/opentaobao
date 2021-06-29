package aeusergrowth

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
第三方平台推荐商品 APIRequest
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

func NewAliexpressUsergrowthRecommendItemsGetRequest() *AliexpressUsergrowthRecommendItemsGetRequest{
    return &AliexpressUsergrowthRecommendItemsGetRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressUsergrowthRecommendItemsGetRequest) GetApiMethodName() string {
    return "aliexpress.usergrowth.recommend.items.get"
}

func (r AliexpressUsergrowthRecommendItemsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressUsergrowthRecommendItemsGetRequest) SetTrackingId(trackingId string) error {
    r.trackingId = trackingId
    r.Set("tracking_id", trackingId)
    return nil
}

func (r AliexpressUsergrowthRecommendItemsGetRequest) GetTrackingId() string {
    return r.trackingId
}

func (r *AliexpressUsergrowthRecommendItemsGetRequest) SetCurrencyCode(currencyCode string) error {
    r.currencyCode = currencyCode
    r.Set("currency_code", currencyCode)
    return nil
}

func (r AliexpressUsergrowthRecommendItemsGetRequest) GetCurrencyCode() string {
    return r.currencyCode
}

func (r *AliexpressUsergrowthRecommendItemsGetRequest) SetLanguage(language string) error {
    r.language = language
    r.Set("language", language)
    return nil
}

func (r AliexpressUsergrowthRecommendItemsGetRequest) GetLanguage() string {
    return r.language
}

func (r *AliexpressUsergrowthRecommendItemsGetRequest) SetUserTypeCode(userTypeCode string) error {
    r.userTypeCode = userTypeCode
    r.Set("user_type_code", userTypeCode)
    return nil
}

func (r AliexpressUsergrowthRecommendItemsGetRequest) GetUserTypeCode() string {
    return r.userTypeCode
}

func (r *AliexpressUsergrowthRecommendItemsGetRequest) SetPageIndex(pageIndex string) error {
    r.pageIndex = pageIndex
    r.Set("page_index", pageIndex)
    return nil
}

func (r AliexpressUsergrowthRecommendItemsGetRequest) GetPageIndex() string {
    return r.pageIndex
}

func (r *AliexpressUsergrowthRecommendItemsGetRequest) SetPageSize(pageSize string) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AliexpressUsergrowthRecommendItemsGetRequest) GetPageSize() string {
    return r.pageSize
}

func (r *AliexpressUsergrowthRecommendItemsGetRequest) SetCountryCode(countryCode string) error {
    r.countryCode = countryCode
    r.Set("country_code", countryCode)
    return nil
}

func (r AliexpressUsergrowthRecommendItemsGetRequest) GetCountryCode() string {
    return r.countryCode
}

