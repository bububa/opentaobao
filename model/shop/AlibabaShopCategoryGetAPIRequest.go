package shop

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaShopCategoryGetAPIRequest 指定店铺分类信息查询接口 API请求
// alibaba.shop.category.get
//
// 按照卖家身份查询指定分类信息
type AlibabaShopCategoryGetAPIRequest struct {
	model.Params
	// 分类id
	_categoryId int64
}

// NewAlibabaShopCategoryGetRequest 初始化AlibabaShopCategoryGetAPIRequest对象
func NewAlibabaShopCategoryGetRequest() *AlibabaShopCategoryGetAPIRequest {
	return &AlibabaShopCategoryGetAPIRequest{
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaShopCategoryGetAPIRequest) Reset() {
	r._categoryId = 0
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaShopCategoryGetAPIRequest) GetApiMethodName() string {
	return "alibaba.shop.category.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaShopCategoryGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaShopCategoryGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetCategoryId is CategoryId Setter
// 分类id
func (r *AlibabaShopCategoryGetAPIRequest) SetCategoryId(_categoryId int64) error {
	r._categoryId = _categoryId
	r.Set("category_id", _categoryId)
	return nil
}

// GetCategoryId CategoryId Getter
func (r AlibabaShopCategoryGetAPIRequest) GetCategoryId() int64 {
	return r._categoryId
}

var poolAlibabaShopCategoryGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaShopCategoryGetRequest()
	},
}

// GetAlibabaShopCategoryGetRequest 从 sync.Pool 获取 AlibabaShopCategoryGetAPIRequest
func GetAlibabaShopCategoryGetAPIRequest() *AlibabaShopCategoryGetAPIRequest {
	return poolAlibabaShopCategoryGetAPIRequest.Get().(*AlibabaShopCategoryGetAPIRequest)
}

// ReleaseAlibabaShopCategoryGetAPIRequest 将 AlibabaShopCategoryGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaShopCategoryGetAPIRequest(v *AlibabaShopCategoryGetAPIRequest) {
	v.Reset()
	poolAlibabaShopCategoryGetAPIRequest.Put(v)
}
