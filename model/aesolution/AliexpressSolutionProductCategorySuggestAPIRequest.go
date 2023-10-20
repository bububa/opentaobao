package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpresssolutionproductcategorysuggestAPIRequest Suggest product categories API请求
// aliexpress.solution.product.category.suggest
//
// Suggest product categories by title and image.
type AliexpresssolutionproductcategorysuggestAPIRequest struct {
	model.Params
	// product title
	_title string
	// language
	_language string
	// product main image
	_imageUrl string
}

// NewAliexpresssolutionproductcategorysuggestRequest 初始化AliexpresssolutionproductcategorysuggestAPIRequest对象
func NewAliexpresssolutionproductcategorysuggestRequest() *AliexpresssolutionproductcategorysuggestAPIRequest {
	return &AliexpresssolutionproductcategorysuggestAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpresssolutionproductcategorysuggestAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.product.category.suggest"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpresssolutionproductcategorysuggestAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AliexpresssolutionproductcategorysuggestAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetTitle is Title Setter
// product title
func (r *AliexpresssolutionproductcategorysuggestAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r AliexpresssolutionproductcategorysuggestAPIRequest) GetTitle() string {
	return r._title
}

// SetLanguage is Language Setter
// language
func (r *AliexpresssolutionproductcategorysuggestAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// GetLanguage Language Getter
func (r AliexpresssolutionproductcategorysuggestAPIRequest) GetLanguage() string {
	return r._language
}

// SetImageUrl is ImageUrl Setter
// product main image
func (r *AliexpresssolutionproductcategorysuggestAPIRequest) SetImageUrl(_imageUrl string) error {
	r._imageUrl = _imageUrl
	r.Set("image_url", _imageUrl)
	return nil
}

// GetImageUrl ImageUrl Getter
func (r AliexpresssolutionproductcategorysuggestAPIRequest) GetImageUrl() string {
	return r._imageUrl
}
