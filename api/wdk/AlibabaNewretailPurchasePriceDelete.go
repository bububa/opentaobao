package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaNewretailPurchasePriceDelete 共享库存 商户删除采购价
// alibaba.newretail.purchase.price.delete
//
// 共享库存 商户删除采购价
func AlibabaNewretailPurchasePriceDelete(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaNewretailPurchasePriceDeleteAPIRequest, resp *wdk.AlibabaNewretailPurchasePriceDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
