package aedropshiper

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressDsImageSearchAPIRequest 图片搜索 API请求
// aliexpress.ds.image.search
//
// 图片搜索
type AliexpressDsImageSearchAPIRequest struct {
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

// NewAliexpressDsImageSearchRequest 初始化AliexpressDsImageSearchAPIRequest对象
func NewAliexpressDsImageSearchRequest() *AliexpressDsImageSearchAPIRequest {
	return &AliexpressDsImageSearchAPIRequest{
		Params: model.NewParams(6),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AliexpressDsImageSearchAPIRequest) Reset() {
	r._targetLanguage = ""
	r._targetCurrency = ""
	r._sort = ""
	r._shptTo = ""
	r._productCnt = 0
	r._imageFileBytes = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressDsImageSearchAPIRequest) GetApiMethodName() string {
	return "aliexpress.ds.image.search"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressDsImageSearchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpressDsImageSearchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTargetLanguage is TargetLanguage Setter
// target_language:EN,RU,PT,ES,FR,ID,IT,TH,JA,AR,VI,TR,DE,HE,KO,NL,PL,MX,CL,IW,IN
func (r *AliexpressDsImageSearchAPIRequest) SetTargetLanguage(_targetLanguage string) error {
	r._targetLanguage = _targetLanguage
	r.Set("target_language", _targetLanguage)
	return nil
}

// GetTargetLanguage TargetLanguage Getter
func (r AliexpressDsImageSearchAPIRequest) GetTargetLanguage() string {
	return r._targetLanguage
}

// SetTargetCurrency is TargetCurrency Setter
// target_currency:USD, GBP, CAD, EUR, UAH, MXN, TRY, RUB, BRL, AUD, INR, JPY, IDR, SEK,KRW
func (r *AliexpressDsImageSearchAPIRequest) SetTargetCurrency(_targetCurrency string) error {
	r._targetCurrency = _targetCurrency
	r.Set("target_currency", _targetCurrency)
	return nil
}

// GetTargetCurrency TargetCurrency Getter
func (r AliexpressDsImageSearchAPIRequest) GetTargetCurrency() string {
	return r._targetCurrency
}

// SetSort is Sort Setter
// SALE_PRICE_ASC, SALE_PRICE_DESC, LAST_VOLUME_ASC, LAST_VOLUME_DESC
func (r *AliexpressDsImageSearchAPIRequest) SetSort(_sort string) error {
	r._sort = _sort
	r.Set("sort", _sort)
	return nil
}

// GetSort Sort Getter
func (r AliexpressDsImageSearchAPIRequest) GetSort() string {
	return r._sort
}

// SetShptTo is ShptTo Setter
// optional  Ship to Country
func (r *AliexpressDsImageSearchAPIRequest) SetShptTo(_shptTo string) error {
	r._shptTo = _shptTo
	r.Set("shpt_to", _shptTo)
	return nil
}

// GetShptTo ShptTo Getter
func (r AliexpressDsImageSearchAPIRequest) GetShptTo() string {
	return r._shptTo
}

// SetProductCnt is ProductCnt Setter
// count of products， max 150.
func (r *AliexpressDsImageSearchAPIRequest) SetProductCnt(_productCnt int64) error {
	r._productCnt = _productCnt
	r.Set("product_cnt", _productCnt)
	return nil
}

// GetProductCnt ProductCnt Getter
func (r AliexpressDsImageSearchAPIRequest) GetProductCnt() int64 {
	return r._productCnt
}

// SetImageFileBytes is ImageFileBytes Setter
// image name in fileserver，max size 100 KB
func (r *AliexpressDsImageSearchAPIRequest) SetImageFileBytes(_imageFileBytes *model.File) error {
	r._imageFileBytes = _imageFileBytes
	r.Set("image_file_bytes", _imageFileBytes)
	return nil
}

// GetImageFileBytes ImageFileBytes Getter
func (r AliexpressDsImageSearchAPIRequest) GetImageFileBytes() *model.File {
	return r._imageFileBytes
}

var poolAliexpressDsImageSearchAPIRequest = sync.Pool{
	New: func() any {
		return NewAliexpressDsImageSearchRequest()
	},
}

// GetAliexpressDsImageSearchRequest 从 sync.Pool 获取 AliexpressDsImageSearchAPIRequest
func GetAliexpressDsImageSearchAPIRequest() *AliexpressDsImageSearchAPIRequest {
	return poolAliexpressDsImageSearchAPIRequest.Get().(*AliexpressDsImageSearchAPIRequest)
}

// ReleaseAliexpressDsImageSearchAPIRequest 将 AliexpressDsImageSearchAPIRequest 放入 sync.Pool
func ReleaseAliexpressDsImageSearchAPIRequest(v *AliexpressDsImageSearchAPIRequest) {
	v.Reset()
	poolAliexpressDsImageSearchAPIRequest.Put(v)
}
