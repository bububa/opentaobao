package shop

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaShopCategoryAllGetAPIRequest 全部店铺分类信息查询接口 API请求
// alibaba.shop.category.all.get
//
// 按照卖家身份查询全部分类信息
type AlibabaShopCategoryAllGetAPIRequest struct {
	model.Params
}

// NewAlibabaShopCategoryAllGetRequest 初始化AlibabaShopCategoryAllGetAPIRequest对象
func NewAlibabaShopCategoryAllGetRequest() *AlibabaShopCategoryAllGetAPIRequest {
	return &AlibabaShopCategoryAllGetAPIRequest{
		Params: model.NewParams(0),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaShopCategoryAllGetAPIRequest) Reset() {
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaShopCategoryAllGetAPIRequest) GetApiMethodName() string {
	return "alibaba.shop.category.all.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaShopCategoryAllGetAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaShopCategoryAllGetAPIRequest) GetRawParams() model.Params {
	return r.Params
}

var poolAlibabaShopCategoryAllGetAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaShopCategoryAllGetRequest()
	},
}

// GetAlibabaShopCategoryAllGetRequest 从 sync.Pool 获取 AlibabaShopCategoryAllGetAPIRequest
func GetAlibabaShopCategoryAllGetAPIRequest() *AlibabaShopCategoryAllGetAPIRequest {
	return poolAlibabaShopCategoryAllGetAPIRequest.Get().(*AlibabaShopCategoryAllGetAPIRequest)
}

// ReleaseAlibabaShopCategoryAllGetAPIRequest 将 AlibabaShopCategoryAllGetAPIRequest 放入 sync.Pool
func ReleaseAlibabaShopCategoryAllGetAPIRequest(v *AlibabaShopCategoryAllGetAPIRequest) {
	v.Reset()
	poolAlibabaShopCategoryAllGetAPIRequest.Put(v)
}
