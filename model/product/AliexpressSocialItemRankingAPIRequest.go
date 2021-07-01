package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AliexpressSocialItemRankingAPIRequest
社交排行榜 API请求
aliexpress.social.item.ranking

社交商品成交排行榜 */
type AliexpressSocialItemRankingAPIRequest struct {
	model.Params
	// 币种
	_currency string
	// 国家列表
	_countryList []string
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
func (r AliexpressSocialItemRankingAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is Currency Setter
// 币种
func (r *AliexpressSocialItemRankingAPIRequest) SetCurrency(_currency string) error {
	r._currency = _currency
	r.Set("currency", _currency)
	return nil
}

// Get Currency Getter
func (r AliexpressSocialItemRankingAPIRequest) GetCurrency() string {
	return r._currency
}

// Set is CountryList Setter
// 国家列表
func (r *AliexpressSocialItemRankingAPIRequest) SetCountryList(_countryList []string) error {
	r._countryList = _countryList
	r.Set("country_list", _countryList)
	return nil
}

// Get CountryList Getter
func (r AliexpressSocialItemRankingAPIRequest) GetCountryList() []string {
	return r._countryList
}

// Set is Locale Setter
// locale，格式为language+"_"+country
func (r *AliexpressSocialItemRankingAPIRequest) SetLocale(_locale string) error {
	r._locale = _locale
	r.Set("locale", _locale)
	return nil
}

// Get Locale Getter
func (r AliexpressSocialItemRankingAPIRequest) GetLocale() string {
	return r._locale
}

// Set is PageNo Setter
// 页码
func (r *AliexpressSocialItemRankingAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// Get PageNo Getter
func (r AliexpressSocialItemRankingAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// Set is CateId Setter
// 类目ID
func (r *AliexpressSocialItemRankingAPIRequest) SetCateId(_cateId int64) error {
	r._cateId = _cateId
	r.Set("cate_id", _cateId)
	return nil
}

// Get CateId Getter
func (r AliexpressSocialItemRankingAPIRequest) GetCateId() int64 {
	return r._cateId
}

// Set is PageSize Setter
// 每页条数
func (r *AliexpressSocialItemRankingAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// Get PageSize Getter
func (r AliexpressSocialItemRankingAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
