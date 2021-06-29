package aecreatives

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
联盟商品详情获取接口 APIRequest
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

func NewAliexpressAffiliateProductdetailGetRequest() *AliexpressAffiliateProductdetailGetRequest{
    return &AliexpressAffiliateProductdetailGetRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressAffiliateProductdetailGetRequest) GetApiMethodName() string {
    return "aliexpress.affiliate.productdetail.get"
}

func (r AliexpressAffiliateProductdetailGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressAffiliateProductdetailGetRequest) SetAppSignature(appSignature string) error {
    r.appSignature = appSignature
    r.Set("app_signature", appSignature)
    return nil
}

func (r AliexpressAffiliateProductdetailGetRequest) GetAppSignature() string {
    return r.appSignature
}

func (r *AliexpressAffiliateProductdetailGetRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r AliexpressAffiliateProductdetailGetRequest) GetFields() string {
    return r.fields
}

func (r *AliexpressAffiliateProductdetailGetRequest) SetProductIds(productIds string) error {
    r.productIds = productIds
    r.Set("product_ids", productIds)
    return nil
}

func (r AliexpressAffiliateProductdetailGetRequest) GetProductIds() string {
    return r.productIds
}

func (r *AliexpressAffiliateProductdetailGetRequest) SetTargetCurrency(targetCurrency string) error {
    r.targetCurrency = targetCurrency
    r.Set("target_currency", targetCurrency)
    return nil
}

func (r AliexpressAffiliateProductdetailGetRequest) GetTargetCurrency() string {
    return r.targetCurrency
}

func (r *AliexpressAffiliateProductdetailGetRequest) SetTargetLanguage(targetLanguage string) error {
    r.targetLanguage = targetLanguage
    r.Set("target_language", targetLanguage)
    return nil
}

func (r AliexpressAffiliateProductdetailGetRequest) GetTargetLanguage() string {
    return r.targetLanguage
}

func (r *AliexpressAffiliateProductdetailGetRequest) SetTrackingId(trackingId string) error {
    r.trackingId = trackingId
    r.Set("tracking_id", trackingId)
    return nil
}

func (r AliexpressAffiliateProductdetailGetRequest) GetTrackingId() string {
    return r.trackingId
}

func (r *AliexpressAffiliateProductdetailGetRequest) SetCountry(country string) error {
    r.country = country
    r.Set("country", country)
    return nil
}

func (r AliexpressAffiliateProductdetailGetRequest) GetCountry() string {
    return r.country
}

