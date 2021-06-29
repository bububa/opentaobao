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
    currency   string
    // 国家列表
    countryList   []string
    // locale，格式为language+"_"+country
    locale   string
    // 页码
    pageNo   int64
    // 类目ID
    cateId   int64
    // 每页条数
    pageSize   int64
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
func (r *AliexpressSocialItemRankingRequest) SetCurrency(currency string) error {
    r.currency = currency
    r.Set("currency", currency)
    return nil
}

// Currency Getter
func (r AliexpressSocialItemRankingRequest) GetCurrency() string {
    return r.currency
}
// CountryList Setter
// 国家列表
func (r *AliexpressSocialItemRankingRequest) SetCountryList(countryList []string) error {
    r.countryList = countryList
    r.Set("country_list", countryList)
    return nil
}

// CountryList Getter
func (r AliexpressSocialItemRankingRequest) GetCountryList() []string {
    return r.countryList
}
// Locale Setter
// locale，格式为language+"_"+country
func (r *AliexpressSocialItemRankingRequest) SetLocale(locale string) error {
    r.locale = locale
    r.Set("locale", locale)
    return nil
}

// Locale Getter
func (r AliexpressSocialItemRankingRequest) GetLocale() string {
    return r.locale
}
// PageNo Setter
// 页码
func (r *AliexpressSocialItemRankingRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r AliexpressSocialItemRankingRequest) GetPageNo() int64 {
    return r.pageNo
}
// CateId Setter
// 类目ID
func (r *AliexpressSocialItemRankingRequest) SetCateId(cateId int64) error {
    r.cateId = cateId
    r.Set("cate_id", cateId)
    return nil
}

// CateId Getter
func (r AliexpressSocialItemRankingRequest) GetCateId() int64 {
    return r.cateId
}
// PageSize Setter
// 每页条数
func (r *AliexpressSocialItemRankingRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AliexpressSocialItemRankingRequest) GetPageSize() int64 {
    return r.pageSize
}
