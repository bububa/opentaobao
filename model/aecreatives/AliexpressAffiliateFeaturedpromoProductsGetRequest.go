package aecreatives

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
联盟主题推广活动商品信息获取 API请求
aliexpress.affiliate.featuredpromo.products.get

根据联盟主题推广活动或主题品库查询对应的商品。如下品库为固定品库，可长期调用。品库类型和名称如下：高佣品（Hot Product）、新品（New Arrival）、热销商品（Best Seller）、每周尖货（weeklydeals）
*/
type AliexpressAffiliateFeaturedpromoProductsGetRequest struct {
    model.Params
    // 请求签名
    _appSignature   string
    // 类目 ID 如何获取category_id，请参考，https://open.taobao.com/api.htm?docId=45801&docType=2&scopeId=17063
    _categoryId   string
    // 返回字段列表
    _fields   string
    // 查询页码
    _pageNo   int64
    // 每页记录数，1-50
    _pageSize   int64
    // 活动结束时间，PST 时区
    _promotionEndTime   string
    // 主题活动的名称，如何获取主题活动，请参考"get featuredpromo info" API. 固定主题：高佣品（Hot Product）、新品（New Arrival）、热销商品（Best Seller）、每周尖货（weeklydeals）
    _promotionName   string
    // 活动开始时间，PST 时区
    _promotionStartTime   string
    // 排序方式：commissionAsc，commissionDesc, priceAsc，priceDesc，volumeAsc、volumeDesc, discountAsc, discountDesc, ratingAsc，ratingDesc, promotionTimeAsc, pr
    _sort   string
    // 目标币种，可根据目标币种返回对应币种：USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
    _targetCurrency   string
    // 目标语言，可根据目标语言返回对应语言：EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IN
    _targetLanguage   string
    // trackingID
    _trackingId   string
    // 商品收货国家，可筛选能销售至该国家的商品，并根据该国家税率政策返回对应商品价格
    _country   string
}

// 初始化AliexpressAffiliateFeaturedpromoProductsGetRequest对象
func NewAliexpressAffiliateFeaturedpromoProductsGetRequest() *AliexpressAffiliateFeaturedpromoProductsGetRequest{
    return &AliexpressAffiliateFeaturedpromoProductsGetRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetApiMethodName() string {
    return "aliexpress.affiliate.featuredpromo.products.get"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppSignature Setter
// 请求签名
func (r *AliexpressAffiliateFeaturedpromoProductsGetRequest) SetAppSignature(_appSignature string) error {
    r._appSignature = _appSignature
    r.Set("app_signature", _appSignature)
    return nil
}

// AppSignature Getter
func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetAppSignature() string {
    return r._appSignature
}
// CategoryId Setter
// 类目 ID 如何获取category_id，请参考，https://open.taobao.com/api.htm?docId=45801&docType=2&scopeId=17063
func (r *AliexpressAffiliateFeaturedpromoProductsGetRequest) SetCategoryId(_categoryId string) error {
    r._categoryId = _categoryId
    r.Set("category_id", _categoryId)
    return nil
}

// CategoryId Getter
func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetCategoryId() string {
    return r._categoryId
}
// Fields Setter
// 返回字段列表
func (r *AliexpressAffiliateFeaturedpromoProductsGetRequest) SetFields(_fields string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetFields() string {
    return r._fields
}
// PageNo Setter
// 查询页码
func (r *AliexpressAffiliateFeaturedpromoProductsGetRequest) SetPageNo(_pageNo int64) error {
    r._pageNo = _pageNo
    r.Set("page_no", _pageNo)
    return nil
}

// PageNo Getter
func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetPageNo() int64 {
    return r._pageNo
}
// PageSize Setter
// 每页记录数，1-50
func (r *AliexpressAffiliateFeaturedpromoProductsGetRequest) SetPageSize(_pageSize int64) error {
    r._pageSize = _pageSize
    r.Set("page_size", _pageSize)
    return nil
}

// PageSize Getter
func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetPageSize() int64 {
    return r._pageSize
}
// PromotionEndTime Setter
// 活动结束时间，PST 时区
func (r *AliexpressAffiliateFeaturedpromoProductsGetRequest) SetPromotionEndTime(_promotionEndTime string) error {
    r._promotionEndTime = _promotionEndTime
    r.Set("promotion_end_time", _promotionEndTime)
    return nil
}

// PromotionEndTime Getter
func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetPromotionEndTime() string {
    return r._promotionEndTime
}
// PromotionName Setter
// 主题活动的名称，如何获取主题活动，请参考"get featuredpromo info" API. 固定主题：高佣品（Hot Product）、新品（New Arrival）、热销商品（Best Seller）、每周尖货（weeklydeals）
func (r *AliexpressAffiliateFeaturedpromoProductsGetRequest) SetPromotionName(_promotionName string) error {
    r._promotionName = _promotionName
    r.Set("promotion_name", _promotionName)
    return nil
}

// PromotionName Getter
func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetPromotionName() string {
    return r._promotionName
}
// PromotionStartTime Setter
// 活动开始时间，PST 时区
func (r *AliexpressAffiliateFeaturedpromoProductsGetRequest) SetPromotionStartTime(_promotionStartTime string) error {
    r._promotionStartTime = _promotionStartTime
    r.Set("promotion_start_time", _promotionStartTime)
    return nil
}

// PromotionStartTime Getter
func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetPromotionStartTime() string {
    return r._promotionStartTime
}
// Sort Setter
// 排序方式：commissionAsc，commissionDesc, priceAsc，priceDesc，volumeAsc、volumeDesc, discountAsc, discountDesc, ratingAsc，ratingDesc, promotionTimeAsc, pr
func (r *AliexpressAffiliateFeaturedpromoProductsGetRequest) SetSort(_sort string) error {
    r._sort = _sort
    r.Set("sort", _sort)
    return nil
}

// Sort Getter
func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetSort() string {
    return r._sort
}
// TargetCurrency Setter
// 目标币种，可根据目标币种返回对应币种：USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
func (r *AliexpressAffiliateFeaturedpromoProductsGetRequest) SetTargetCurrency(_targetCurrency string) error {
    r._targetCurrency = _targetCurrency
    r.Set("target_currency", _targetCurrency)
    return nil
}

// TargetCurrency Getter
func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetTargetCurrency() string {
    return r._targetCurrency
}
// TargetLanguage Setter
// 目标语言，可根据目标语言返回对应语言：EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IN
func (r *AliexpressAffiliateFeaturedpromoProductsGetRequest) SetTargetLanguage(_targetLanguage string) error {
    r._targetLanguage = _targetLanguage
    r.Set("target_language", _targetLanguage)
    return nil
}

// TargetLanguage Getter
func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetTargetLanguage() string {
    return r._targetLanguage
}
// TrackingId Setter
// trackingID
func (r *AliexpressAffiliateFeaturedpromoProductsGetRequest) SetTrackingId(_trackingId string) error {
    r._trackingId = _trackingId
    r.Set("tracking_id", _trackingId)
    return nil
}

// TrackingId Getter
func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetTrackingId() string {
    return r._trackingId
}
// Country Setter
// 商品收货国家，可筛选能销售至该国家的商品，并根据该国家税率政策返回对应商品价格
func (r *AliexpressAffiliateFeaturedpromoProductsGetRequest) SetCountry(_country string) error {
    r._country = _country
    r.Set("country", _country)
    return nil
}

// Country Getter
func (r AliexpressAffiliateFeaturedpromoProductsGetRequest) GetCountry() string {
    return r._country
}
