package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkItemMerchantCategoryQueryAPIRequest
查询商品的商家叶子类目 API请求
alibaba.wdk.item.merchant.category.query

查询商品的商家叶子类目 */
type AlibabaWdkItemMerchantCategoryQueryAPIRequest struct {
	model.Params
	// 请求
	_queryRequest *WdkOpenSkuMerchantCatServiceQueryRequest
}

// New
