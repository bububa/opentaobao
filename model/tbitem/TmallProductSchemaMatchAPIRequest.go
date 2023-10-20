package tbitem

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// TmallproductschemamatchAPIRequest product匹配接口 API请求
// tmall.product.schema.match
//
// 根据tmall.product.match.schema.get获取到的规则，填充相应地的字段值以及类目，匹配符合条件的产品，返回匹配product结果，注意，有可能返回多个产品ID，以逗号分隔（尤其是图书类目）；
type TmallproductschemamatchAPIRequest struct {
	model.Params
	// 根据tmall.product.match.schema.get获取到的模板，ISV将需要的字段填充好相应的值结果XML。
	_propvalues string
	// 商品发布的目标类目，必须是叶子类目
	_categoryId int64
}

// NewTmallproductschemamatchRequest 初始化TmallproductschemamatchAPIRequest对象
func NewTmallproductschemamatchRequest() *TmallproductschemamatchAPIRequest {
	return &TmallproductschemamatchAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r TmallproductschemamatchAPIRequest) GetApiMethodName() string {
	return "tmall.product.schema.match"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r TmallproductschemamatchAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r TmallproductschemamatchAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetPropvalues is Propvalues Setter
// 根据tmall.product.match.schema.get获取到的模板，ISV将需要的字段填充好相应的值结果XML。
func (r *TmallproductschemamatchAPIRequest) SetPropvalues(_propvalues string) error {
	r._propvalues = _propvalues
	r.Set("propvalues", _propvalues)
	return nil
}

// GetPropvalues Propvalues Getter
func (r TmallproductschemamatchAPIRequest) GetPropvalues() string {
	return r._propvalues
}

// SetCategoryId is CategoryId Setter
// 商品发布的目标类目，必须是叶子类目
func (r *TmallproductschemamatchAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r TmallproductschemamatchAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}
