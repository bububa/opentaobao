package aecreatives

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
联盟物料智能推荐api APIRequest
aliexpress.affiliate.product.smartmatch

联盟物料算法智能推荐
*/
type AliexpressAffiliateProductSmartmatchRequest struct {
    model.Params

    // 接入APP信息
    app   string 

    // 请求签名
    appSignature   string 

    // 设备信息
    device   string 

    // adid或者idfa
    deviceId   string 

    // 返回字段列表
    fields   string 

    // 关键词
    keywords   string 

    // 商品ID
    productId   string 

    // 站点信息
    site   string 

    // 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
    targetCurrency   string 

    // 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
    targetLanguage   string 

    // trackingId
    trackingId   string 

    // 用户信息
    user   string 

    // 请求页数
    pageNo   int64 

    // 商品收货国家，可筛选能销售至该国家的商品，并根据该国家税率政策返回对应商品价格
    country   string 

}

func NewAliexpressAffiliateProductSmartmatchRequest() *AliexpressAffiliateProductSmartmatchRequest{
    return &AliexpressAffiliateProductSmartmatchRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressAffiliateProductSmartmatchRequest) GetApiMethodName() string {
    return "aliexpress.affiliate.product.smartmatch"
}

func (r AliexpressAffiliateProductSmartmatchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressAffiliateProductSmartmatchRequest) SetApp(app string) error {
    r.app = app
    r.Set("app", app)
    return nil
}

func (r AliexpressAffiliateProductSmartmatchRequest) GetApp() string {
    return r.app
}

func (r *AliexpressAffiliateProductSmartmatchRequest) SetAppSignature(appSignature string) error {
    r.appSignature = appSignature
    r.Set("app_signature", appSignature)
    return nil
}

func (r AliexpressAffiliateProductSmartmatchRequest) GetAppSignature() string {
    return r.appSignature
}

func (r *AliexpressAffiliateProductSmartmatchRequest) SetDevice(device string) error {
    r.device = device
    r.Set("device", device)
    return nil
}

func (r AliexpressAffiliateProductSmartmatchRequest) GetDevice() string {
    return r.device
}

func (r *AliexpressAffiliateProductSmartmatchRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

func (r AliexpressAffiliateProductSmartmatchRequest) GetDeviceId() string {
    return r.deviceId
}

func (r *AliexpressAffiliateProductSmartmatchRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r AliexpressAffiliateProductSmartmatchRequest) GetFields() string {
    return r.fields
}

func (r *AliexpressAffiliateProductSmartmatchRequest) SetKeywords(keywords string) error {
    r.keywords = keywords
    r.Set("keywords", keywords)
    return nil
}

func (r AliexpressAffiliateProductSmartmatchRequest) GetKeywords() string {
    return r.keywords
}

func (r *AliexpressAffiliateProductSmartmatchRequest) SetProductId(productId string) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

func (r AliexpressAffiliateProductSmartmatchRequest) GetProductId() string {
    return r.productId
}

func (r *AliexpressAffiliateProductSmartmatchRequest) SetSite(site string) error {
    r.site = site
    r.Set("site", site)
    return nil
}

func (r AliexpressAffiliateProductSmartmatchRequest) GetSite() string {
    return r.site
}

func (r *AliexpressAffiliateProductSmartmatchRequest) SetTargetCurrency(targetCurrency string) error {
    r.targetCurrency = targetCurrency
    r.Set("target_currency", targetCurrency)
    return nil
}

func (r AliexpressAffiliateProductSmartmatchRequest) GetTargetCurrency() string {
    return r.targetCurrency
}

func (r *AliexpressAffiliateProductSmartmatchRequest) SetTargetLanguage(targetLanguage string) error {
    r.targetLanguage = targetLanguage
    r.Set("target_language", targetLanguage)
    return nil
}

func (r AliexpressAffiliateProductSmartmatchRequest) GetTargetLanguage() string {
    return r.targetLanguage
}

func (r *AliexpressAffiliateProductSmartmatchRequest) SetTrackingId(trackingId string) error {
    r.trackingId = trackingId
    r.Set("tracking_id", trackingId)
    return nil
}

func (r AliexpressAffiliateProductSmartmatchRequest) GetTrackingId() string {
    return r.trackingId
}

func (r *AliexpressAffiliateProductSmartmatchRequest) SetUser(user string) error {
    r.user = user
    r.Set("user", user)
    return nil
}

func (r AliexpressAffiliateProductSmartmatchRequest) GetUser() string {
    return r.user
}

func (r *AliexpressAffiliateProductSmartmatchRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r AliexpressAffiliateProductSmartmatchRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *AliexpressAffiliateProductSmartmatchRequest) SetCountry(country string) error {
    r.country = country
    r.Set("country", country)
    return nil
}

func (r AliexpressAffiliateProductSmartmatchRequest) GetCountry() string {
    return r.country
}

