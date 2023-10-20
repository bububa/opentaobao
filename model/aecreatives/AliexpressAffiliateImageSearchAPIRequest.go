package aecreatives

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressaffiliateimagesearchAPIRequest 图搜 API请求
// aliexpress.affiliate.image.search
//
// 图片搜索接口
type AliexpressaffiliateimagesearchAPIRequest struct {
	model.Params
	// API signature
	_appSignature string
	// 请求字段
	_fields string
	// 图片类目倾向，不填则为最佳匹配。0 - 服装；3 - 包；4 - 鞋子；88888888 - 其他类目
	_imgCid string
	// 媒体用户唯一识别号
	_mediaUserId string
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
	// 图片文件字节数组，最大不超过 100 KB
	_imageFileBytes *model.File
	// 搜索结果数量，最高 50
	_productCnt int64
}

// NewAliexpressaffiliateimagesearchRequest 初始化AliexpressaffiliateimagesearchAPIRequest对象
func NewAliexpressaffiliateimagesearchRequest() *AliexpressaffiliateimagesearchAPIRequest {
	return &AliexpressaffiliateimagesearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressaffiliateimagesearchAPIRequest) GetApiMethodName() string {
	return "aliexpress.affiliate.image.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressaffiliateimagesearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressaffiliateimagesearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetAppSignature is AppSignature Setter
// API signature
func (r *AliexpressaffiliateimagesearchAPIRequest) SetAppSignature(_appSignature string) error {
	r._appSignature = _appSignature
	r.Set("app_signature", _appSignature)
	return nil
}

// GetAppSignature AppSignature Getter
func (r AliexpressaffiliateimagesearchAPIRequest) GetAppSignature() string {
	return r._appSignature
}

// SetFields is Fields Setter
// 请求字段
func (r *AliexpressaffiliateimagesearchAPIRequest) SetFields(_fields string) error {
	r._fields = _fields
	r.Set("fields", _fields)
	return nil
}

// GetFields Fields Getter
func (r AliexpressaffiliateimagesearchAPIRequest) GetFields() string {
	return r._fields
}

// SetImgCid is ImgCid Setter
// 图片类目倾向，不填则为最佳匹配。0 - 服装；3 - 包；4 - 鞋子；88888888 - 其他类目
func (r *AliexpressaffiliateimagesearchAPIRequest) SetImgCid(_imgCid string) error {
	r._imgCid = _imgCid
	r.Set("img_cid", _imgCid)
	return nil
}

// GetImgCid ImgCid Getter
func (r AliexpressaffiliateimagesearchAPIRequest) GetImgCid() string {
	return r._imgCid
}

// SetMediaUserId is MediaUserId Setter
// 媒体用户唯一识别号
func (r *AliexpressaffiliateimagesearchAPIRequest) SetMediaUserId(_mediaUserId string) error {
	r._mediaUserId = _mediaUserId
	r.Set("media_user_id", _mediaUserId)
	return nil
}

// GetMediaUserId MediaUserId Getter
func (r AliexpressaffiliateimagesearchAPIRequest) GetMediaUserId() string {
	return r._mediaUserId
}

// SetShptTo is ShptTo Setter
// ship-to 国家
func (r *AliexpressaffiliateimagesearchAPIRequest) SetShptTo(_shptTo string) error {
	r._shptTo = _shptTo
	r.Set("shpt_to", _shptTo)
	return nil
}

// GetShptTo ShptTo Getter
func (r AliexpressaffiliateimagesearchAPIRequest) GetShptTo() string {
	return r._shptTo
}

// SetSort is Sort Setter
// 排序方式:SALE_PRICE_ASC, SALE_PRICE_DESC,LAST_VOLUME_ASC, LAST_VOLUME_DESC
func (r *AliexpressaffiliateimagesearchAPIRequest) SetSort(_sort string) error {
	r._sort = _sort
	r.Set("sort", _sort)
	return nil
}

// GetSort Sort Getter
func (r AliexpressaffiliateimagesearchAPIRequest) GetSort() string {
	return r._sort
}

// SetTargetCurrency is TargetCurrency Setter
// 目标币种:USD, GBP, CAD, EUR, UAH, MXN,TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
func (r *AliexpressaffiliateimagesearchAPIRequest) SetTargetCurrency(_targetCurrency string) error {
	r._targetCurrency = _targetCurrency
	r.Set("target_currency", _targetCurrency)
	return nil
}

// GetTargetCurrency TargetCurrency Getter
func (r AliexpressaffiliateimagesearchAPIRequest) GetTargetCurrency() string {
	return r._targetCurrency
}

// SetTargetLanguage is TargetLanguage Setter
// 目标语言:en,ru,pt,es,fr,id,it,th,ja,ar,vi,tr,de,he,ko,nl,pl,mx,cl,iw,in
func (r *AliexpressaffiliateimagesearchAPIRequest) SetTargetLanguage(_targetLanguage string) error {
	r._targetLanguage = _targetLanguage
	r.Set("target_language", _targetLanguage)
	return nil
}

// GetTargetLanguage TargetLanguage Getter
func (r AliexpressaffiliateimagesearchAPIRequest) GetTargetLanguage() string {
	return r._targetLanguage
}

// SetTrackingId is TrackingId Setter
// 媒体 trackingid
func (r *AliexpressaffiliateimagesearchAPIRequest) SetTrackingId(_trackingId string) error {
	r._trackingId = _trackingId
	r.Set("tracking_id", _trackingId)
	return nil
}

// GetTrackingId TrackingId Getter
func (r AliexpressaffiliateimagesearchAPIRequest) GetTrackingId() string {
	return r._trackingId
}

// SetImageFileBytes is ImageFileBytes Setter
// 图片文件字节数组，最大不超过 100 KB
func (r *AliexpressaffiliateimagesearchAPIRequest) SetImageFileBytes(_imageFileBytes *model.File) error {
	r._imageFileBytes = _imageFileBytes
	r.Set("image_file_bytes", _imageFileBytes)
	return nil
}

// GetImageFileBytes ImageFileBytes Getter
func (r AliexpressaffiliateimagesearchAPIRequest) GetImageFileBytes() *model.File {
	return r._imageFileBytes
}

// SetProductCnt is ProductCnt Setter
// 搜索结果数量，最高 50
func (r *AliexpressaffiliateimagesearchAPIRequest) SetProductCnt(_productCnt int64) error {
	r._productCnt = _productCnt
	r.Set("product_cnt", _productCnt)
	return nil
}

// GetProductCnt ProductCnt Getter
func (r AliexpressaffiliateimagesearchAPIRequest) GetProductCnt() int64 {
	return r._productCnt
}
