package aecreatives

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
联盟推广商品获取接口 API请求
aliexpress.affiliate.product.query

联盟推广商品搜索接口，用于搜索联盟推广商品数据
*/
type AliexpressAffiliateProductQueryRequest struct {
    model.Params
    // 安全签名
    _appSignature   string
    // 类目ID列表
    _categoryIds   string
    // 返回字段列表
    _fields   string
    // 关键词
    _keywords   string
    // 最大售价
    _maxSalePrice   int64
    // 最小售价
    _minSalePrice   int64
    // 查询页码
    _pageNo   int64
    // 每页记录数
    _pageSize   int64
    // 平台商品类型：ALL,PLAZA,TMALL
    _platformProductType   string
    // 排序方式:SALE_PRICE_ASC, SALE_PRICE_DESC, LAST_VOLUME_ASC, LAST_VOLUME_DESC
    _sort   string
    // 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
    _targetCurrency   string
    // 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
    _targetLanguage   string
    // trackingId
    _trackingId   string
    // 商品收货国家，根据该国家税率政策返回对应商品价格
    _shipToCountry   string
    // 物流到达时间。3：3日达，5：5 日达，7：7日达，10：10日达
    _deliveryDays   string
}

// 初始化AliexpressAffiliateProductQueryRequest对象
func NewAliexpressAffiliateProductQueryRequest() *AliexpressAffiliateProductQueryRequest{
    return &AliexpressAffiliateProductQueryRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressAffiliateProductQueryRequest) GetApiMethodName() string {
    return "aliexpress.affiliate.product.query"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressAffiliateProductQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppSignature Setter
// 安全签名
func (r *AliexpressAffiliateProductQueryRequest) SetAppSignature(_appSignature string) error {
    r._appSignature = _appSignature
    r.Set("app_signature", _appSignature)
    return nil
}

// AppSignature Getter
func (r AliexpressAffiliateProductQueryRequest) GetAppSignature() string {
    return r._appSignature
}
// CategoryIds Setter
// 类目ID列表
func (r *AliexpressAffiliateProductQueryRequest) SetCategoryIds(_categoryIds string) error {
    r._categoryIds = _categoryIds
    r.Set("category_ids", _categoryIds)
    return nil
}

// CategoryIds Getter
func (r AliexpressAffiliateProductQueryRequest) GetCategoryIds() string {
    return r._categoryIds
}
// Fields Setter
// 返回字段列表
func (r *AliexpressAffiliateProductQueryRequest) SetFields(_fields string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r AliexpressAffiliateProductQueryRequest) GetFields() string {
    return r._fields
}
// Keywords Setter
// 关键词
func (r *AliexpressAffiliateProductQueryRequest) SetKeywords(_keywords string) error {
    r._keywords = _keywords
    r.Set("keywords", _keywords)
    return nil
}

// Keywords Getter
func (r AliexpressAffiliateProductQueryRequest) GetKeywords() string {
    return r._keywords
}
// MaxSalePrice Setter
// 最大售价
func (r *AliexpressAffiliateProductQueryRequest) SetMaxSalePrice(_maxSalePrice int64) error {
    r._maxSalePrice = _maxSalePrice
    r.Set("max_sale_price", _maxSalePrice)
    return nil
}

// MaxSalePrice Getter
func (r AliexpressAffiliateProductQueryRequest) GetMaxSalePrice() int64 {
    return r._maxSalePrice
}
// MinSalePrice Setter
// 最小售价
func (r *AliexpressAffiliateProductQueryRequest) SetMinSalePrice(_minSalePrice int64) error {
    r._minSalePrice = _minSalePrice
    r.Set("min_sale_price", _minSalePrice)
    return nil
}

// MinSalePrice Getter
func (r AliexpressAffiliateProductQueryRequest) GetMinSalePrice() int64 {
    return r._minSalePrice
}
// PageNo Setter
// 查询页码
func (r *AliexpressAffiliateProductQueryRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r AliexpressAffiliateProductQueryRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页记录数
func (r *AliexpressAffiliateProductQueryRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AliexpressAffiliateProductQueryRequest) GetPageSize() int64 {
    return r._pageSize
}
// PlatformProductType Setter
// 平台商品类型：ALL,PLAZA,TMALL
func (r *AliexpressAffiliateProductQueryRequest) SetPlatformProductType(_platformProductType string) error {
    r._platformProductType = _platformProductType
    r.Set("platform_product_type", _platformProductType)
    return nil
}

// PlatformProductType Getter
func (r AliexpressAffiliateProductQueryRequest) GetPlatformProductType() string {
    return r._platformProductType
}
// Sort Setter
// 排序方式:SALE_PRICE_ASC, SALE_PRICE_DESC, LAST_VOLUME_ASC, LAST_VOLUME_DESC
func (r *AliexpressAffiliateProductQueryRequest) SetSort(_sort string) error {
    r._sort = _sort
    r.Set("sort", _sort)
    return nil
}

// Sort Getter
func (r AliexpressAffiliateProductQueryRequest) GetSort() string {
    return r._sort
}
// TargetCurrency Setter
// 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
func (r *AliexpressAffiliateProductQueryRequest) SetTargetCurrency(_targetCurrency string) error {
    r._targetCurrency = _targetCurrency
    r.Set("target_currency", _targetCurrency)
    return nil
}

// TargetCurrency Getter
func (r AliexpressAffiliateProductQueryRequest) GetTargetCurrency() string {
    return r._targetCurrency
}
// TargetLanguage Setter
// 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
func (r *AliexpressAffiliateProductQueryRequest) SetTargetLanguage(_targetLanguage string) error {
    r._targetLanguage = _targetLanguage
    r.Set("target_language", _targetLanguage)
    return nil
}

// TargetLanguage Getter
func (r AliexpressAffiliateProductQueryRequest) GetTargetLanguage() string {
    return r._targetLanguage
}
// TrackingId Setter
// trackingId
func (r *AliexpressAffiliateProductQueryRequest) SetTrackingId(_trackingId string) error {
    r._trackingId = _trackingId
    r.Set("tracking_id", _trackingId)
    return nil
}

// TrackingId Getter
func (r AliexpressAffiliateProductQueryRequest) GetTrackingId() string {
    return r._trackingId
}
// ShipToCountry Setter
// 商品收货国家，根据该国家税率政策返回对应商品价格
func (r *AliexpressAffiliateProductQueryRequest) SetShipToCountry(_shipToCountry string) error {
    r._shipToCountry = _shipToCountry
    r.Set("ship_to_country", _shipToCountry)
    return nil
}

// ShipToCountry Getter
func (r AliexpressAffiliateProductQueryRequest) GetShipToCountry() string {
    return r._shipToCountry
}
// DeliveryDays Setter
// 物流到达时间。3：3日达，5：5 日达，7：7日达，10：10日达
func (r *AliexpressAffiliateProductQueryRequest) SetDeliveryDays(_deliveryDays string) error {
    r._deliveryDays = _deliveryDays
    r.Set("delivery_days", _deliveryDays)
    return nil
}

// DeliveryDays Getter
func (r AliexpressAffiliateProductQueryRequest) GetDeliveryDays() string {
    return r._deliveryDays
}
