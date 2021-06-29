package aecreatives

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
联盟营销爆品下载接口 APIRequest
aliexpress.affiliate.hotproduct.download

查询联盟爆品API
*/
type AliexpressAffiliateHotproductDownloadRequest struct {
    model.Params

    // trackingId
    trackingId   string 

    // 请求签名
    appSignature   string 

    // 类目ID
    categoryId   string 

    // 返回字段列表
    fields   string 

    // 站点商品标：global,it_site,es_site,ru_site
    localeSite   string 

    // 请求页数
    pageNo   int64 

    // 每次请求数量
    pageSize   int64 

    // 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
    targetCurrency   string 

    // 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
    targetLanguage   string 

    // 收货国家，可筛选能销售至该国家的商品，并根据该国家税率政策返回对应商品价格
    country   string 

}

func NewAliexpressAffiliateHotproductDownloadRequest() *AliexpressAffiliateHotproductDownloadRequest{
    return &AliexpressAffiliateHotproductDownloadRequest{
        Params: model.NewParams(),
    }
}

func (r AliexpressAffiliateHotproductDownloadRequest) GetApiMethodName() string {
    return "aliexpress.affiliate.hotproduct.download"
}

func (r AliexpressAffiliateHotproductDownloadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}


func (r *AliexpressAffiliateHotproductDownloadRequest) SetTrackingId(trackingId string) error {
    r.trackingId = trackingId
    r.Set("tracking_id", trackingId)
    return nil
}

func (r AliexpressAffiliateHotproductDownloadRequest) GetTrackingId() string {
    return r.trackingId
}

func (r *AliexpressAffiliateHotproductDownloadRequest) SetAppSignature(appSignature string) error {
    r.appSignature = appSignature
    r.Set("app_signature", appSignature)
    return nil
}

func (r AliexpressAffiliateHotproductDownloadRequest) GetAppSignature() string {
    return r.appSignature
}

func (r *AliexpressAffiliateHotproductDownloadRequest) SetCategoryId(categoryId string) error {
    r.categoryId = categoryId
    r.Set("category_id", categoryId)
    return nil
}

func (r AliexpressAffiliateHotproductDownloadRequest) GetCategoryId() string {
    return r.categoryId
}

func (r *AliexpressAffiliateHotproductDownloadRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

func (r AliexpressAffiliateHotproductDownloadRequest) GetFields() string {
    return r.fields
}

func (r *AliexpressAffiliateHotproductDownloadRequest) SetLocaleSite(localeSite string) error {
    r.localeSite = localeSite
    r.Set("locale_site", localeSite)
    return nil
}

func (r AliexpressAffiliateHotproductDownloadRequest) GetLocaleSite() string {
    return r.localeSite
}

func (r *AliexpressAffiliateHotproductDownloadRequest) SetPageNo(pageNo int64) error {
    r.pageNo = pageNo
    r.Set("page_no", pageNo)
    return nil
}

func (r AliexpressAffiliateHotproductDownloadRequest) GetPageNo() int64 {
    return r.pageNo
}

func (r *AliexpressAffiliateHotproductDownloadRequest) SetPageSize(pageSize int64) error {
    r.pageSize = pageSize
    r.Set("page_size", pageSize)
    return nil
}

func (r AliexpressAffiliateHotproductDownloadRequest) GetPageSize() int64 {
    return r.pageSize
}

func (r *AliexpressAffiliateHotproductDownloadRequest) SetTargetCurrency(targetCurrency string) error {
    r.targetCurrency = targetCurrency
    r.Set("target_currency", targetCurrency)
    return nil
}

func (r AliexpressAffiliateHotproductDownloadRequest) GetTargetCurrency() string {
    return r.targetCurrency
}

func (r *AliexpressAffiliateHotproductDownloadRequest) SetTargetLanguage(targetLanguage string) error {
    r.targetLanguage = targetLanguage
    r.Set("target_language", targetLanguage)
    return nil
}

func (r AliexpressAffiliateHotproductDownloadRequest) GetTargetLanguage() string {
    return r.targetLanguage
}

func (r *AliexpressAffiliateHotproductDownloadRequest) SetCountry(country string) error {
    r.country = country
    r.Set("country", country)
    return nil
}

func (r AliexpressAffiliateHotproductDownloadRequest) GetCountry() string {
    return r.country
}

