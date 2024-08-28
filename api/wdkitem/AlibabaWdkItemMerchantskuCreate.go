package wdkitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// AlibabaWdkItemMerchantskuCreate 商家商品信息新建
// alibaba.wdk.item.merchantsku.create
//
// 商家商品信息新建
func AlibabaWdkItemMerchantskuCreate(ctx context.Context, clt *core.SDKClient, req *wdkitem.AlibabaWdkItemMerchantskuCreateAPIRequest, resp *wdkitem.AlibabaWdkItemMerchantskuCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
