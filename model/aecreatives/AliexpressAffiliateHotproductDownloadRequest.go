package aecreatives

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
联盟营销爆品下载接口 API请求
aliexpress.affiliate.hotproduct.download

查询联盟爆品API
*/
type AliexpressAffiliateHotproductDownloadRequest struct {
    model.Params
    // trackingId
    _trackingId   string
    // 请求签名
    _appSignature   string
    // 类目ID
    _categoryId   string
    // 返回字段列表
    _fields   string
    // 站点商品标：global,it_site,es_site,ru_site
    _localeSite   string
    // 请求页数
    _pageNo   int64
    // 每次请求数量
    _pageSize   int64
    // 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
    _targetCurrency   string
    // 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
    _targetLanguage   string
    // 收货国家，可筛选能销售至该国家的商品，并根据该国家税率政策返回对应商品价格
    _country   string
}

// 初始化AliexpressAffiliateHotproductDownloadRequest对象
func NewAliexpressAffiliateHotproductDownloadRequest() *AliexpressAffiliateHotproductDownloadRequest{
    return &AliexpressAffiliateHotproductDownloadRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressAffiliateHotproductDownloadRequest) GetApiMethodName() string {
    return "aliexpress.affiliate.hotproduct.download"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressAffiliateHotproductDownloadRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// TrackingId Setter
// trackingId
func (r *AliexpressAffiliateHotproductDownloadRequest) SetTrackingId(_trackingId string) error {
    r._trackingId = _trackingId
    r.Set("tracking_id", _trackingId)
    return nil
}

// TrackingId Getter
func (r AliexpressAffiliateHotproductDownloadRequest) GetTrackingId() string {
    return r._trackingId
}
// AppSignature Setter
// 请求签名
func (r *AliexpressAffiliateHotproductDownloadRequest) SetAppSignature(_appSignature string) error {
    r._appSignature = _appSignature
    r.Set("app_signature", _appSignature)
    return nil
}

// AppSignature Getter
func (r AliexpressAffiliateHotproductDownloadRequest) GetAppSignature() string {
    return r._appSignature
}
// CategoryId Setter
// 类目ID
func (r *AliexpressAffiliateHotproductDownloadRequest) SetCategoryId(_categoryId string) error {
    r._categoryId = _categoryId
    r.Set("category_id", _categoryId)
    return nil
}

// CategoryId Getter
func (r AliexpressAffiliateHotproductDownloadRequest) GetCategoryId() string {
    return r._categoryId
}
// Fields Setter
// 返回字段列表
func (r *AliexpressAffiliateHotproductDownloadRequest) SetFields(_fields string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r AliexpressAffiliateHotproductDownloadRequest) GetFields() string {
    return r._fields
}
// LocaleSite Setter
// 站点商品标：global,it_site,es_site,ru_site
func (r *AliexpressAffiliateHotproductDownloadRequest) SetLocaleSite(_localeSite string) error {
    r._localeSite = _localeSite
    r.Set("locale_site", _localeSite)
    return nil
}

// LocaleSite Getter
func (r AliexpressAffiliateHotproductDownloadRequest) GetLocaleSite() string {
    return r._localeSite
}
// PageNo Setter
// 请求页数
func (r *AliexpressAffiliateHotproductDownloadRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r AliexpressAffiliateHotproductDownloadRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每次请求数量
func (r *AliexpressAffiliateHotproductDownloadRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AliexpressAffiliateHotproductDownloadRequest) GetPageSize() int64 {
    return r._pageSize
}
// TargetCurrency Setter
// 目标币种:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
func (r *AliexpressAffiliateHotproductDownloadRequest) SetTargetCurrency(_targetCurrency string) error {
    r._targetCurrency = _targetCurrency
    r.Set("target_currency", _targetCurrency)
    return nil
}

// TargetCurrency Getter
func (r AliexpressAffiliateHotproductDownloadRequest) GetTargetCurrency() string {
    return r._targetCurrency
}
// TargetLanguage Setter
// 目标语言:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
func (r *AliexpressAffiliateHotproductDownloadRequest) SetTargetLanguage(_targetLanguage string) error {
    r._targetLanguage = _targetLanguage
    r.Set("target_language", _targetLanguage)
    return nil
}

// TargetLanguage Getter
func (r AliexpressAffiliateHotproductDownloadRequest) GetTargetLanguage() string {
    return r._targetLanguage
}
// Country Setter
// 收货国家，可筛选能销售至该国家的商品，并根据该国家税率政策返回对应商品价格
func (r *AliexpressAffiliateHotproductDownloadRequest) SetCountry(_country string) error {
    r._country = _country
    r.Set("country", _country)
    return nil
}

// Country Getter
func (r AliexpressAffiliateHotproductDownloadRequest) GetCountry() string {
    return r._country
}
