package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabaJymItemPropertyDefQueryAPIRequest 交易猫商品属性定义查询 API请求
// alibaba.jym.item.property.def.query
//
// 查询商品发布属性定义详情
type AlibabaJymItemPropertyDefQueryAPIRequest struct {
	model.Params
	// 二级类目ID
	_categoryId int64
}

// NewAlibabaJymItemPropertyDefQueryRequest 初始化AlibabaJymItemPropertyDefQueryAPIRequest对象
func NewAlibabaJymItemPropertyDefQueryRequest() *AlibabaJymItemPropertyDefQueryAPIRequest {
	return &AlibabaJymItemPropertyDefQueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaJymItemPropertyDefQueryAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.item.property.def.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaJymItemPropertyDefQueryAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// SetCategoryId is CategoryId Setter
// 二级类目ID
func (r *AlibabaJymItemPropertyDefQueryAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r AlibabaJymItemPropertyDefQueryAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}
