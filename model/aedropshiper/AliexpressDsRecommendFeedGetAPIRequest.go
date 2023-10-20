package aedropshiper

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressDsRecommendFeedGetAPIRequest 获取推荐商品信息流接口 API请求
// aliexpress.ds.recommend.feed.get
//
// 获取推荐商品信息流
type AliexpressDsRecommendFeedGetAPIRequest struct {
	model.Params
	// screens the subject product library for the target country
	_country string
	// target currency:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
	_targetCurrency string
	// target language:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IN
	_targetLanguage string
	// sort by：priceAsc，priceDesc，volumeAsc、volumeDesc, discountAsc, discountDesc, DSRratingAsc，DSRratingDesc,
	_sort string
	// Category ID, you can get category ID via "get category" API https://developers.aliexpress.com/en/doc.htm?docId=45801&docType=2
	_categoryId string
	// feed name, eg. "DS bestseller"
	_feedName string
	// record count of each page, 1 - 50
	_pageSize int64
	// Page number
	_pageNo int64
}

// NewAliexpressDsRecommendFeedGetRequest 初始化AliexpressDsRecommendFeedGetAPIRequest对象
func NewAliexpressDsRecommendFeedGetRequest() *AliexpressDsRecommendFeedGetAPIRequest {
	return &AliexpressDsRecommendFeedGetAPIRequest{
		Params: model.NewParams(8),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressDsRecommendFeedGetAPIRequest) Reset() {
	r._country = ""
	r._targetCurrency = ""
	r._targetLanguage = ""
	r._sort = ""
	r._categoryId = ""
	r._feedName = ""
	r._pageSize = 0
	r._pageNo = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressDsRecommendFeedGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.ds.recommend.feed.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressDsRecommendFeedGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressDsRecommendFeedGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCountry is Country Setter
// screens the subject product library for the target country
func (r *AliexpressDsRecommendFeedGetAPIRequest) SetCountry(_country string) error {
	r._country = _country
	r.Set("country", _country)
	return nil
}

// GetCountry Country Getter
func (r AliexpressDsRecommendFeedGetAPIRequest) GetCountry() string {
	return r._country
}

// SetTargetCurrency is TargetCurrency Setter
// target currency:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
func (r *AliexpressDsRecommendFeedGetAPIRequest) SetTargetCurrency(_targetCurrency string) error {
	r._targetCurrency = _targetCurrency
	r.Set("target_currency", _targetCurrency)
	return nil
}

// GetTargetCurrency TargetCurrency Getter
func (r AliexpressDsRecommendFeedGetAPIRequest) GetTargetCurrency() string {
	return r._targetCurrency
}

// SetTargetLanguage is TargetLanguage Setter
// target language:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IN
func (r *AliexpressDsRecommendFeedGetAPIRequest) SetTargetLanguage(_targetLanguage string) error {
	r._targetLanguage = _targetLanguage
	r.Set("target_language", _targetLanguage)
	return nil
}

// GetTargetLanguage TargetLanguage Getter
func (r AliexpressDsRecommendFeedGetAPIRequest) GetTargetLanguage() string {
	return r._targetLanguage
}

// SetSort is Sort Setter
// sort by：priceAsc，priceDesc，volumeAsc、volumeDesc, discountAsc, discountDesc, DSRratingAsc，DSRratingDesc,
func (r *AliexpressDsRecommendFeedGetAPIRequest) SetSort(_sort string) error {
	r._sort = _sort
	r.Set("sort", _sort)
	return nil
}

// GetSort Sort Getter
func (r AliexpressDsRecommendFeedGetAPIRequest) GetSort() string {
	return r._sort
}

// SetCategoryId is CategoryId Setter
// Category ID, you can get category ID via &#34;get category&#34; API https://developers.aliexpress.com/en/doc.htm?docId=45801&amp;docType=2
func (r *AliexpressDsRecommendFeedGetAPIRequest) SetCategoryId(_categoryId string) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r AliexpressDsRecommendFeedGetAPIRequest) GetCategoryId() string {
	return r._categoryId
}

// SetFeedName is FeedName Setter
// feed name, eg. &#34;DS bestseller&#34;
func (r *AliexpressDsRecommendFeedGetAPIRequest) SetFeedName(_feedName string) error {
	r._feedName = _feedName
	r.Set("feed_name", _feedName)
	return nil
}

// GetFeedName FeedName Getter
func (r AliexpressDsRecommendFeedGetAPIRequest) GetFeedName() string {
	return r._feedName
}

// SetPageSize is PageSize Setter
// record count of each page, 1 - 50
func (r *AliexpressDsRecommendFeedGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AliexpressDsRecommendFeedGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

// SetPageNo is PageNo Setter
// Page number
func (r *AliexpressDsRecommendFeedGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r AliexpressDsRecommendFeedGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

var poolAliexpressDsRecommendFeedGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressDsRecommendFeedGetRequest()
	},
}

// GetAliexpressDsRecommendFeedGetRequest 从 sync.Pool 获取 AliexpressDsRecommendFeedGetAPIRequest
func GetAliexpressDsRecommendFeedGetAPIRequest() *AliexpressDsRecommendFeedGetAPIRequest {
	return poolAliexpressDsRecommendFeedGetAPIRequest.Get().(*AliexpressDsRecommendFeedGetAPIRequest)
}

// ReleaseAliexpressDsRecommendFeedGetAPIRequest 将 AliexpressDsRecommendFeedGetAPIRequest 放入 sync.Pool
func ReleaseAliexpressDsRecommendFeedGetAPIRequest(v *AliexpressDsRecommendFeedGetAPIRequest) {
	v.Reset()
	poolAliexpressDsRecommendFeedGetAPIRequest.Put(v)
}
