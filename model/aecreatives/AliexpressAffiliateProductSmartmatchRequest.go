package aecreatives

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
联盟物料智能推荐api API请求
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

// 初始化AliexpressAffiliateProductSmartmatchRequest对象
func NewAliexpressAffiliateProductSmartmatchRequest() *AliexpressAffiliateProductSmartmatchRequest{
    return &AliexpressAffiliateProductSmartmatchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressAffiliateProductSmartmatchRequest) GetApiMethodName() string {
    return "aliexpress.affiliate.product.smartmatch"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressAffiliateProductSmartmatchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// App Setter
// 接入APP信息
func (r *AliexpressAffiliateProductSmartmatchRequest) SetApp(app string) error {
    r.app = app
    r.Set("app", app)
    return nil
}

// App Getter
func (r AliexpressAffiliateProductSmartmatchRequest) GetApp() string {
    return r.app
}
// AppSignature Setter
// 请求签名
func (r *AliexpressAffiliateProductSmartmatchRequest) SetAppSignature(appSignature string) error {
    r.appSignature = appSignature
    r.Set("app_signature", appSignature)
    return nil
}

// AppSignature Getter
func (r AliexpressAffiliateProductSmartmatchRequest) GetAppSignature() string {
    return r.appSignature
}
// Device Setter
// 设备信息
func (r *AliexpressAffiliateProductSmartmatchRequest) SetDevice(device string) error {
    r.device = device
    r.Set("device", device)
    return nil
}

// Device Getter
func (r AliexpressAffiliateProductSmartmatchRequest) GetDevice() string {
    return r.device
}
// DeviceId Setter
// adid或者idfa
func (r *AliexpressAffiliateProductSmartmatchRequest) SetDeviceId(deviceId string) error {
    r.deviceId = deviceId
    r.Set("device_id", deviceId)
    return nil
}

// DeviceId Getter
func (r AliexpressAffiliateProductSmartmatchRequest) GetDeviceId() string {
    return r.deviceId
}
// Fields Setter
// 返回字段列表
func (r *AliexpressAffiliateProductSmartmatchRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r AliexpressAffiliateProductSmartmatchRequest) GetFields() string {
    return r.fields
}
// Keywords Setter
// 关键词
func (r *AliexpressAffiliateProductSmartmatchRequest) SetKeywords(keywords string) error {
    r.keywords = keywords
    r.Set("keywords", keywords)
    return nil
}

// Keywords Getter
func (r AliexpressAffiliateProductSmartmatchRequest) GetKeywords() string {
    return r.keywords
}
// ProductId Setter
// 商品ID
func (r *AliexpressAffiliateProductSmartmatchRequest) SetProductId(productId string) error {
    r.productId = productId
    r.Set("product_id", productId)
    return nil
}

// ProductId Getter
func (r AliexpressAffiliateProductSmartmatchRequest) GetProductId() string {
    return r.productId
}
// Site Setter
// 站点信息
func (r *AliexpressAffiliateProductSmartmatchRequest) SetSite(site string) error {
    r.site = site
    r.Set("site", site)
    return nil
}

// Site Getter
func (r AliexpressAffiliateProductSmartmatchRequest) GetSite() string {
    return r.site
}
// TargetCurrency Setter
// 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
func (r *AliexpressAffiliateProductSmartmatchRequest) SetTargetCurrency(targetCurrency string) error {
    r.targetCurrency = targetCurrency
    r.Set("target_currency", targetCurrency)
    return nil
}

// TargetCurrency Getter
func (r AliexpressAffiliateProductSmartmatchRequest) GetTargetCurrency() string {
    return r.targetCurrency
}
// TargetLanguage Setter
// 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
func (r *AliexpressAffiliateProductSmartmatchRequest) SetTargetLanguage(targetLanguage string) error {
    r.targetLanguage = targetLanguage
    r.Set("target_language", targetLanguage)
    return nil
}

// TargetLanguage Getter
func (r AliexpressAffiliateProductSmartmatchRequest) GetTargetLanguage() string {
    return r.targetLanguage
}
// TrackingId Setter
// trackingId
func (r *AliexpressAffiliateProductSmartmatchRequest) SetTrackingId(trackingId string) error {
    r.trackingId = trackingId
    r.Set("tracking_id", trackingId)
    return nil
}

// TrackingId Getter
func (r AliexpressAffiliateProductSmartmatchRequest) GetTrackingId() string {
    return r.trackingId
}
// User Setter
// 用户信息
func (r *AliexpressAffiliateProductSmartmatchRequest) SetUser(user string) error {
    r.user = user
    r.Set("user", user)
    return nil
}

// User Getter
func (r AliexpressAffiliateProductSmartmatchRequest) GetUser() string {
    return r.user
}
// PageNo Setter
// 请求页数
func (r *AliexpressAffiliateProductSmartmatchRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

// PageNo Getter
func (r AliexpressAffiliateProductSmartmatchRequest) GetPageNo() int64 {
    return r.pageNo
}
// Country Setter
// 商品收货国家，可筛选能销售至该国家的商品，并根据该国家税率政策返回对应商品价格
func (r *AliexpressAffiliateProductSmartmatchRequest) SetCountry(country string) error {
    r.country = country
    r.Set("country", country)
    return nil
}

// Country Getter
func (r AliexpressAffiliateProductSmartmatchRequest) GetCountry() string {
    return r.country
}
