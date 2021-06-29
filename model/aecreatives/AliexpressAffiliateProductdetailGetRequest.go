package aecreatives

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
联盟商品详情获取接口 API请求
aliexpress.affiliate.productdetail.get

联盟推广商品搜索接口，用于搜索联盟推广商品数据
*/
type AliexpressAffiliateProductdetailGetRequest struct {
    model.Params
    // 安全签名
    appSignature   string
    // commission_rate,sale_price
    fields   string
    // 商品ID列表
    productIds   string
    // 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
    targetCurrency   string
    // 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
    targetLanguage   string
    // trackingId
    trackingId   string
    // 商品收货国家，可筛选能销售至该国家的商品，并根据该国家税率政策返回对应商品价格
    country   string
}

// 初始化AliexpressAffiliateProductdetailGetRequest对象
func NewAliexpressAffiliateProductdetailGetRequest() *AliexpressAffiliateProductdetailGetRequest{
    return &AliexpressAffiliateProductdetailGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressAffiliateProductdetailGetRequest) GetApiMethodName() string {
    return "aliexpress.affiliate.productdetail.get"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressAffiliateProductdetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppSignature Setter
// 安全签名
func (r *AliexpressAffiliateProductdetailGetRequest) SetAppSignature(appSignature string) error {
    r.appSignature = appSignature
    r.Set("app_signature", appSignature)
    return nil
}

// AppSignature Getter
func (r AliexpressAffiliateProductdetailGetRequest) GetAppSignature() string {
    return r.appSignature
}
// Fields Setter
// commission_rate,sale_price
func (r *AliexpressAffiliateProductdetailGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r AliexpressAffiliateProductdetailGetRequest) GetFields() string {
    return r.fields
}
// ProductIds Setter
// 商品ID列表
func (r *AliexpressAffiliateProductdetailGetRequest) SetProductIds(productIds string) error {
    r.productIds = productIds
    r.Set("product_ids", productIds)
    return nil
}

// ProductIds Getter
func (r AliexpressAffiliateProductdetailGetRequest) GetProductIds() string {
    return r.productIds
}
// TargetCurrency Setter
// 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
func (r *AliexpressAffiliateProductdetailGetRequest) SetTargetCurrency(targetCurrency string) error {
    r.targetCurrency = targetCurrency
    r.Set("target_currency", targetCurrency)
    return nil
}

// TargetCurrency Getter
func (r AliexpressAffiliateProductdetailGetRequest) GetTargetCurrency() string {
    return r.targetCurrency
}
// TargetLanguage Setter
// 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
func (r *AliexpressAffiliateProductdetailGetRequest) SetTargetLanguage(targetLanguage string) error {
    r.targetLanguage = targetLanguage
    r.Set("target_language", targetLanguage)
    return nil
}

// TargetLanguage Getter
func (r AliexpressAffiliateProductdetailGetRequest) GetTargetLanguage() string {
    return r.targetLanguage
}
// TrackingId Setter
// trackingId
func (r *AliexpressAffiliateProductdetailGetRequest) SetTrackingId(trackingId string) error {
    r.trackingId = trackingId
    r.Set("tracking_id", trackingId)
    return nil
}

// TrackingId Getter
func (r AliexpressAffiliateProductdetailGetRequest) GetTrackingId() string {
    return r.trackingId
}
// Country Setter
// 商品收货国家，可筛选能销售至该国家的商品，并根据该国家税率政策返回对应商品价格
func (r *AliexpressAffiliateProductdetailGetRequest) SetCountry(country string) error {
    r.country = country
    r.Set("country", country)
    return nil
}

// Country Getter
func (r AliexpressAffiliateProductdetailGetRequest) GetCountry() string {
    return r.country
}
