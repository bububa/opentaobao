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
    appSignature   string
    // 类目ID列表
    categoryIds   string
    // 返回字段列表
    fields   string
    // 关键词
    keywords   string
    // 最大售价
    maxSalePrice   int64
    // 最小售价
    minSalePrice   int64
    // 查询页码
    pageNo   int64
    // 每页记录数
    pageSize   int64
    // 平台商品类型：ALL,PLAZA,TMALL
    platformProductType   string
    // 排序方式:SALE_PRICE_ASC, SALE_PRICE_DESC, LAST_VOLUME_ASC, LAST_VOLUME_DESC
    sort   string
    // 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
    targetCurrency   string
    // 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
    targetLanguage   string
    // trackingId
    trackingId   string
    // 商品收货国家，根据该国家税率政策返回对应商品价格
    shipToCountry   string
    // 物流到达时间。3：3日达，5：5 日达，7：7日达，10：10日达
    deliveryDays   string
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
func (r *AliexpressAffiliateProductQueryRequest) SetAppSignature(appSignature string) error {
    r.appSignature = appSignature
    r.Set("app_signature", appSignature)
    return nil
}

// AppSignature Getter
func (r AliexpressAffiliateProductQueryRequest) GetAppSignature() string {
    return r.appSignature
}
// CategoryIds Setter
// 类目ID列表
func (r *AliexpressAffiliateProductQueryRequest) SetCategoryIds(categoryIds string) error {
    r.categoryIds = categoryIds
    r.Set("category_ids", categoryIds)
    return nil
}

// CategoryIds Getter
func (r AliexpressAffiliateProductQueryRequest) GetCategoryIds() string {
    return r.categoryIds
}
// Fields Setter
// 返回字段列表
func (r *AliexpressAffiliateProductQueryRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r AliexpressAffiliateProductQueryRequest) GetFields() string {
    return r.fields
}
// Keywords Setter
// 关键词
func (r *AliexpressAffiliateProductQueryRequest) SetKeywords(keywords string) error {
    r.keywords = keywords
    r.Set("keywords", keywords)
    return nil
}

// Keywords Getter
func (r AliexpressAffiliateProductQueryRequest) GetKeywords() string {
    return r.keywords
}
// MaxSalePrice Setter
// 最大售价
func (r *AliexpressAffiliateProductQueryRequest) SetMaxSalePrice(maxSalePrice int64) error {
    r.maxSalePrice = maxSalePrice
    r.Set("max_sale_price", maxSalePrice)
    return nil
}

// MaxSalePrice Getter
func (r AliexpressAffiliateProductQueryRequest) GetMaxSalePrice() int64 {
    return r.maxSalePrice
}
// MinSalePrice Setter
// 最小售价
func (r *AliexpressAffiliateProductQueryRequest) SetMinSalePrice(minSalePrice int64) error {
    r.minSalePrice = minSalePrice
    r.Set("min_sale_price", minSalePrice)
    return nil
}

// MinSalePrice Getter
func (r AliexpressAffiliateProductQueryRequest) GetMinSalePrice() int64 {
    return r.minSalePrice
}
// PageNo Setter
// 查询页码
func (r *AliexpressAffiliateProductQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r AliexpressAffiliateProductQueryRequest) GetPageNo() int64 {
    return r.pageNo
}
// PageSize Setter
// 每页记录数
func (r *AliexpressAffiliateProductQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

// PageSize Getter
func (r AliexpressAffiliateProductQueryRequest) GetPageSize() int64 {
    return r.pageSize
}
// PlatformProductType Setter
// 平台商品类型：ALL,PLAZA,TMALL
func (r *AliexpressAffiliateProductQueryRequest) SetPlatformProductType(platformProductType string) error {
    r.platformProductType = platformProductType
    r.Set("platform_product_type", platformProductType)
    return nil
}

// PlatformProductType Getter
func (r AliexpressAffiliateProductQueryRequest) GetPlatformProductType() string {
    return r.platformProductType
}
// Sort Setter
// 排序方式:SALE_PRICE_ASC, SALE_PRICE_DESC, LAST_VOLUME_ASC, LAST_VOLUME_DESC
func (r *AliexpressAffiliateProductQueryRequest) SetSort(sort string) error {
    r.sort = sort
    r.Set("sort", sort)
    return nil
}

// Sort Getter
func (r AliexpressAffiliateProductQueryRequest) GetSort() string {
    return r.sort
}
// TargetCurrency Setter
// 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
func (r *AliexpressAffiliateProductQueryRequest) SetTargetCurrency(targetCurrency string) error {
    r.targetCurrency = targetCurrency
    r.Set("target_currency", targetCurrency)
    return nil
}

// TargetCurrency Getter
func (r AliexpressAffiliateProductQueryRequest) GetTargetCurrency() string {
    return r.targetCurrency
}
// TargetLanguage Setter
// 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
func (r *AliexpressAffiliateProductQueryRequest) SetTargetLanguage(targetLanguage string) error {
    r.targetLanguage = targetLanguage
    r.Set("target_language", targetLanguage)
    return nil
}

// TargetLanguage Getter
func (r AliexpressAffiliateProductQueryRequest) GetTargetLanguage() string {
    return r.targetLanguage
}
// TrackingId Setter
// trackingId
func (r *AliexpressAffiliateProductQueryRequest) SetTrackingId(trackingId string) error {
    r.trackingId = trackingId
    r.Set("tracking_id", trackingId)
    return nil
}

// TrackingId Getter
func (r AliexpressAffiliateProductQueryRequest) GetTrackingId() string {
    return r.trackingId
}
// ShipToCountry Setter
// 商品收货国家，根据该国家税率政策返回对应商品价格
func (r *AliexpressAffiliateProductQueryRequest) SetShipToCountry(shipToCountry string) error {
    r.shipToCountry = shipToCountry
    r.Set("ship_to_country", shipToCountry)
    return nil
}

// ShipToCountry Getter
func (r AliexpressAffiliateProductQueryRequest) GetShipToCountry() string {
    return r.shipToCountry
}
// DeliveryDays Setter
// 物流到达时间。3：3日达，5：5 日达，7：7日达，10：10日达
func (r *AliexpressAffiliateProductQueryRequest) SetDeliveryDays(deliveryDays string) error {
    r.deliveryDays = deliveryDays
    r.Set("delivery_days", deliveryDays)
    return nil
}

// DeliveryDays Getter
func (r AliexpressAffiliateProductQueryRequest) GetDeliveryDays() string {
    return r.deliveryDays
}
