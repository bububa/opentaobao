package shop

import (
	"net/url"

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
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r AlibabaShopCategoryAllGetAPIRequest) GetApiMethodName() string {
	return "alibaba.shop.category.all.get"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r AlibabaShopCategoryAllGetAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}
