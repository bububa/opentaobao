package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
社交排行榜 API请求
aliexpress.social.item.ranking

社交商品成交排行榜
*/
type AliexpressSocialItemRankingAPIRequest struct {
    model.Params
    // 币种
    _currency   string
    // 国家列表
    _countryList   []string
    // locale，格式为language+"_"+country
    _locale   string
    // 页码
    _pageNo   int64
    // 类目ID
    _cateId   int64
    // 每页条数
    _pageSize   int64
}

// 初始化AliexpressSocialItemRankingAPIRequest对象
func NewAliexpressSocialItemRankingRequest() *AliexpressSocialItemRankingAPIRequest{
    return &AliexpressSocialItemRankingAPIRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSocialItemRankingAPIRequest) GetApiMethodName() string {
    return "aliexpress.social.item.ranking"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSocialItemRankingAPIRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Currency Setter
// 币种
func (r *AliexpressSocialItemRankingAPIRequest) SetCurrency(_currency string) error {
    r._currency = _currency
    r.Set("currency", _currency)
    return nil
}

// Currency Getter
func (r AliexpressSocialItemRankingAPIRequest) GetCurrency() string {
    return r._currency
}
// CountryList Setter
// 国家列表
func (r *AliexpressSocialItemRankingAPIRequest) SetCountryList(_countryList []string) error {
    r._countryList = _countryList
    r.Set("country_list", _countryList)
    return nil
}

// CountryList Getter
func (r AliexpressSocialItemRankingAPIRequest) GetCountryList() []string {
    return r._countryList
}
// Locale Setter
// locale，格式为language+"_"+country
func (r *AliexpressSocialItemRankingAPIRequest) SetLocale(_locale string) error {
    r._locale = _locale
    r.Set("locale", _locale)
    return nil
}

// Locale Getter
func (r AliexpressSocialItemRankingAPIRequest) GetLocale() string {
    return r._locale
}
// PageNo Setter
// 页码
func (r *AliexpressSocialItemRankingAPIRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r AliexpressSocialItemRankingAPIRequest) GetPageNo() int64 {
    return r._pageNo
}
// CateId Setter
// 类目ID
func (r *AliexpressSocialItemRankingAPIRequest) SetCateId(_cateId int64) error {
    r._cateId = _cateId
    r.Set("cate_id", _cateId)
    return nil
}

// CateId Getter
func (r AliexpressSocialItemRankingAPIRequest) GetCateId() int64 {
    return r._cateId
}
// PageSize Setter
// 每页条数
func (r *AliexpressSocialItemRankingAPIRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AliexpressSocialItemRankingAPIRequest) GetPageSize() int64 {
    return r._pageSize
}
