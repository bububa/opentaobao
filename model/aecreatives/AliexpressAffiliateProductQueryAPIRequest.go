package aecreatives

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAffiliateProductQueryAPIRequest 联盟推广商品获取接口 API请求
// aliexpress.affiliate.product.query
//
// 联盟推广商品搜索接口，用于搜索联盟推广商品数据
type AliexpressAffiliateProductQueryAPIRequest struct {
	model.Params
	// 安全签名
	_appSignature string
	// 类目ID列表
	_categoryIds string
	// 返回字段列表
	_fields string
	// 关键词
	_keywords string
	// 平台商品类型：ALL,PLAZA,TMALL
	_platformProductType string
	// 排序方式:SALE_PRICE_ASC, SALE_PRICE_DESC, LAST_VOLUME_ASC, LAST_VOLUME_DESC
	_sort string
	// 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
	_targetCurrency string
	// 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
	_targetLanguage string
	// trackingId
	_trackingId string
	// 商品收货国家，根据该国家税率政策返回对应商品价格
	_shipToCountry string
	// 物流到达时间。3：3日达，5：5 日达，7：7日达，10：10日达
	_deliveryDays string
	// 最大售价
	_maxSalePrice int64
	// 最小售价
	_minSalePrice int64
	// 查询页码
	_pageNo int64
	// 每页记录数
	_pageSize int64
}

