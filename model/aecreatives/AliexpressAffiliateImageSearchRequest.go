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
    appSignature   string
    // 请求字段
    fields   string
    // 图片文件字节数组，最大不超过 100 KB
    imageFileBytes   []*model.File
    // 图片类目倾向，不填则为最佳匹配。0 - 服装；3 - 包；4 - 鞋子；88888888 - 其他类目
    imgCid   string
    // 媒体用户唯一识别号
    mediaUserId   string
    // 搜索结果数量，最高 50
    productCnt   int64
    // ship-to 国家
    shptTo   string
    // 排序方式:SALE_PRICE_ASC, SALE_PRICE_DESC,LAST_VOLUME_ASC, LAST_VOLUME_DESC
    sort   string
    // 目标币种:USD, GBP, CAD, EUR, UAH, MXN,TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
    targetCurrency   string
    // 目标语言:en,ru,pt,es,fr,id,it,th,ja,ar,vi,tr,de,he,ko,nl,pl,mx,cl,iw,in
    targetLanguage   string
    // 媒体 trackingid
    trackingId   string
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
func (r *AliexpressAffiliateImageSearchRequest) SetAppSignature(appSignature string) error {
    r.appSignature = appSignature
    r.Set("app_signature", appSignature)
    return nil
}

// AppSignature Getter
func (r AliexpressAffiliateImageSearchRequest) GetAppSignature() string {
    return r.appSignature
}
// Fields Setter
// 请求字段
func (r *AliexpressAffiliateImageSearchRequest) SetFields(fields string) error {
    r.fields = fields
    r.Set("fields", fields)
    return nil
}

// Fields Getter
func (r AliexpressAffiliateImageSearchRequest) GetFields() string {
    return r.fields
}
// ImageFileBytes Setter
// 图片文件字节数组，最大不超过 100 KB
func (r *AliexpressAffiliateImageSearchRequest) SetImageFileBytes(imageFileBytes []*model.File) error {
    r.imageFileBytes = imageFileBytes
    r.Set("image_file_bytes", imageFileBytes)
    return nil
}

// ImageFileBytes Getter
func (r AliexpressAffiliateImageSearchRequest) GetImageFileBytes() []*model.File {
    return r.imageFileBytes
}
// ImgCid Setter
// 图片类目倾向，不填则为最佳匹配。0 - 服装；3 - 包；4 - 鞋子；88888888 - 其他类目
func (r *AliexpressAffiliateImageSearchRequest) SetImgCid(imgCid string) error {
    r.imgCid = imgCid
    r.Set("img_cid", imgCid)
    return nil
}

// ImgCid Getter
func (r AliexpressAffiliateImageSearchRequest) GetImgCid() string {
    return r.imgCid
}
// MediaUserId Setter
// 媒体用户唯一识别号
func (r *AliexpressAffiliateImageSearchRequest) SetMediaUserId(mediaUserId string) error {
    r.mediaUserId = mediaUserId
    r.Set("media_user_id", mediaUserId)
    return nil
}

// MediaUserId Getter
func (r AliexpressAffiliateImageSearchRequest) GetMediaUserId() string {
    return r.mediaUserId
}
// ProductCnt Setter
// 搜索结果数量，最高 50
func (r *AliexpressAffiliateImageSearchRequest) SetProductCnt(productCnt int64) error {
    r.productCnt = productCnt
    r.Set("product_cnt", productCnt)
    return nil
}

// ProductCnt Getter
func (r AliexpressAffiliateImageSearchRequest) GetProductCnt() int64 {
    return r.productCnt
}
// ShptTo Setter
// ship-to 国家
func (r *AliexpressAffiliateImageSearchRequest) SetShptTo(shptTo string) error {
    r.shptTo = shptTo
    r.Set("shpt_to", shptTo)
    return nil
}

// ShptTo Getter
func (r AliexpressAffiliateImageSearchRequest) GetShptTo() string {
    return r.shptTo
}
// Sort Setter
// 排序方式:SALE_PRICE_ASC, SALE_PRICE_DESC,LAST_VOLUME_ASC, LAST_VOLUME_DESC
func (r *AliexpressAffiliateImageSearchRequest) SetSort(sort string) error {
    r.sort = sort
    r.Set("sort", sort)
    return nil
}

// Sort Getter
func (r AliexpressAffiliateImageSearchRequest) GetSort() string {
    return r.sort
}
// TargetCurrency Setter
// 目标币种:USD, GBP, CAD, EUR, UAH, MXN,TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
func (r *AliexpressAffiliateImageSearchRequest) SetTargetCurrency(targetCurrency string) error {
    r.targetCurrency = targetCurrency
    r.Set("target_currency", targetCurrency)
    return nil
}

// TargetCurrency Getter
func (r AliexpressAffiliateImageSearchRequest) GetTargetCurrency() string {
    return r.targetCurrency
}
// TargetLanguage Setter
// 目标语言:en,ru,pt,es,fr,id,it,th,ja,ar,vi,tr,de,he,ko,nl,pl,mx,cl,iw,in
func (r *AliexpressAffiliateImageSearchRequest) SetTargetLanguage(targetLanguage string) error {
    r.targetLanguage = targetLanguage
    r.Set("target_language", targetLanguage)
    return nil
}

// TargetLanguage Getter
func (r AliexpressAffiliateImageSearchRequest) GetTargetLanguage() string {
    return r.targetLanguage
}
// TrackingId Setter
// 媒体 trackingid
func (r *AliexpressAffiliateImageSearchRequest) SetTrackingId(trackingId string) error {
    r.trackingId = trackingId
    r.Set("tracking_id", trackingId)
    return nil
}

// TrackingId Getter
func (r AliexpressAffiliateImageSearchRequest) GetTrackingId() string {
    return r.trackingId
}
