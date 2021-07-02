package aecreatives

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressAffiliateImageSearchAPIRequest 图搜 API请求
// aliexpress.affiliate.image.search
//
// 图片搜索接口
type AliexpressAffiliateImageSearchAPIRequest struct {
	model.Params
	// API signature
	_appSignature string
	// 请求字段
	_fields string
	// 图片文件字节数组，最大不超过 100 KB
	_imageFileBytes *model.File
	// 图片类目倾向，不填则为最佳匹配。0 - 服装；3 - 包；4 - 鞋子；88888888 - 其他类目
	_imgCid string
	// 媒体用户唯一识别号
	_mediaUserId string
	// 搜索结果数量，最高 50
	_productCnt int64
	// ship-to 国家
	_shptTo string
	// 排序方式:SALE_PRICE_ASC, SALE_PRICE_DESC,LAST_VOLUME_ASC, LAST_VOLUME_DESC
	_sort string
	// 目标币种:USD, GBP, CAD, EUR, UAH, MXN,TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
	_targetCurrency string
	// 目标语言:en,ru,pt,es,fr,id,it,th,ja,ar,vi,tr,de,he,ko,nl,pl,mx,cl,iw,in
	_targetLanguage string
	// 媒体 trackingid
	_trackingId string
}

// NewAliexpressAffiliateImageSearchRequest 初始化AliexpressAffiliateImageSearchAPIRequest对象
func NewAliexpressAffiliateImageSearchRequest() *AliexpressAffiliateImageSearchAPIRequest {
	return &AliexpressAffiliateImageSearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressAffiliateImageSearchAPIRequest) GetApiMethodName() string {
	return "aliexpress.affiliate.image.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressAffiliateImageSearchAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is AppSignature Setter
// API signature
func (r *AliexpressAffiliateImageSearchAPIRequest) SetAppSignature(_appSignature string) error {
	r._appSignature = _appSignature
	r.Set("app_signature", _appSignature)
	return nil
}

// Get AppSignature Getter
func (r AliexpressAffiliateImageSearchAPIRequest) GetAppSignature() string {
	return r._appSignature
}

// Set is Fields Setter
// 请求字段
func (r *AliexpressAffiliateImageSearchAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// Get Fields Getter
func (r AliexpressAffiliateImageSearchAPIRequest) GetFields() string {
	return r._fields
}

// Set is ImageFileBytes Setter
// 图片文件字节数组，最大不超过 100 KB
func (r *AliexpressAffiliateImageSearchAPIRequest) SetImageFileBytes(_imageFileBytes *model.File) error {
	r._imageFileBytes = _imageFileBytes
	r.Set("image_file_bytes", _imageFileBytes)
	return nil
}

// Get ImageFileBytes Getter
func (r AliexpressAffiliateImageSearchAPIRequest) GetImageFileBytes() *model.File {
	return r._imageFileBytes
}

// Set is ImgCid Setter
// 图片类目倾向，不填则为最佳匹配。0 - 服装；3 - 包；4 - 鞋子；88888888 - 其他类目
func (r *AliexpressAffiliateImageSearchAPIRequest) SetImgCid(_imgCid string) error {
	r._imgCid = _imgCid
	r.Set("img_cid", _imgCid)
	return nil
}

// Get ImgCid Getter
func (r AliexpressAffiliateImageSearchAPIRequest) GetImgCid() string {
	return r._imgCid
}

// Set is MediaUserId Setter
// 媒体用户唯一识别号
func (r *AliexpressAffiliateImageSearchAPIRequest) SetMediaUserId(_mediaUserId string) error {
	r._mediaUserId = _mediaUserId
	r.Set("media_user_id", _mediaUserId)
	return nil
}

// Get MediaUserId Getter
func (r AliexpressAffiliateImageSearchAPIRequest) GetMediaUserId() string {
	return r._mediaUserId
}

// Set is ProductCnt Setter
// 搜索结果数量，最高 50
func (r *AliexpressAffiliateImageSearchAPIRequest) SetProductCnt(_productCnt int64) error {
	r._productCnt = _productCnt
	r.Set("product_cnt", _productCnt)
	return nil
}

// Get ProductCnt Getter
func (r AliexpressAffiliateImageSearchAPIRequest) GetProductCnt() int64 {
	return r._productCnt
}

// Set is ShptTo Setter
// ship-to 国家
func (r *AliexpressAffiliateImageSearchAPIRequest) SetShptTo(_shptTo string) error {
	r._shptTo = _shptTo
	r.Set("shpt_to", _shptTo)
	return nil
}

// Get ShptTo Getter
func (r AliexpressAffiliateImageSearchAPIRequest) GetShptTo() string {
	return r._shptTo
}

// Set is Sort Setter
// 排序方式:SALE_PRICE_ASC, SALE_PRICE_DESC,LAST_VOLUME_ASC, LAST_VOLUME_DESC
func (r *AliexpressAffiliateImageSearchAPIRequest) SetSort(_sort string) error {
	r._sort = _sort
	r.Set("sort", _sort)
	return nil
}

// Get Sort Getter
func (r AliexpressAffiliateImageSearchAPIRequest) GetSort() string {
	return r._sort
}

// Set is TargetCurrency Setter
// 目标币种:USD, GBP, CAD, EUR, UAH, MXN,TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
func (r *AliexpressAffiliateImageSearchAPIRequest) SetTargetCurrency(_targetCurrency string) error {
	r._targetCurrency = _targetCurrency
	r.Set("target_currency", _targetCurrency)
	return nil
}

// Get TargetCurrency Getter
func (r AliexpressAffiliateImageSearchAPIRequest) GetTargetCurrency() string {
	return r._targetCurrency
}

// Set is TargetLanguage Setter
// 目标语言:en,ru,pt,es,fr,id,it,th,ja,ar,vi,tr,de,he,ko,nl,pl,mx,cl,iw,in
func (r *AliexpressAffiliateImageSearchAPIRequest) SetTargetLanguage(_targetLanguage string) error {
	r._targetLanguage = _targetLanguage
	r.Set("target_language", _targetLanguage)
	return nil
}

// Get TargetLanguage Getter
func (r AliexpressAffiliateImageSearchAPIRequest) GetTargetLanguage() string {
	return r._targetLanguage
}

// Set is TrackingId Setter
// 媒体 trackingid
func (r *AliexpressAffiliateImageSearchAPIRequest) SetTrackingId(_trackingId string) error {
	r._trackingId = _trackingId
	r.Set("tracking_id", _trackingId)
	return nil
}

// Get TrackingId Getter
func (r AliexpressAffiliateImageSearchAPIRequest) GetTrackingId() string {
	return r._trackingId
}
