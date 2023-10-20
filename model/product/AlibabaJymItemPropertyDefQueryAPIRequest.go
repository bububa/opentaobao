package product

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabajymitempropertydefqueryAPIRequest 交易猫商品属性定义查询 API请求
// alibaba.jym.item.property.def.query
//
// 查询商品发布属性定义详情
type AlibabajymitempropertydefqueryAPIRequest struct {
	model.Params
	// 二级类目ID
	_categoryId int64
}

// NewAlibabajymitempropertydefqueryRequest 初始化AlibabajymitempropertydefqueryAPIRequest对象
func NewAlibabajymitempropertydefqueryRequest() *AlibabajymitempropertydefqueryAPIRequest {
	return &AlibabajymitempropertydefqueryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabajymitempropertydefqueryAPIRequest) GetApiMethodName() string {
	return "alibaba.jym.item.property.def.query"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabajymitempropertydefqueryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabajymitempropertydefqueryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCategoryId is CategoryId Setter
// 二级类目ID
func (r *AlibabajymitempropertydefqueryAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r AlibabajymitempropertydefqueryAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}
