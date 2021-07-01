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
type AliexpressUsergrowthSearchItemsGetAPIRequest struct {
    model.Params
    // user input keypods
    _keywords   string
    // Third party tracking_id
    _trackingId   string
    // currency code
    _currencyCode   string
    // language
    _language   string
    // page size
    _pageSize   string
    // page index
    _pageIndex   string
    // ship to country
    _countryCode   string
}

// 初始化AliexpressUsergrowthSearchItemsGetAPIRequest对象
func NewAliexpressUsergrowthSearchItemsGetRequest() *AliexpressUsergrowthSearchItemsGetAPIRequest{
    return &AliexpressUsergrowthSearchItemsGetAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressUsergrowthSearchItemsGetAPIRequest) GetApiMethodName() string {
    return "aliexpress.usergrowth.search.items.get"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressUsergrowthSearchItemsGetAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Keywords Setter
// user input keypods
func (r *AliexpressUsergrowthSearchItemsGetAPIRequest) SetKeywords(_keywords string) error {
    r._keywords = _keywords
    r.Set("keywords", _keywords)
    return nil
}

// Keywords Getter
func (r AliexpressUsergrowthSearchItemsGetAPIRequest) GetKeywords() string {
    return r._keywords
}
// TrackingId Setter
// Third party tracking_id
func (r *AliexpressUsergrowthSearchItemsGetAPIRequest) SetTrackingId(_trackingId string) error {
    r._trackingId = _trackingId
    r.Set("tracking_id", _trackingId)
    return nil
}

// TrackingId Getter
func (r AliexpressUsergrowthSearchItemsGetAPIRequest) GetTrackingId() string {
    return r._trackingId
}
// CurrencyCode Setter
// currency code
func (r *AliexpressUsergrowthSearchItemsGetAPIRequest) SetCurrencyCode(_currencyCode string) error {
    r._currencyCode = _currencyCode
    r.Set("currency_code", _currencyCode)
    return nil
}

// CurrencyCode Getter
func (r AliexpressUsergrowthSearchItemsGetAPIRequest) GetCurrencyCode() string {
    return r._currencyCode
}
// Language Setter
// language
func (r *AliexpressUsergrowthSearchItemsGetAPIRequest) SetLanguage(_language string) error {
    r._language = _language
    r.Set("language", _language)
    return nil
}

// Language Getter
func (r AliexpressUsergrowthSearchItemsGetAPIRequest) GetLanguage() string {
    return r._language
}
// PageSize Setter
// page size
func (r *AliexpressUsergrowthSearchItemsGetAPIRequest) SetPageSize(_pageSize string) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AliexpressUsergrowthSearchItemsGetAPIRequest) GetPageSize() string {
    return r._pageSize
}
// PageIndex Setter
// page index
func (r *AliexpressUsergrowthSearchItemsGetAPIRequest) SetPageIndex(_pageIndex string) error {
    r._pageIndex = _pageIndex
    r.Set("page_index", _pageIndex)
    return nil
}

// PageIndex Getter
func (r AliexpressUsergrowthSearchItemsGetAPIRequest) GetPageIndex() string {
    return r._pageIndex
}
// CountryCode Setter
// ship to country
func (r *AliexpressUsergrowthSearchItemsGetAPIRequest) SetCountryCode(_countryCode string) error {
    r._countryCode = _countryCode
    r.Set("country_code", _countryCode)
    return nil
}

// CountryCode Getter
func (r AliexpressUsergrowthSearchItemsGetAPIRequest) GetCountryCode() string {
    return r._countryCode
}
