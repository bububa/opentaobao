package wdkitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// AlibabaWdkItemMerchantskuQuery 商家商品信息查询
// alibaba.wdk.item.merchantsku.query
//
// 商家商品信息查询
func AlibabaWdkItemMerchantskuQuery(ctx context.Context, clt *core.SDKClient, req *wdkitem.AlibabaWdkItemMerchantskuQueryAPIRequest, resp *wdkitem.AlibabaWdkItemMerchantskuQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
