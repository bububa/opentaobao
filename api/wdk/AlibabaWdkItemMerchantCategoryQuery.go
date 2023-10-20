package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkItemMerchantCategoryQuery 查询商品的商家叶子类目
// alibaba.wdk.item.merchant.category.query
//
// 查询商品的商家叶子类目
func AlibabaWdkItemMerchantCategoryQuery(clt *core.SDKClient, req *wdk.AlibabaWdkItemMerchantCategoryQueryAPIRequest, resp *wdk.AlibabaWdkItemMerchantCategoryQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
