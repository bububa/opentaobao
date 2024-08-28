package lstwarehouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/lstwarehouse"
)

// AlibabaLstIcStockItemsUpdate 零售通经销商商品库存设置
// alibaba.lst.ic.stock.items.update
//
// 零售通经销商商品库存设置
func AlibabaLstIcStockItemsUpdate(ctx context.Context, clt *core.SDKClient, req *lstwarehouse.AlibabaLstIcStockItemsUpdateAPIRequest, resp *lstwarehouse.AlibabaLstIcStockItemsUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
