package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallProductMatchSchemaGetAPIRequest 获取匹配产品规则 API请求
// tmall.product.match.schema.get
//
// ISV发布商品前，需要先查找到产品ID，这个接口返回查找产品规则入参规则
type TmallProductMatchSchemaGetAPIRequest struct {
	model.Params
	// 商品发布的目标类目，必须是叶子类目
	_categoryId int64
}

// NewTmallProductMatchSchemaGetRequest 初始化TmallProductMatchSchemaGetAPIRequest对象
func NewTmallProductMatchSchemaGetRequest() *TmallProductMatchSchemaGetAPIRequest {
	return &TmallProductMatchSchemaGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallProductMatchSchemaGetAPIRequest) GetApiMethodName() string {
	return "tmall.product.match.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallProductMatchSchemaGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is CategoryId Setter
// 商品发布的目标类目，必须是叶子类目
func (r *TmallProductMatchSchemaGetAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// Get CategoryId Getter
func (r TmallProductMatchSchemaGetAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}
