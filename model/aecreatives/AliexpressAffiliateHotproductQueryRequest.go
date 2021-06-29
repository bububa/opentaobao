package aecreatives

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
查询联盟爆品数据 APIRequest
aliexpress.affiliate.hotproduct.query

查询联盟爆品API
*/
type AliexpressAffiliateHotproductQueryRequest struct {
    model.Params

    // 请求签名
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

    // 请求页数
    pageNo   int64 

    // 每次请求数量
    pageSize   int64 

    // 平台商家类型：ALL,PLAZA,TMALL
    platformProductType   string 

    // 排序方式:SALE_PRICE_ASC, SALE_PRICE_DESC, DISCOUNT_ASC, DISCOUNT_DESC, LAST_VOLUME_ASC, LAST_VOLUME_DESC
    sort   string 

    // 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
    targetCurrency   string 

    // 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
    targetLanguage   string 

    // trackingId
    trackingId   string 

    // 物流到达时间。3：3日达，5：5 日达，7：7日达，10：10日达
    deliveryDays   string 

    // 商品收货国家，根据该国家税率政策返回对应商品价格
    shipToCountry   string 

}

func NewAliexpressAffiliateHotproductQueryRequest() *AliexpressAffiliateHotproductQueryRequest{
    return &AliexpressAffiliateHotproductQueryRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressAffiliateHotproductQueryRequest) GetApiMethodName() string {
    return "aliexpress.affiliate.hotproduct.query"
}

func (r AliexpressAffiliateHotproductQueryRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressAffiliateHotproductQueryRequest) SetAppSignature(appSignature string) error {
    r.appSignature = appSignature
    r.Set("app_signature", appSignature)
    return nil
}

func (r AliexpressAffiliateHotproductQueryRequest) GetAppSignature() string {
    return r.appSignature
}

func (r *AliexpressAffiliateHotproductQueryRequest) SetCategoryIds(categoryIds string) error {
    r.categoryIds = categoryIds
    r.Set("category_ids", categoryIds)
    return nil
}

func (r AliexpressAffiliateHotproductQueryRequest) GetCategoryIds() string {
    return r.categoryIds
}

func (r *AliexpressAffiliateHotproductQueryRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r AliexpressAffiliateHotproductQueryRequest) GetFields() string {
    return r.fields
}

func (r *AliexpressAffiliateHotproductQueryRequest) SetKeywords(keywords string) error {
    r.keywords = keywords
    r.Set("keywords", keywords)
    return nil
}

func (r AliexpressAffiliateHotproductQueryRequest) GetKeywords() string {
    return r.keywords
}

func (r *AliexpressAffiliateHotproductQueryRequest) SetMaxSalePrice(maxSalePrice int64) error {
    r.maxSalePrice = maxSalePrice
    r.Set("max_sale_price", maxSalePrice)
    return nil
}

func (r AliexpressAffiliateHotproductQueryRequest) GetMaxSalePrice() int64 {
    return r.maxSalePrice
}

func (r *AliexpressAffiliateHotproductQueryRequest) SetMinSalePrice(minSalePrice int64) error {
    r.minSalePrice = minSalePrice
    r.Set("min_sale_price", minSalePrice)
    return nil
}

func (r AliexpressAffiliateHotproductQueryRequest) GetMinSalePrice() int64 {
    return r.minSalePrice
}

func (r *AliexpressAffiliateHotproductQueryRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r AliexpressAffiliateHotproductQueryRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *AliexpressAffiliateHotproductQueryRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AliexpressAffiliateHotproductQueryRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AliexpressAffiliateHotproductQueryRequest) SetPlatformProductType(platformProductType string) error {
    r.platformProductType = platformProductType
    r.Set("platform_product_type", platformProductType)
    return nil
}

func (r AliexpressAffiliateHotproductQueryRequest) GetPlatformProductType() string {
    return r.platformProductType
}

func (r *AliexpressAffiliateHotproductQueryRequest) SetSort(sort string) error {
    r.sort = sort
    r.Set("sort", sort)
    return nil
}

func (r AliexpressAffiliateHotproductQueryRequest) GetSort() string {
    return r.sort
}

func (r *AliexpressAffiliateHotproductQueryRequest) SetTargetCurrency(targetCurrency string) error {
    r.targetCurrency = targetCurrency
    r.Set("target_currency", targetCurrency)
    return nil
}

func (r AliexpressAffiliateHotproductQueryRequest) GetTargetCurrency() string {
    return r.targetCurrency
}

func (r *AliexpressAffiliateHotproductQueryRequest) SetTargetLanguage(targetLanguage string) error {
    r.targetLanguage = targetLanguage
    r.Set("target_language", targetLanguage)
    return nil
}

func (r AliexpressAffiliateHotproductQueryRequest) GetTargetLanguage() string {
    return r.targetLanguage
}

func (r *AliexpressAffiliateHotproductQueryRequest) SetTrackingId(trackingId string) error {
    r.trackingId = trackingId
    r.Set("tracking_id", trackingId)
    return nil
}

func (r AliexpressAffiliateHotproductQueryRequest) GetTrackingId() string {
    return r.trackingId
}

func (r *AliexpressAffiliateHotproductQueryRequest) SetDeliveryDays(deliveryDays string) error {
    r.deliveryDays = deliveryDays
    r.Set("delivery_days", deliveryDays)
    return nil
}

func (r AliexpressAffiliateHotproductQueryRequest) GetDeliveryDays() string {
    return r.deliveryDays
}

func (r *AliexpressAffiliateHotproductQueryRequest) SetShipToCountry(shipToCountry string) error {
    r.shipToCountry = shipToCountry
    r.Set("ship_to_country", shipToCountry)
    return nil
}

func (r AliexpressAffiliateHotproductQueryRequest) GetShipToCountry() string {
    return r.shipToCountry
}

