package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabawdkmarketingitempooladdcategoryAPIRequest 增加商品池里面的类目 API请求
// alibaba.wdk.marketing.itempool.addcategory
//
// 增加商品池里面的类目
type AlibabawdkmarketingitempooladdcategoryAPIRequest struct {
	model.Params
	// 类目对象
	_itemPoolActivityCategory *ItemPoolActivityCategory
	// 活动对象
	_commonActivityParam *CommonActivityParam
}

// NewAlibabawdkmarketingitempooladdcategoryRequest 初始化AlibabawdkmarketingitempooladdcategoryAPIRequest对象
func NewAlibabawdkmarketingitempooladdcategoryRequest() *AlibabawdkmarketingitempooladdcategoryAPIRequest {
	return &AlibabawdkmarketingitempooladdcategoryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabawdkmarketingitempooladdcategoryAPIRequest) GetApiMethodName() string {
	return "alibaba.wdk.marketing.itempool.addcategory"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabawdkmarketingitempooladdcategoryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabawdkmarketingitempooladdcategoryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemPoolActivityCategory is ItemPoolActivityCategory Setter
// 类目对象
func (r *AlibabawdkmarketingitempooladdcategoryAPIRequest) SetItemPoolActivityCategory(_itemPoolActivityCategory *ItemPoolActivityCategory) error {
	r._itemPoolActivityCategory = _itemPoolActivityCategory
	r.Set("item_pool_activity_category", _itemPoolActivityCategory)
	return nil
}

// GetItemPoolActivityCategory ItemPoolActivityCategory Getter
func (r AlibabawdkmarketingitempooladdcategoryAPIRequest) GetItemPoolActivityCategory() *ItemPoolActivityCategory {
	return r._itemPoolActivityCategory
}

// SetCommonActivityParam is CommonActivityParam Setter
// 活动对象
func (r *AlibabawdkmarketingitempooladdcategoryAPIRequest) SetCommonActivityParam(_commonActivityParam *CommonActivityParam) error {
	r._commonActivityParam = _commonActivityParam
	r.Set("common_activity_param", _commonActivityParam)
	return nil
}

// GetCommonActivityParam CommonActivityParam Getter
func (r AlibabawdkmarketingitempooladdcategoryAPIRequest) GetCommonActivityParam() *CommonActivityParam {
	return r._commonActivityParam
}