// NewAliexpressAffiliateProductQueryRequest 初始化AliexpressAffiliateProductQueryAPIRequest对象
func NewAliexpressAffiliateProductQueryRequest() *AliexpressAffiliateProductQueryAPIRequest {
	return &AliexpressAffiliateProductQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAffiliateProductQueryAPIRequest) GetApiMethodName() string {
	return "aliexpress.affiliate.product.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAffiliateProductQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetAppSignature is AppSignature Setter
// 安全签名
func (r *AliexpressAffiliateProductQueryAPIRequest) SetAppSignature(_appSignature string) error {
	r._appSignature = _appSignature
	r.Set("app_signature", _appSignature)
	return nil
}

// GetAppSignature AppSignature Getter
func (r AliexpressAffiliateProductQueryAPIRequest) GetAppSignature() string {
	return r._appSignature
}

// SetCategoryIds is CategoryIds Setter
// 类目ID列表
func (r *AliexpressAffiliateProductQueryAPIRequest) SetCategoryIds(_categoryIds string) error {
	r._categoryIds = _categoryIds
	r.Set("category_ids", _categoryIds)
	return nil
}

// GetCategoryIds CategoryIds Getter
func (r AliexpressAffiliateProductQueryAPIRequest) GetCategoryIds() string {
	return r._categoryIds
}

// SetFields is Fields Setter
// 返回字段列表
func (r *AliexpressAffiliateProductQueryAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r AliexpressAffiliateProductQueryAPIRequest) GetFields() string {
	return r._fields
}

// SetKeywords is Keywords Setter
// 关键词
func (r *AliexpressAffiliateProductQueryAPIRequest) SetKeywords(_keywords string) error {
	r._keywords = _keywords
	r.Set("keywords", _keywords)
	return nil
}

// GetKeywords Keywords Getter
func (r AliexpressAffiliateProductQueryAPIRequest) GetKeywords() string {
	return r._keywords
}

// SetPlatformProductType is PlatformProductType Setter
// 平台商品类型：ALL,PLAZA,TMALL
func (r *AliexpressAffiliateProductQueryAPIRequest) SetPlatformProductType(_platformProductType string) error {
	r._platformProductType = _platformProductType
	r.Set("platform_product_type", _platformProductType)
	return nil
}

// GetPlatformProductType PlatformProductType Getter
func (r AliexpressAffiliateProductQueryAPIRequest) GetPlatformProductType() string {
	return r._platformProductType
}

// SetSort is Sort Setter
// 排序方式:SALE_PRICE_ASC, SALE_PRICE_DESC, LAST_VOLUME_ASC, LAST_VOLUME_DESC
func (r *AliexpressAffiliateProductQueryAPIRequest) SetSort(_sort string) error {
	r._sort = _sort
	r.Set("sort", _sort)
	return nil
}

// GetSort Sort Getter
func (r AliexpressAffiliateProductQueryAPIRequest) GetSort() string {
	return r._sort
}

// SetTargetCurrency is TargetCurrency Setter
// 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
func (r *AliexpressAffiliateProductQueryAPIRequest) SetTargetCurrency(_targetCurrency string) error {
	r._targetCurrency = _targetCurrency
	r.Set("target_currency", _targetCurrency)
	return nil
}

// GetTargetCurrency TargetCurrency Getter
func (r AliexpressAffiliateProductQueryAPIRequest) GetTargetCurrency() string {
	return r._targetCurrency
}

// SetTargetLanguage is TargetLanguage Setter
// 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
func (r *AliexpressAffiliateProductQueryAPIRequest) SetTargetLanguage(_targetLanguage string) error {
	r._targetLanguage = _targetLanguage
	r.Set("target_language", _targetLanguage)
	return nil
}

// GetTargetLanguage TargetLanguage Getter
func (r AliexpressAffiliateProductQueryAPIRequest) GetTargetLanguage() string {
	return r._targetLanguage
}

// SetTrackingId is TrackingId Setter
// trackingId
func (r *AliexpressAffiliateProductQueryAPIRequest) SetTrackingId(_trackingId string) error {
	r._trackingId = _trackingId
	r.Set("tracking_id", _trackingId)
	return nil
}

// GetTrackingId TrackingId Getter
func (r AliexpressAffiliateProductQueryAPIRequest) GetTrackingId() string {
	return r._trackingId
}

// SetShipToCountry is ShipToCountry Setter
// 商品收货国家，根据该国家税率政策返回对应商品价格
func (r *AliexpressAffiliateProductQueryAPIRequest) SetShipToCountry(_shipToCountry string) error {
	r._shipToCountry = _shipToCountry
	r.Set("ship_to_country", _shipToCountry)
	return nil
}

// GetShipToCountry ShipToCountry Getter
func (r AliexpressAffiliateProductQueryAPIRequest) GetShipToCountry() string {
	return r._shipToCountry
}

// SetDeliveryDays is DeliveryDays Setter
// 物流到达时间。3：3日达，5：5 日达，7：7日达，10：10日达
func (r *AliexpressAffiliateProductQueryAPIRequest) SetDeliveryDays(_deliveryDays string) error {
	r._deliveryDays = _deliveryDays
	r.Set("delivery_days", _deliveryDays)
	return nil
}

// GetDeliveryDays DeliveryDays Getter
func (r AliexpressAffiliateProductQueryAPIRequest) GetDeliveryDays() string {
	return r._deliveryDays
}

// SetMaxSalePrice is MaxSalePrice Setter
// 最大售价
func (r *AliexpressAffiliateProductQueryAPIRequest) SetMaxSalePrice(_maxSalePrice int64) error {
	r._maxSalePrice = _maxSalePrice
	r.Set("max_sale_price", _maxSalePrice)
	return nil
}

// GetMaxSalePrice MaxSalePrice Getter
func (r AliexpressAffiliateProductQueryAPIRequest) GetMaxSalePrice() int64 {
	return r._maxSalePrice
}

// SetMinSalePrice is MinSalePrice Setter
// 最小售价
func (r *AliexpressAffiliateProductQueryAPIRequest) SetMinSalePrice(_minSalePrice int64) error {
	r._minSalePrice = _minSalePrice
	r.Set("min_sale_price", _minSalePrice)
	return nil
}

// GetMinSalePrice MinSalePrice Getter
func (r AliexpressAffiliateProductQueryAPIRequest) GetMinSalePrice() int64 {
	return r._minSalePrice
}

// SetPageNo is PageNo Setter
// 查询页码
func (r *AliexpressAffiliateProductQueryAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r AliexpressAffiliateProductQueryAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页记录数
func (r *AliexpressAffiliateProductQueryAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AliexpressAffiliateProductQueryAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
