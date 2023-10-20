package aedropshiper

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressdsimagesearchAPIRequest 图片搜索 API请求
// aliexpress.ds.image.search
//
// 图片搜索
type AliexpressdsimagesearchAPIRequest struct {
	model.Params
	// target_language:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
	_targetLanguage string
	// target_currency:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
	_targetCurrency string
	// SALE_PRICE_ASC, SALE_PRICE_DESC, LAST_VOLUME_ASC, LAST_VOLUME_DESC
	_sort string
	// optional  Ship to Country
	_shptTo string
	// count of products， max 150.
	_productCnt int64
	// image name in fileserver，max size 100 KB
	_imageFileBytes *model.File
}

// NewAliexpressdsimagesearchRequest 初始化AliexpressdsimagesearchAPIRequest对象
func NewAliexpressdsimagesearchRequest() *AliexpressdsimagesearchAPIRequest {
	return &AliexpressdsimagesearchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressdsimagesearchAPIRequest) GetApiMethodName() string {
	return "aliexpress.ds.image.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressdsimagesearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressdsimagesearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTargetLanguage is TargetLanguage Setter
// target_language:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
func (r *AliexpressdsimagesearchAPIRequest) SetTargetLanguage(_targetLanguage string) error {
	r._targetLanguage = _targetLanguage
	r.Set("target_language", _targetLanguage)
	return nil
}

// GetTargetLanguage TargetLanguage Getter
func (r AliexpressdsimagesearchAPIRequest) GetTargetLanguage() string {
	return r._targetLanguage
}

// SetTargetCurrency is TargetCurrency Setter
// target_currency:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
func (r *AliexpressdsimagesearchAPIRequest) SetTargetCurrency(_targetCurrency string) error {
	r._targetCurrency = _targetCurrency
	r.Set("target_currency", _targetCurrency)
	return nil
}

// GetTargetCurrency TargetCurrency Getter
func (r AliexpressdsimagesearchAPIRequest) GetTargetCurrency() string {
	return r._targetCurrency
}

// SetSort is Sort Setter
// SALE_PRICE_ASC, SALE_PRICE_DESC, LAST_VOLUME_ASC, LAST_VOLUME_DESC
func (r *AliexpressdsimagesearchAPIRequest) SetSort(_sort string) error {
	r._sort = _sort
	r.Set("sort", _sort)
	return nil
}

// GetSort Sort Getter
func (r AliexpressdsimagesearchAPIRequest) GetSort() string {
	return r._sort
}

// SetShptTo is ShptTo Setter
// optional  Ship to Country
func (r *AliexpressdsimagesearchAPIRequest) SetShptTo(_shptTo string) error {
	r._shptTo = _shptTo
	r.Set("shpt_to", _shptTo)
	return nil
}

// GetShptTo ShptTo Getter
func (r AliexpressdsimagesearchAPIRequest) GetShptTo() string {
	return r._shptTo
}

// SetProductCnt is ProductCnt Setter
// count of products， max 150.
func (r *AliexpressdsimagesearchAPIRequest) SetProductCnt(_productCnt int64) error {
	r._productCnt = _productCnt
	r.Set("product_cnt", _productCnt)
	return nil
}

// GetProductCnt ProductCnt Getter
func (r AliexpressdsimagesearchAPIRequest) GetProductCnt() int64 {
	return r._productCnt
}

// SetImageFileBytes is ImageFileBytes Setter
// image name in fileserver，max size 100 KB
func (r *AliexpressdsimagesearchAPIRequest) SetImageFileBytes(_imageFileBytes *model.File) error {
	r._imageFileBytes = _imageFileBytes
	r.Set("image_file_bytes", _imageFileBytes)
	return nil
}

// GetImageFileBytes ImageFileBytes Getter
func (r AliexpressdsimagesearchAPIRequest) GetImageFileBytes() *model.File {
	return r._imageFileBytes
}
