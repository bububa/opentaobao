package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssocialitemrankingAPIRequest 社交排行榜 API请求
// aliexpress.social.item.ranking
//
// 社交商品成交排行榜
type AliexpresssocialitemrankingAPIRequest struct {
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

// NewAliexpresssocialitemrankingRequest 初始化AliexpresssocialitemrankingAPIRequest对象
func NewAliexpresssocialitemrankingRequest() *AliexpresssocialitemrankingAPIRequest {
	return &AliexpresssocialitemrankingAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssocialitemrankingAPIRequest) GetApiMethodName() string {
	return "aliexpress.social.item.ranking"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssocialitemrankingAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssocialitemrankingAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCountryList is CountryList Setter
// 国家列表
func (r *AliexpresssocialitemrankingAPIRequest) SetCountryList(_countryList []string) error {
	r._countryList = _countryList
	r.Set("country_list", _countryList)
	return nil
}

// GetCountryList CountryList Getter
func (r AliexpresssocialitemrankingAPIRequest) GetCountryList() []string {
	return r._countryList
}

// SetCurrency is Currency Setter
// 币种
func (r *AliexpresssocialitemrankingAPIRequest) SetCurrency(_currency string) error {
	r._currency = _currency
	r.Set("currency", _currency)
	return nil
}

// GetCurrency Currency Getter
func (r AliexpresssocialitemrankingAPIRequest) GetCurrency() string {
	return r._currency
}

// SetLocale is Locale Setter
// locale，格式为language+&#34;_&#34;+country
func (r *AliexpresssocialitemrankingAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// GetLocale Locale Getter
func (r AliexpresssocialitemrankingAPIRequest) GetLocale() string {
	return r._locale
}

// SetPageNo is PageNo Setter
// 页码
func (r *AliexpresssocialitemrankingAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r AliexpresssocialitemrankingAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetCateId is CateId Setter
// 类目ID
func (r *AliexpresssocialitemrankingAPIRequest) SetCateId(_cateId int64) error {
	r._cateId = _cateId
	r.Set("cate_id", _cateId)
	return nil
}

// GetCateId CateId Getter
func (r AliexpresssocialitemrankingAPIRequest) GetCateId() int64 {
	return r._cateId
}

// SetPageSize is PageSize Setter
// 每页条数
func (r *AliexpresssocialitemrankingAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AliexpresssocialitemrankingAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
