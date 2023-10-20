package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallproductmatchschemagetAPIRequest 获取匹配产品规则 API请求
// tmall.product.match.schema.get
//
// ISV发布商品前，需要先查找到产品ID，这个接口返回查找产品规则入参规则
type TmallproductmatchschemagetAPIRequest struct {
	model.Params
	// 商品发布的目标类目，必须是叶子类目
	_categoryId int64
}

// NewTmallproductmatchschemagetRequest 初始化TmallproductmatchschemagetAPIRequest对象
func NewTmallproductmatchschemagetRequest() *TmallproductmatchschemagetAPIRequest {
	return &TmallproductmatchschemagetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallproductmatchschemagetAPIRequest) GetApiMethodName() string {
	return "tmall.product.match.schema.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallproductmatchschemagetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallproductmatchschemagetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCategoryId is CategoryId Setter
// 商品发布的目标类目，必须是叶子类目
func (r *TmallproductmatchschemagetAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r TmallproductmatchschemagetAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}
