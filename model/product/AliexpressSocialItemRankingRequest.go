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
type AliexpressSocialItemRankingRequest struct {
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

// 初始化AliexpressSocialItemRankingRequest对象
func NewAliexpressSocialItemRankingRequest() *AliexpressSocialItemRankingRequest{
    return &AliexpressSocialItemRankingRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressSocialItemRankingRequest) GetApiMethodName() string {
    return "aliexpress.social.item.ranking"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressSocialItemRankingRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// Currency Setter
// 币种
func (r *AliexpressSocialItemRankingRequest) SetCurrency(_currency string) error {
    r._currency = _currency
    r.Set("currency", _currency)
    return nil
}

// Currency Getter
func (r AliexpressSocialItemRankingRequest) GetCurrency() string {
    return r._currency
}
// CountryList Setter
// 国家列表
func (r *AliexpressSocialItemRankingRequest) SetCountryList(_countryList []string) error {
    r._countryList = _countryList
    r.Set("country_list", _countryList)
    return nil
}

// CountryList Getter
func (r AliexpressSocialItemRankingRequest) GetCountryList() []string {
    return r._countryList
}
// Locale Setter
// locale，格式为language+"_"+country
func (r *AliexpressSocialItemRankingRequest) SetLocale(_locale string) error {
    r._locale = _locale
    r.Set("locale", _locale)
    return nil
}

// Locale Getter
func (r AliexpressSocialItemRankingRequest) GetLocale() string {
    return r._locale
}
// PageNo Setter
// 页码
func (r *AliexpressSocialItemRankingRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r AliexpressSocialItemRankingRequest) GetPageNo() int64 {
    return r._pageNo
}
// CateId Setter
// 类目ID
func (r *AliexpressSocialItemRankingRequest) SetCateId(_cateId int64) error {
    r._cateId = _cateId
    r.Set("cate_id", _cateId)
    return nil
}

// CateId Getter
func (r AliexpressSocialItemRankingRequest) GetCateId() int64 {
    return r._cateId
}
// PageSize Setter
// 每页条数
func (r *AliexpressSocialItemRankingRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AliexpressSocialItemRankingRequest) GetPageSize() int64 {
    return r._pageSize
}
