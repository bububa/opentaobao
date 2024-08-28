package retail

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/retail"
)

// AlibabaRetailVendingPriceWhitelistRemove 价格管控白名单去除
// alibaba.retail.vending.price.whitelist.remove
//
// 商家价格管控白名单去除
func AlibabaRetailVendingPriceWhitelistRemove(ctx context.Context, clt *core.SDKClient, req *retail.AlibabaRetailVendingPriceWhitelistRemoveAPIRequest, resp *retail.AlibabaRetailVendingPriceWhitelistRemoveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
