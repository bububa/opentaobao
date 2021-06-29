package aecreatives

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
联盟主题推广活动商品信息获取 APIRequest
aliexpress.affiliate.featuredpromo.products.get

根据联盟主题推广活动或主题品库查询对应的商品。如下品库为固定品库，可长期调用。品库类型和名称如下：高佣品（Hot Product）、新品（New Arrival）、热销商品（Best Seller）、每周尖货（weeklydeals）
*/
type AliexpressAffiliateFeaturedpromoProductsGetRequest struct {
    model.Params

    // 请求签名
    appSignature   string 

    // 类目 ID 如何获取category_id，请参考，https://open.taobao.com/api.htm?docId=45801&docType=2&scopeId=17063
    categoryId   string 

    // 返回字段列表
    fields   string 

    // 查询页码
    pageNo   int64 

    // 每页记录数，1-50
    pageSize   int64 

    // 活动结束时间，PST 时区
    promotionEndTime   string 

    // 主题活动的名称，如何获取主题活动，请参考"get featuredpromo info" API. 固定主题：高佣品（Hot Product）、新品（New Arrival）、热销商品（Best Seller）、每周尖货（weeklydeals）
    promotionName   string 

    // 活动开始时间，PST 时区
    promotionStartTime   string 

    // 排序方式：commissionAsc，commissionDesc, priceAsc，priceDesc，volumeAsc、volumeDesc, discountAsc, discountDesc, ratingAsc，ratingDesc, promotionTimeAsc, pr
    sort   string 

    // 目标币种，可根据目标币种返回对应币种：USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
    targetCurrency   string 

    // 目标语言，可根据目标语言返回对应语言：EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IN
    targetLanguage   string 

    // trackingID
    trackingId   string 

    // 商品收货国家，可筛选能销售至该国家的商品，并根据该国家税率政策返回对应商品价格
    country   string 

}

func NewAliexpressAffiliateFeaturedpromoProductsGetRequest() *AliexpressAffiliateFeaturedpromoProductsGetRequest{
    return &AliexpressAffiliateFeaturedpromoProductsGetRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetApiMethodName() string {
    return "aliexpress.affiliate.featuredpromo.products.get"
}

func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressAffiliateFeaturedpromoProductsGetRequest) SetAppSignature(appSignature string) error {
    r.appSignature = appSignature
    r.Set("app_signature", appSignature)
    return nil
}

func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetAppSignature() string {
    return r.appSignature
}

func (r *AliexpressAffiliateFeaturedpromoProductsGetRequest) SetCategoryId(categoryId string) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetCategoryId() string {
    return r.categoryId
}

func (r *AliexpressAffiliateFeaturedpromoProductsGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetFields() string {
    return r.fields
}

func (r *AliexpressAffiliateFeaturedpromoProductsGetRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *AliexpressAffiliateFeaturedpromoProductsGetRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AliexpressAffiliateFeaturedpromoProductsGetRequest) SetPromotionEndTime(promotionEndTime string) error {
    r.promotionEndTime = promotionEndTime
    r.Set("promotion_end_time", promotionEndTime)
    return nil
}

func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetPromotionEndTime() string {
    return r.promotionEndTime
}

func (r *AliexpressAffiliateFeaturedpromoProductsGetRequest) SetPromotionName(promotionName string) error {
    r.promotionName = promotionName
    r.Set("promotion_name", promotionName)
    return nil
}

func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetPromotionName() string {
    return r.promotionName
}

func (r *AliexpressAffiliateFeaturedpromoProductsGetRequest) SetPromotionStartTime(promotionStartTime string) error {
    r.promotionStartTime = promotionStartTime
    r.Set("promotion_start_time", promotionStartTime)
    return nil
}

func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetPromotionStartTime() string {
    return r.promotionStartTime
}

func (r *AliexpressAffiliateFeaturedpromoProductsGetRequest) SetSort(sort string) error {
    r.sort = sort
    r.Set("sort", sort)
    return nil
}

func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetSort() string {
    return r.sort
}

func (r *AliexpressAffiliateFeaturedpromoProductsGetRequest) SetTargetCurrency(targetCurrency string) error {
    r.targetCurrency = targetCurrency
    r.Set("target_currency", targetCurrency)
    return nil
}

func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetTargetCurrency() string {
    return r.targetCurrency
}

func (r *AliexpressAffiliateFeaturedpromoProductsGetRequest) SetTargetLanguage(targetLanguage string) error {
    r.targetLanguage = targetLanguage
    r.Set("target_language", targetLanguage)
    return nil
}

func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetTargetLanguage() string {
    return r.targetLanguage
}

func (r *AliexpressAffiliateFeaturedpromoProductsGetRequest) SetTrackingId(trackingId string) error {
    r.trackingId = trackingId
    r.Set("tracking_id", trackingId)
    return nil
}

func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetTrackingId() string {
    return r.trackingId
}

func (r *AliexpressAffiliateFeaturedpromoProductsGetRequest) SetCountry(country string) error {
    r.country = country
    r.Set("country", country)
    return nil
}

func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetCountry() string {
    return r.country
}

