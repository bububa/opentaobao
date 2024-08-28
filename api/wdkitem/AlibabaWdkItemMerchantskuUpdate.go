package wdkitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// AlibabaWdkItemMerchantskuUpdate 商家商品修改
// alibaba.wdk.item.merchantsku.update
//
// 商家商品修改
func AlibabaWdkItemMerchantskuUpdate(ctx context.Context, clt *core.SDKClient, req *wdkitem.AlibabaWdkItemMerchantskuUpdateAPIRequest, resp *wdkitem.AlibabaWdkItemMerchantskuUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
