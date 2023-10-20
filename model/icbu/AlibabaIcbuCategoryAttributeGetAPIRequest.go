package icbu

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaIcbuCategoryAttributeGetAPIRequest 类目属性获取 API请求
// alibaba.icbu.category.attribute.get
//
// 根据类目ID获取系统定义的属性
type AlibabaIcbuCategoryAttributeGetAPIRequest struct {
	model.Params
	// 发布类目id
	_catId int64
}

// NewAlibabaIcbuCategoryAttributeGetRequest 初始化AlibabaIcbuCategoryAttributeGetAPIRequest对象
func NewAlibabaIcbuCategoryAttributeGetRequest() *AlibabaIcbuCategoryAttributeGetAPIRequest {
	return &AlibabaIcbuCategoryAttributeGetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaIcbuCategoryAttributeGetAPIRequest) GetApiMethodName() string {
	return "alibaba.icbu.category.attribute.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaIcbuCategoryAttributeGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaIcbuCategoryAttributeGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCatId is CatId Setter
// 发布类目id
func (r *AlibabaIcbuCategoryAttributeGetAPIRequest) SetCatId(_catId int64) error {
	r._catId = _catId
	r.Set("cat_id", _catId)
	return nil
}

// GetCatId CatId Getter
func (r AlibabaIcbuCategoryAttributeGetAPIRequest) GetCatId() int64 {
	return r._catId
}
