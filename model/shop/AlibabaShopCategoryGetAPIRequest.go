package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabashopcategorygetAPIRequest 指定店铺分类信息查询接口 API请求
// alibaba.shop.category.get
//
// 按照卖家身份查询指定分类信息
type AlibabashopcategorygetAPIRequest struct {
	model.Params
	// 分类id
	_categoryId int64
}

// NewAlibabashopcategorygetRequest 初始化AlibabashopcategorygetAPIRequest对象
func NewAlibabashopcategorygetRequest() *AlibabashopcategorygetAPIRequest {
	return &AlibabashopcategorygetAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabashopcategorygetAPIRequest) GetApiMethodName() string {
	return "alibaba.shop.category.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabashopcategorygetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabashopcategorygetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCategoryId is CategoryId Setter
// 分类id
func (r *AlibabashopcategorygetAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r AlibabashopcategorygetAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}
