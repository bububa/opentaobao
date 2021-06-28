package product

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
社交排行榜 APIRequest
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

func NewAliexpressSocialItemRankingRequest() *AliexpressSocialItemRankingRequest{
    return &AliexpressSocialItemRankingRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressSocialItemRankingRequest) GetApiMethodName() string {
    return "aliexpress.social.item.ranking"
}

func (r AliexpressSocialItemRankingRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressSocialItemRankingRequest) SetCurrency(currency string) error {
    r.currency = currency
    r.Set("currency", currency)
    return nil
}

func (r AliexpressSocialItemRankingRequest) GetCurrency() string {
    return r.currency
}

func (r *AliexpressSocialItemRankingRequest) SetCountryList(countryList []string) error {
    r.countryList = countryList
    r.Set("country_list", countryList)
    return nil
}

func (r AliexpressSocialItemRankingRequest) GetCountryList() []string {
    return r.countryList
}

func (r *AliexpressSocialItemRankingRequest) SetLocale(locale string) error {
    r.locale = locale
    r.Set("locale", locale)
    return nil
}

func (r AliexpressSocialItemRankingRequest) GetLocale() string {
    return r.locale
}

func (r *AliexpressSocialItemRankingRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r AliexpressSocialItemRankingRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *AliexpressSocialItemRankingRequest) SetCateId(cateId int64) error {
    r.cateId = cateId
    r.Set("cate_id", cateId)
    return nil
}

func (r AliexpressSocialItemRankingRequest) GetCateId() int64 {
    return r.cateId
}

func (r *AliexpressSocialItemRankingRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AliexpressSocialItemRankingRequest) GetPageSize() int64 {
    return r.pageSize
}

