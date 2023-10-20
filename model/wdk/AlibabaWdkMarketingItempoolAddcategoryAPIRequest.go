package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingItempoolAddcategoryAPIRequest 增加商品池里面的类目 API请求
// alibaba.wdk.marketing.itempool.addcategory
//
// 增加商品池里面的类目
type AlibabaWdkMarketingItempoolAddcategoryAPIRequest struct {
	model.Params
	// 类目对象
	_itemPoolActivityCategory *ItemPoolActivityCategory
	// 活动对象
	_commonActivityParam *CommonActivityParam
}

// NewAlibabaWdkMarketingItempoolAddcategoryRequest 初始化AlibabaWdkMarketingItempoolAddcategoryAPIRequest对象
func NewAlibabaWdkMarketingItempoolAddcategoryRequest() *AlibabaWdkMarketingItempoolAddcategoryAPIRequest {
	return &AlibabaWdkMarketingItempoolAddcategoryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaWdkMarketingItempoolAddcategoryAPIRequest) Reset() {
	r._itemPoolActivityCategory = nil
	r._commonActivityParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaWdkMarketingItempoolAddcategoryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itempool.addcategory"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaWdkMarketingItempoolAddcategoryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaWdkMarketingItempoolAddcategoryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemPoolActivityCategory is ItemPoolActivityCategory Setter
// 类目对象
func (r *AlibabaWdkMarketingItempoolAddcategoryAPIRequest) SetItemPoolActivityCategory(_itemPoolActivityCategory *ItemPoolActivityCategory) error {
	r._itemPoolActivityCategory = _itemPoolActivityCategory
	r.Set("item_pool_activity_category", _itemPoolActivityCategory)
	return nil
}

// GetItemPoolActivityCategory ItemPoolActivityCategory Getter
func (r AlibabaWdkMarketingItempoolAddcategoryAPIRequest) GetItemPoolActivityCategory() *ItemPoolActivityCategory {
	return r._itemPoolActivityCategory
}

// SetCommonActivityParam is CommonActivityParam Setter
// 活动对象
func (r *AlibabaWdkMarketingItempoolAddcategoryAPIRequest) SetCommonActivityParam(_commonActivityParam *CommonActivityParam) error {
	r._commonActivityParam = _commonActivityParam
	r.Set("common_activity_param", _commonActivityParam)
	return nil
}

// GetCommonActivityParam CommonActivityParam Getter
func (r AlibabaWdkMarketingItempoolAddcategoryAPIRequest) GetCommonActivityParam() *CommonActivityParam {
	return r._commonActivityParam
}

var poolAlibabaWdkMarketingItempoolAddcategoryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaWdkMarketingItempoolAddcategoryRequest()
	},
}

// GetAlibabaWdkMarketingItempoolAddcategoryRequest 从 sync.Pool 获取 AlibabaWdkMarketingItempoolAddcategoryAPIRequest
func GetAlibabaWdkMarketingItempoolAddcategoryAPIRequest() *AlibabaWdkMarketingItempoolAddcategoryAPIRequest {
	return poolAlibabaWdkMarketingItempoolAddcategoryAPIRequest.Get().(*AlibabaWdkMarketingItempoolAddcategoryAPIRequest)
}

// ReleaseAlibabaWdkMarketingItempoolAddcategoryAPIRequest 将 AlibabaWdkMarketingItempoolAddcategoryAPIRequest 放入 sync.Pool
func ReleaseAlibabaWdkMarketingItempoolAddcategoryAPIRequest(v *AlibabaWdkMarketingItempoolAddcategoryAPIRequest) {
	v.Reset()
	poolAlibabaWdkMarketingItempoolAddcategoryAPIRequest.Put(v)
}
