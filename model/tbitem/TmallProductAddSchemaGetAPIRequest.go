package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallproductaddschemagetAPIRequest 产品发布规则获取接口 API请求
// tmall.product.add.schema.get
//
// 获取用户发布产品的规则
type TmallproductaddschemagetAPIRequest struct {
	model.Params
	// 商品发布的目标类目，必须是叶子类目
	_categoryId int64
	// 品牌ID
	_brandId int64
}

// NewTmallproductaddschemagetRequest 初始化TmallproductaddschemagetAPIRequest对象
func NewTmallproductaddschemagetRequest() *TmallproductaddschemagetAPIRequest {
	return &TmallproductaddschemagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallproductaddschemagetAPIRequest) GetApiMethodName() string {
	return "tmall.product.add.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallproductaddschemagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallproductaddschemagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCategoryId is CategoryId Setter
// 商品发布的目标类目，必须是叶子类目
func (r *TmallproductaddschemagetAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r TmallproductaddschemagetAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}

// SetBrandId is BrandId Setter
// 品牌ID
func (r *TmallproductaddschemagetAPIRequest) SetBrandId(_brandId int64) error {
	r._brandId = _brandId
	r.Set("brand_id", _brandId)
	return nil
}

// GetBrandId BrandId Getter
func (r TmallproductaddschemagetAPIRequest) GetBrandId() int64 {
	return r._brandId
}
