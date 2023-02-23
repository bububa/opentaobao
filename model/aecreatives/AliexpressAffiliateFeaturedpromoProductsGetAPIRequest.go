package aecreatives

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAffiliateFeaturedpromoProductsGetAPIRequest 联盟主题推广活动商品信息获取 API请求
// aliexpress.affiliate.featuredpromo.products.get
//
// 根据联盟主题推广活动或主题品库查询对应的商品。如下品库为固定品库，可长期调用。品库类型和名称如下：高佣品（Hot Product）、新品（New Arrival）、热销商品（Best Seller）、每周尖货（weeklydeals）
type AliexpressAffiliateFeaturedpromoProductsGetAPIRequest struct {
	model.Params
	// 请求签名
	_appSignature string
	// 类目 ID 如何获取category_id，请参考，https://open.taobao.com/api.htm?docId=45801&docType=2&scopeId=17063
	_categoryId string
	// 返回字段列表
	_fields string
	// 活动结束时间，PST 时区
	_promotionEndTime string
	// 主题活动的名称，如何获取主题活动，请参考"get featuredpromo info" API. 固定主题：高佣品（Hot Product）、新品（New Arrival）、热销商品（Best Seller）、每周尖货（weeklydeals）
	_promotionName string
	// 活动开始时间，PST 时区
	_promotionStartTime string
	// 排序方式：commissionAsc，commissionDesc, priceAsc，priceDesc，volumeAsc、volumeDesc, discountAsc, discountDesc, ratingAsc，ratingDesc, promotionTimeAsc, pr
	_sort string
	// 目标币种，可根据目标币种返回对应币种：USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
	_targetCurrency string
	// 目标语言，可根据目标语言返回对应语言：EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IN
	_targetLanguage string
	// trackingID
	_trackingId string
	// 商品收货国家，可筛选能销售至该国家的商品，并根据该国家税率政策返回对应商品价格
	_country string
	// 查询页码
	_pageNo int64
	// 每页记录数，1-50
	_pageSize int64
}

// NewAliexpressAffiliateFeaturedpromoProductsGetRequest 初始化AliexpressAffiliateFeaturedpromoProductsGetAPIRequest对象
func NewAliexpressAffiliateFeaturedpromoProductsGetRequest() *AliexpressAffiliateFeaturedpromoProductsGetAPIRequest {
	return &AliexpressAffiliateFeaturedpromoProductsGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) GetApiMethodName() string {
	return "aliexpress.affiliate.featuredpromo.products.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppSignature is AppSignature Setter
// 请求签名
func (r *AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) SetAppSignature(_appSignature string) error {
	r._appSignature = _appSignature
	r.Set("app_signature", _appSignature)
	return nil
}

// GetAppSignature AppSignature Getter
func (r AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) GetAppSignature() string {
	return r._appSignature
}

// SetCategoryId is CategoryId Setter
// 类目 ID 如何获取category_id，请参考，https://open.taobao.com/api.htm?docId=45801&amp;docType=2&amp;scopeId=17063
func (r *AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) SetCategoryId(_categoryId string) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) GetCategoryId() string {
	return r._categoryId
}

// SetFields is Fields Setter
// 返回字段列表
func (r *AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) GetFields() string {
	return r._fields
}

// SetPromotionEndTime is PromotionEndTime Setter
// 活动结束时间，PST 时区
func (r *AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) SetPromotionEndTime(_promotionEndTime string) error {
	r._promotionEndTime = _promotionEndTime
	r.Set("promotion_end_time", _promotionEndTime)
	return nil
}

// GetPromotionEndTime PromotionEndTime Getter
func (r AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) GetPromotionEndTime() string {
	return r._promotionEndTime
}

// SetPromotionName is PromotionName Setter
// 主题活动的名称，如何获取主题活动，请参考&#34;get featuredpromo info&#34; API. 固定主题：高佣品（Hot Product）、新品（New Arrival）、热销商品（Best Seller）、每周尖货（weeklydeals）
func (r *AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) SetPromotionName(_promotionName string) error {
	r._promotionName = _promotionName
	r.Set("promotion_name", _promotionName)
	return nil
}

// GetPromotionName PromotionName Getter
func (r AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) GetPromotionName() string {
	return r._promotionName
}

// SetPromotionStartTime is PromotionStartTime Setter
// 活动开始时间，PST 时区
func (r *AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) SetPromotionStartTime(_promotionStartTime string) error {
	r._promotionStartTime = _promotionStartTime
	r.Set("promotion_start_time", _promotionStartTime)
	return nil
}

// GetPromotionStartTime PromotionStartTime Getter
func (r AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) GetPromotionStartTime() string {
	return r._promotionStartTime
}

// SetSort is Sort Setter
// 排序方式：commissionAsc，commissionDesc, priceAsc，priceDesc，volumeAsc、volumeDesc, discountAsc, discountDesc, ratingAsc，ratingDesc, promotionTimeAsc, pr
func (r *AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) SetSort(_sort string) error {
	r._sort = _sort
	r.Set("sort", _sort)
	return nil
}

// GetSort Sort Getter
func (r AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) GetSort() string {
	return r._sort
}

// SetTargetCurrency is TargetCurrency Setter
// 目标币种，可根据目标币种返回对应币种：USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
func (r *AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) SetTargetCurrency(_targetCurrency string) error {
	r._targetCurrency = _targetCurrency
	r.Set("target_currency", _targetCurrency)
	return nil
}

// GetTargetCurrency TargetCurrency Getter
func (r AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) GetTargetCurrency() string {
	return r._targetCurrency
}

// SetTargetLanguage is TargetLanguage Setter
// 目标语言，可根据目标语言返回对应语言：EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IN
func (r *AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) SetTargetLanguage(_targetLanguage string) error {
	r._targetLanguage = _targetLanguage
	r.Set("target_language", _targetLanguage)
	return nil
}

// GetTargetLanguage TargetLanguage Getter
func (r AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) GetTargetLanguage() string {
	return r._targetLanguage
}

// SetTrackingId is TrackingId Setter
// trackingID
func (r *AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) SetTrackingId(_trackingId string) error {
	r._trackingId = _trackingId
	r.Set("tracking_id", _trackingId)
	return nil
}

// GetTrackingId TrackingId Getter
func (r AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) GetTrackingId() string {
	return r._trackingId
}

// SetCountry is Country Setter
// 商品收货国家，可筛选能销售至该国家的商品，并根据该国家税率政策返回对应商品价格
func (r *AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) SetCountry(_country string) error {
	r._country = _country
	r.Set("country", _country)
	return nil
}

// GetCountry Country Getter
func (r AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) GetCountry() string {
	return r._country
}

// SetPageNo is PageNo Setter
// 查询页码
func (r *AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) SetPageNo(_pageNo int64) error {
	r._pageNo = _pageNo
	r.Set("page_no", _pageNo)
	return nil
}

// GetPageNo PageNo Getter
func (r AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) GetPageNo() int64 {
	return r._pageNo
}

// SetPageSize is PageSize Setter
// 每页记录数，1-50
func (r *AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) SetPageSize(_pageSize int64) error {
	r._pageSize = _pageSize
	r.Set("page_size", _pageSize)
	return nil
}

// GetPageSize PageSize Getter
func (r AliexpressAffiliateFeaturedpromoProductsGetAPIRequest) GetPageSize() int64 {
	return r._pageSize
}
