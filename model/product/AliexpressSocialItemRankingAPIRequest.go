package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSocialItemRankingAPIRequest 社交排行榜 API请求
// aliexpress.social.item.ranking
//
// 社交商品成交排行榜
type AliexpressSocialItemRankingAPIRequest struct {
	model.Params
	// 国家列表
	_countryList []string
	// 币种
	_currency string
	// locale，格式为language+"_"+country
	_locale string
	// 页码
	_pageNo int64
	// 类目ID
	_cateId int64
	// 每页条数
	_pageSize int64
}

// NewAliexpressSocialItemRankingRequest 初始化AliexpressSocialItemRankingAPIRequest对象
func NewAliexpressSocialItemRankingRequest() *AliexpressSocialItemRankingAPIRequest {
	return &AliexpressSocialItemRankingAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSocialItemRankingAPIRequest) GetApiMethodName() string {
	return "aliexpress.social.item.ranking"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSocialItemRankingAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressSocialItemRankingAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCountryList is CountryList Setter
// 国家列表
func (r *AliexpressSocialItemRankingAPIRequest) SetCountryList(_countryList []string) error {
	r._countryList = _countryList
	r.Set("country_list", _countryList)
	return nil
}

// GetCountryList CountryList Getter
func (r AliexpressSocialItemRankingAPIRequest) GetCountryList() []string {
	return r._countryList
}

// SetCurrency is Currency Setter
// 币种
func (r *AliexpressSocialItemRankingAPIRequest) SetCurrency(_currency string) error {
	r._currency = _currency
	r.Set("currency", _currency)
	return nil
}

// GetCurrency Currency Getter
func (r AliexpressSocialItemRankingAPIRequest) GetCurrency() string {
	return r._currency
}

// SetLocale is Locale Setter
// locale，格式为language+&#34;_&#34;+country
func (r *AliexpressSocialItemRankingAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// GetLocale Locale Getter
func (r AliexpressSocialItemRankingAPIRequest) GetLocale() string {
	return r._locale
}

// SetPageNo is PageNo Setter
// 页码
func (r *AliexpressSocialItemRankingAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r AliexpressSocialItemRankingAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetCateId is CateId Setter
// 类目ID
func (r *AliexpressSocialItemRankingAPIRequest) SetCateId(_cateId int64) error {
	r._cateId = _cateId
	r.Set("cate_id", _cateId)
	return nil
}

// GetCateId CateId Getter
func (r AliexpressSocialItemRankingAPIRequest) GetCateId() int64 {
	return r._cateId
}

// SetPageSize is PageSize Setter
// 每页条数
func (r *AliexpressSocialItemRankingAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AliexpressSocialItemRankingAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
