package aecreatives

import (
    "net/url"

    "github.com/bububa/opentaobao/model"
)

/* 
图搜 API请求
aliexpress.affiliate.image.search

图片搜索接口
*/
type AliexpressAffiliateImageSearchRequest struct {
    model.Params
    // API signature
    _appSignature   string
    // 请求字段
    _fields   string
    // 图片文件字节数组，最大不超过 100 KB
    _imageFileBytes   []*model.File
    // 图片类目倾向，不填则为最佳匹配。0 - 服装；3 - 包；4 - 鞋子；88888888 - 其他类目
    _imgCid   string
    // 媒体用户唯一识别号
    _mediaUserId   string
    // 搜索结果数量，最高 50
    _productCnt   int64
    // ship-to 国家
    _shptTo   string
    // 排序方式:SALE_PRICE_ASC, SALE_PRICE_DESC,LAST_VOLUME_ASC, LAST_VOLUME_DESC
    _sort   string
    // 目标币种:USD, GBP, CAD, EUR, UAH, MXN,TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
    _targetCurrency   string
    // 目标语言:en,ru,pt,es,fr,id,it,th,ja,ar,vi,tr,de,he,ko,nl,pl,mx,cl,iw,in
    _targetLanguage   string
    // 媒体 trackingid
    _trackingId   string
}

// 初始化AliexpressAffiliateImageSearchRequest对象
func NewAliexpressAffiliateImageSearchRequest() *AliexpressAffiliateImageSearchRequest{
    return &AliexpressAffiliateImageSearchRequest{
        Params: model.NewParams(),
    }
}

// IRequest interface 方法, 获取Api method
func (r AliexpressAffiliateImageSearchRequest) GetApiMethodName() string {
    return "aliexpress.affiliate.image.search"
}

// IRequest interface 方法, 获取API参数
func (r AliexpressAffiliateImageSearchRequest) GetApiParams() url.Values {
    params := url.Values{}
    for k, v := range r.GetRawParams() {
        params.Set(k, v.String())
    }
    return params
}
// AppSignature Setter
// API signature
func (r *AliexpressAffiliateImageSearchRequest) SetAppSignature(_appSignature string) error {
    r._appSignature = _appSignature
    r.Set("app_signature", _appSignature)
    return nil
}

// AppSignature Getter
func (r AliexpressAffiliateImageSearchRequest) GetAppSignature() string {
    return r._appSignature
}
// Fields Setter
// 请求字段
func (r *AliexpressAffiliateImageSearchRequest) SetFields(_fields string) error {
    r._fields = _fields
    r.Set("fields", _fields)
    return nil
}

// Fields Getter
func (r AliexpressAffiliateImageSearchRequest) GetFields() string {
    return r._fields
}
// ImageFileBytes Setter
// 图片文件字节数组，最大不超过 100 KB
func (r *AliexpressAffiliateImageSearchRequest) SetImageFileBytes(_imageFileBytes []*model.File) error {
    r._imageFileBytes = _imageFileBytes
    r.Set("image_file_bytes", _imageFileBytes)
    return nil
}

// ImageFileBytes Getter
func (r AliexpressAffiliateImageSearchRequest) GetImageFileBytes() []*model.File {
    return r._imageFileBytes
}
// ImgCid Setter
// 图片类目倾向，不填则为最佳匹配。0 - 服装；3 - 包；4 - 鞋子；88888888 - 其他类目
func (r *AliexpressAffiliateImageSearchRequest) SetImgCid(_imgCid string) error {
    r._imgCid = _imgCid
    r.Set("img_cid", _imgCid)
    return nil
}

// ImgCid Getter
func (r AliexpressAffiliateImageSearchRequest) GetImgCid() string {
    return r._imgCid
}
// MediaUserId Setter
// 媒体用户唯一识别号
func (r *AliexpressAffiliateImageSearchRequest) SetMediaUserId(_mediaUserId string) error {
    r._mediaUserId = _mediaUserId
    r.Set("media_user_id", _mediaUserId)
    return nil
}

// MediaUserId Getter
func (r AliexpressAffiliateImageSearchRequest) GetMediaUserId() string {
    return r._mediaUserId
}
// ProductCnt Setter
// 搜索结果数量，最高 50
func (r *AliexpressAffiliateImageSearchRequest) SetProductCnt(_productCnt int64) error {
    r._productCnt = _productCnt
    r.Set("product_cnt", _productCnt)
    return nil
}

// ProductCnt Getter
func (r AliexpressAffiliateImageSearchRequest) GetProductCnt() int64 {
    return r._productCnt
}
// ShptTo Setter
// ship-to 国家
func (r *AliexpressAffiliateImageSearchRequest) SetShptTo(_shptTo string) error {
    r._shptTo = _shptTo
    r.Set("shpt_to", _shptTo)
    return nil
}

// ShptTo Getter
func (r AliexpressAffiliateImageSearchRequest) GetShptTo() string {
    return r._shptTo
}
// Sort Setter
// 排序方式:SALE_PRICE_ASC, SALE_PRICE_DESC,LAST_VOLUME_ASC, LAST_VOLUME_DESC
func (r *AliexpressAffiliateImageSearchRequest) SetSort(_sort string) error {
    r._sort = _sort
    r.Set("sort", _sort)
    return nil
}

// Sort Getter
func (r AliexpressAffiliateImageSearchRequest) GetSort() string {
    return r._sort
}
// TargetCurrency Setter
// 目标币种:USD, GBP, CAD, EUR, UAH, MXN,TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
func (r *AliexpressAffiliateImageSearchRequest) SetTargetCurrency(_targetCurrency string) error {
    r._targetCurrency = _targetCurrency
    r.Set("target_currency", _targetCurrency)
    return nil
}

// TargetCurrency Getter
func (r AliexpressAffiliateImageSearchRequest) GetTargetCurrency() string {
    return r._targetCurrency
}
// TargetLanguage Setter
// 目标语言:en,ru,pt,es,fr,id,it,th,ja,ar,vi,tr,de,he,ko,nl,pl,mx,cl,iw,in
func (r *AliexpressAffiliateImageSearchRequest) SetTargetLanguage(_targetLanguage string) error {
    r._targetLanguage = _targetLanguage
    r.Set("target_language", _targetLanguage)
    return nil
}

// TargetLanguage Getter
func (r AliexpressAffiliateImageSearchRequest) GetTargetLanguage() string {
    return r._targetLanguage
}
// TrackingId Setter
// 媒体 trackingid
func (r *AliexpressAffiliateImageSearchRequest) SetTrackingId(_trackingId string) error {
    r._trackingId = _trackingId
    r.Set("tracking_id", _trackingId)
    return nil
}

// TrackingId Getter
func (r AliexpressAffiliateImageSearchRequest) GetTrackingId() string {
    return r._trackingId
}
