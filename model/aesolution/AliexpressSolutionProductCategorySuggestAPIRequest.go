package aesolution

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionProductCategorySuggestAPIRequest Suggest product categories API请求
// aliexpress.solution.product.category.suggest
//
// Suggest product categories by title and image.
type AliexpressSolutionProductCategorySuggestAPIRequest struct {
	model.Params
	// product title
	_title string
	// language
	_language string
	// product main image
	_imageUrl string
}

// NewAliexpressSolutionProductCategorySuggestRequest 初始化AliexpressSolutionProductCategorySuggestAPIRequest对象
func NewAliexpressSolutionProductCategorySuggestRequest() *AliexpressSolutionProductCategorySuggestAPIRequest {
	return &AliexpressSolutionProductCategorySuggestAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AliexpressSolutionProductCategorySuggestAPIRequest) GetApiMethodName() string {
	return "aliexpress.solution.product.category.suggest"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AliexpressSolutionProductCategorySuggestAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetTitle is Title Setter
// product title
func (r *AliexpressSolutionProductCategorySuggestAPIRequest) SetTitle(_title string) error {
	r._title = _title
	r.Set("title", _title)
	return nil
}

// GetTitle Title Getter
func (r AliexpressSolutionProductCategorySuggestAPIRequest) GetTitle() string {
	return r._title
}

// SetLanguage is Language Setter
// language
func (r *AliexpressSolutionProductCategorySuggestAPIRequest) SetLanguage(_language string) error {
	r._language = _language
	r.Set("language", _language)
	return nil
}

// GetLanguage Language Getter
func (r AliexpressSolutionProductCategorySuggestAPIRequest) GetLanguage() string {
	return r._language
}

// SetImageUrl is ImageUrl Setter
// product main image
func (r *AliexpressSolutionProductCategorySuggestAPIRequest) SetImageUrl(_imageUrl string) error {
	r._imageUrl = _imageUrl
	r.Set("image_url", _imageUrl)
	return nil
}

// GetImageUrl ImageUrl Getter
func (r AliexpressSolutionProductCategorySuggestAPIRequest) GetImageUrl() string {
	return r._imageUrl
}
