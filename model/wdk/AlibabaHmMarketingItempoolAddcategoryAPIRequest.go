package wdk

import (
	"net/url"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingItempoolAddcategoryAPIRequest 增加商品池里面的类目 API请求
// alibaba.hm.marketing.itempool.addcategory
//
// 增加商品池里面的类目
type AlibabaHmMarketingItempoolAddcategoryAPIRequest struct {
	model.Params
	// 类目对象
	_itemPoolActivityCategory *ItemPoolActivityCategory
	// 活动对象
	_commonActivityParam *CommonActivityParam
}

// NewAlibabaHmMarketingItempoolAddcategoryRequest 初始化AlibabaHmMarketingItempoolAddcategoryAPIRequest对象
func NewAlibabaHmMarketingItempoolAddcategoryRequest() *AlibabaHmMarketingItempoolAddcategoryAPIRequest {
	return &AlibabaHmMarketingItempoolAddcategoryAPIRequest{
		Params: model.NewParams(2),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *AlibabaHmMarketingItempoolAddcategoryAPIRequest) Reset() {
	r._itemPoolActivityCategory = nil
	r._commonActivityParam = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaHmMarketingItempoolAddcategoryAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itempool.addcategory"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaHmMarketingItempoolAddcategoryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabaHmMarketingItempoolAddcategoryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemPoolActivityCategory is ItemPoolActivityCategory Setter
// 类目对象
func (r *AlibabaHmMarketingItempoolAddcategoryAPIRequest) SetItemPoolActivityCategory(_itemPoolActivityCategory *ItemPoolActivityCategory) error {
	r._itemPoolActivityCategory = _itemPoolActivityCategory
	r.Set("item_pool_activity_category", _itemPoolActivityCategory)
	return nil
}

// GetItemPoolActivityCategory ItemPoolActivityCategory Getter
func (r AlibabaHmMarketingItempoolAddcategoryAPIRequest) GetItemPoolActivityCategory() *ItemPoolActivityCategory {
	return r._itemPoolActivityCategory
}

// SetCommonActivityParam is CommonActivityParam Setter
// 活动对象
func (r *AlibabaHmMarketingItempoolAddcategoryAPIRequest) SetCommonActivityParam(_commonActivityParam *CommonActivityParam) error {
	r._commonActivityParam = _commonActivityParam
	r.Set("common_activity_param", _commonActivityParam)
	return nil
}

// GetCommonActivityParam CommonActivityParam Getter
func (r AlibabaHmMarketingItempoolAddcategoryAPIRequest) GetCommonActivityParam() *CommonActivityParam {
	return r._commonActivityParam
}

var poolAlibabaHmMarketingItempoolAddcategoryAPIRequest = sync.Pool{
	New: func() any {
		return NewAlibabaHmMarketingItempoolAddcategoryRequest()
	},
}

// GetAlibabaHmMarketingItempoolAddcategoryRequest 从 sync.Pool 获取 AlibabaHmMarketingItempoolAddcategoryAPIRequest
func GetAlibabaHmMarketingItempoolAddcategoryAPIRequest() *AlibabaHmMarketingItempoolAddcategoryAPIRequest {
	return poolAlibabaHmMarketingItempoolAddcategoryAPIRequest.Get().(*AlibabaHmMarketingItempoolAddcategoryAPIRequest)
}

// ReleaseAlibabaHmMarketingItempoolAddcategoryAPIRequest 将 AlibabaHmMarketingItempoolAddcategoryAPIRequest 放入 sync.Pool
func ReleaseAlibabaHmMarketingItempoolAddcategoryAPIRequest(v *AlibabaHmMarketingItempoolAddcategoryAPIRequest) {
	v.Reset()
	poolAlibabaHmMarketingItempoolAddcategoryAPIRequest.Put(v)
}
