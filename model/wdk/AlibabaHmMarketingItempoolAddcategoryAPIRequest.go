package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// AlibabahmmarketingitempooladdcategoryAPIRequest 增加商品池里面的类目 API请求
// alibaba.hm.marketing.itempool.addcategory
//
// 增加商品池里面的类目
type AlibabahmmarketingitempooladdcategoryAPIRequest struct {
	model.Params
	// 类目对象
	_itemPoolActivityCategory *ItemPoolActivityCategory
	// 活动对象
	_commonActivityParam *CommonActivityParam
}

// NewAlibabahmmarketingitempooladdcategoryRequest 初始化AlibabahmmarketingitempooladdcategoryAPIRequest对象
func NewAlibabahmmarketingitempooladdcategoryRequest() *AlibabahmmarketingitempooladdcategoryAPIRequest {
	return &AlibabahmmarketingitempooladdcategoryAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabahmmarketingitempooladdcategoryAPIRequest) GetApiMethodName() string {
	return "alibaba.hm.marketing.itempool.addcategory"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabahmmarketingitempooladdcategoryAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r AlibabahmmarketingitempooladdcategoryAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetItemPoolActivityCategory is ItemPoolActivityCategory Setter
// 类目对象
func (r *AlibabahmmarketingitempooladdcategoryAPIRequest) SetItemPoolActivityCategory(_itemPoolActivityCategory *ItemPoolActivityCategory) error {
	r._itemPoolActivityCategory = _itemPoolActivityCategory
	r.Set("item_pool_activity_category", _itemPoolActivityCategory)
	return nil
}

// GetItemPoolActivityCategory ItemPoolActivityCategory Getter
func (r AlibabahmmarketingitempooladdcategoryAPIRequest) GetItemPoolActivityCategory() *ItemPoolActivityCategory {
	return r._itemPoolActivityCategory
}

// SetCommonActivityParam is CommonActivityParam Setter
// 活动对象
func (r *AlibabahmmarketingitempooladdcategoryAPIRequest) SetCommonActivityParam(_commonActivityParam *CommonActivityParam) error {
	r._commonActivityParam = _commonActivityParam
	r.Set("common_activity_param", _commonActivityParam)
	return nil
}

// GetCommonActivityParam CommonActivityParam Getter
func (r AlibabahmmarketingitempooladdcategoryAPIRequest) GetCommonActivityParam() *CommonActivityParam {
	return r._commonActivityParam
}
