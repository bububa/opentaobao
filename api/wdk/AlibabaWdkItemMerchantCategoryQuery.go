package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkItemMerchantCategoryQuery 查询商品的商家叶子类目
// alibaba.wdk.item.merchant.category.query
//
// 查询商品的商家叶子类目
func AlibabaWdkItemMerchantCategoryQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkItemMerchantCategoryQueryAPIRequest, resp *wdk.AlibabaWdkItemMerchantCategoryQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
