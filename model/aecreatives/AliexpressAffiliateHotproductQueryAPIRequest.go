package aecreatives

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAffiliateHotproductQueryAPIRequest 查询联盟爆品数据 API请求
// aliexpress.affiliate.hotproduct.query
//
// 查询联盟爆品API
type AliexpressAffiliateHotproductQueryAPIRequest struct {
	model.Params
	// 请求签名
	_appSignature string
	// 类目ID列表
	_categoryIds string
	// 返回字段列表
	_fields string
	// 关键词
	_keywords string
	// 平台商家类型：ALL,PLAZA,TMALL
	_platformProductType string
	// 排序方式:SALE_PRICE_ASC, SALE_PRICE_DESC, DISCOUNT_ASC, DISCOUNT_DESC, LAST_VOLUME_ASC, LAST_VOLUME_DESC
	_sort string
	// 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
	_targetCurrency string
	// 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
	_targetLanguage string
	// trackingId
	_trackingId string
	// 物流到达时间。3：3日达，5：5 日达，7：7日达，10：10日达
	_deliveryDays string
	// 商品收货国家，根据该国家税率政策返回对应商品价格
	_shipToCountry string
	// 最大售价
	_maxSalePrice int64
	// 最小售价
	_minSalePrice int64
	// 请求页数
	_pageNo int64
	// 每次请求数量
	_pageSize int64
}

