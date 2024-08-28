package icbuproduct

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/icbuproduct"
)

// AlibabaIcbuProductInventoryUpdate icbu商品库存更新
// alibaba.icbu.product.inventory.update
//
// 更新库存信息
func AlibabaIcbuProductInventoryUpdate(ctx context.Context, clt *core.SDKClient, req *icbuproduct.AlibabaIcbuProductInventoryUpdateAPIRequest, resp *icbuproduct.AlibabaIcbuProductInventoryUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
