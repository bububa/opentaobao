package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallProductAddSchemaGetAPIRequest 产品发布规则获取接口 API请求
// tmall.product.add.schema.get
//
// 获取用户发布产品的规则
type TmallProductAddSchemaGetAPIRequest struct {
	model.Params
	// 商品发布的目标类目，必须是叶子类目
	_categoryId int64
	// 品牌ID
	_brandId int64
}

// NewTmallProductAddSchemaGetRequest 初始化TmallProductAddSchemaGetAPIRequest对象
func NewTmallProductAddSchemaGetRequest() *TmallProductAddSchemaGetAPIRequest {
	return &TmallProductAddSchemaGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallProductAddSchemaGetAPIRequest) GetApiMethodName() string {
	return "tmall.product.add.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallProductAddSchemaGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallProductAddSchemaGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCategoryId is CategoryId Setter
// 商品发布的目标类目，必须是叶子类目
func (r *TmallProductAddSchemaGetAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r TmallProductAddSchemaGetAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}

// SetBrandId is BrandId Setter
// 品牌ID
func (r *TmallProductAddSchemaGetAPIRequest) SetBrandId(_brandId int64) error {
	r._brandId = _brandId
	r.Set("brand_id", _brandId)
	return nil
}

// GetBrandId BrandId Getter
func (r TmallProductAddSchemaGetAPIRequest) GetBrandId() int64 {
	return r._brandId
}