// NewAliexpressAffiliateHotproductQueryRequest 初始化AliexpressAffiliateHotproductQueryAPIRequest对象
func NewAliexpressAffiliateHotproductQueryRequest() *AliexpressAffiliateHotproductQueryAPIRequest {
	return &AliexpressAffiliateHotproductQueryAPIRequest{
		Params: model.NewParams(15),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressAffiliateHotproductQueryAPIRequest) Reset() {
	r._appSignature = ""
	r._categoryIds = ""
	r._fields = ""
	r._keywords = ""
	r._platformProductType = ""
	r._sort = ""
	r._targetCurrency = ""
	r._targetLanguage = ""
	r._trackingId = ""
	r._deliveryDays = ""
	r._shipToCountry = ""
	r._maxSalePrice = 0
	r._minSalePrice = 0
	r._pageNo = 0
	r._pageSize = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAffiliateHotproductQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.affiliate.hotproduct.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAffiliateHotproductQueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressAffiliateHotproductQueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppSignature is AppSignature Setter
// 请求签名
func (r *AliexpressAffiliateHotproductQueryAPIRequest) SetAppSignature(_appSignature string) error {
	r._appSignature = _appSignature
	r.Set("app_signature", _appSignature)
	return nil
}

// GetAppSignature AppSignature Getter
func (r AliexpressAffiliateHotproductQueryAPIRequest) GetAppSignature() string {
	return r._appSignature
}

// SetCategoryIds is CategoryIds Setter
// 类目ID列表
func (r *AliexpressAffiliateHotproductQueryAPIRequest) SetCategoryIds(_categoryIds string) error {
	r._categoryIds = _categoryIds
	r.Set("category_ids", _categoryIds)
	return nil
}

// GetCategoryIds CategoryIds Getter
func (r AliexpressAffiliateHotproductQueryAPIRequest) GetCategoryIds() string {
	return r._categoryIds
}

// SetFields is Fields Setter
// 返回字段列表
func (r *AliexpressAffiliateHotproductQueryAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r AliexpressAffiliateHotproductQueryAPIRequest) GetFields() string {
	return r._fields
}

// SetKeywords is Keywords Setter
// 关键词
func (r *AliexpressAffiliateHotproductQueryAPIRequest) SetKeywords(_keywords string) error {
	r._keywords = _keywords
	r.Set("keywords", _keywords)
	return nil
}

// GetKeywords Keywords Getter
func (r AliexpressAffiliateHotproductQueryAPIRequest) GetKeywords() string {
	return r._keywords
}

// SetPlatformProductType is PlatformProductType Setter
// 平台商家类型：ALL,PLAZA,TMALL
func (r *AliexpressAffiliateHotproductQueryAPIRequest) SetPlatformProductType(_platformProductType string) error {
	r._platformProductType = _platformProductType
	r.Set("platform_product_type", _platformProductType)
	return nil
}

// GetPlatformProductType PlatformProductType Getter
func (r AliexpressAffiliateHotproductQueryAPIRequest) GetPlatformProductType() string {
	return r._platformProductType
}

// SetSort is Sort Setter
// 排序方式:SALE_PRICE_ASC, SALE_PRICE_DESC, DISCOUNT_ASC, DISCOUNT_DESC, LAST_VOLUME_ASC, LAST_VOLUME_DESC
func (r *AliexpressAffiliateHotproductQueryAPIRequest) SetSort(_sort string) error {
	r._sort = _sort
	r.Set("sort", _sort)
	return nil
}

// GetSort Sort Getter
func (r AliexpressAffiliateHotproductQueryAPIRequest) GetSort() string {
	return r._sort
}

// SetTargetCurrency is TargetCurrency Setter
// 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
func (r *AliexpressAffiliateHotproductQueryAPIRequest) SetTargetCurrency(_targetCurrency string) error {
	r._targetCurrency = _targetCurrency
	r.Set("target_currency", _targetCurrency)
	return nil
}

// GetTargetCurrency TargetCurrency Getter
func (r AliexpressAffiliateHotproductQueryAPIRequest) GetTargetCurrency() string {
	return r._targetCurrency
}

// SetTargetLanguage is TargetLanguage Setter
// 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
func (r *AliexpressAffiliateHotproductQueryAPIRequest) SetTargetLanguage(_targetLanguage string) error {
	r._targetLanguage = _targetLanguage
	r.Set("target_language", _targetLanguage)
	return nil
}

// GetTargetLanguage TargetLanguage Getter
func (r AliexpressAffiliateHotproductQueryAPIRequest) GetTargetLanguage() string {
	return r._targetLanguage
}

// SetTrackingId is TrackingId Setter
// trackingId
func (r *AliexpressAffiliateHotproductQueryAPIRequest) SetTrackingId(_trackingId string) error {
	r._trackingId = _trackingId
	r.Set("tracking_id", _trackingId)
	return nil
}

// GetTrackingId TrackingId Getter
func (r AliexpressAffiliateHotproductQueryAPIRequest) GetTrackingId() string {
	return r._trackingId
}

// SetDeliveryDays is DeliveryDays Setter
// 物流到达时间。3：3日达，5：5 日达，7：7日达，10：10日达
func (r *AliexpressAffiliateHotproductQueryAPIRequest) SetDeliveryDays(_deliveryDays string) error {
	r._deliveryDays = _deliveryDays
	r.Set("delivery_days", _deliveryDays)
	return nil
}

// GetDeliveryDays DeliveryDays Getter
func (r AliexpressAffiliateHotproductQueryAPIRequest) GetDeliveryDays() string {
	return r._deliveryDays
}

// SetShipToCountry is ShipToCountry Setter
// 商品收货国家，根据该国家税率政策返回对应商品价格
func (r *AliexpressAffiliateHotproductQueryAPIRequest) SetShipToCountry(_shipToCountry string) error {
	r._shipToCountry = _shipToCountry
	r.Set("ship_to_country", _shipToCountry)
	return nil
}

// GetShipToCountry ShipToCountry Getter
func (r AliexpressAffiliateHotproductQueryAPIRequest) GetShipToCountry() string {
	return r._shipToCountry
}

// SetMaxSalePrice is MaxSalePrice Setter
// 最大售价
func (r *AliexpressAffiliateHotproductQueryAPIRequest) SetMaxSalePrice(_maxSalePrice int64) error {
	r._maxSalePrice = _maxSalePrice
	r.Set("max_sale_price", _maxSalePrice)
	return nil
}

// GetMaxSalePrice MaxSalePrice Getter
func (r AliexpressAffiliateHotproductQueryAPIRequest) GetMaxSalePrice() int64 {
	return r._maxSalePrice
}

// SetMinSalePrice is MinSalePrice Setter
// 最小售价
func (r *AliexpressAffiliateHotproductQueryAPIRequest) SetMinSalePrice(_minSalePrice int64) error {
	r._minSalePrice = _minSalePrice
	r.Set("min_sale_price", _minSalePrice)
	return nil
}

// GetMinSalePrice MinSalePrice Getter
func (r AliexpressAffiliateHotproductQueryAPIRequest) GetMinSalePrice() int64 {
	return r._minSalePrice
}

// SetPageNo is PageNo Setter
// 请求页数
func (r *AliexpressAffiliateHotproductQueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r AliexpressAffiliateHotproductQueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每次请求数量
func (r *AliexpressAffiliateHotproductQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AliexpressAffiliateHotproductQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}

var poolAliexpressAffiliateHotproductQueryAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressAffiliateHotproductQueryRequest()
	},
}

// GetAliexpressAffiliateHotproductQueryRequest 从 sync.Pool 获取 AliexpressAffiliateHotproductQueryAPIRequest
func GetAliexpressAffiliateHotproductQueryAPIRequest() *AliexpressAffiliateHotproductQueryAPIRequest {
	return poolAliexpressAffiliateHotproductQueryAPIRequest.Get().(*AliexpressAffiliateHotproductQueryAPIRequest)
}

// ReleaseAliexpressAffiliateHotproductQueryAPIRequest 将 AliexpressAffiliateHotproductQueryAPIRequest 放入 sync.Pool
func ReleaseAliexpressAffiliateHotproductQueryAPIRequest(v *AliexpressAffiliateHotproductQueryAPIRequest) {
	v.Reset()
	poolAliexpressAffiliateHotproductQueryAPIRequest.Put(v)
}
