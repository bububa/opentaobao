package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoInventoryInitial 库存初始化
// taobao.inventory.initial
//
// 建议使用新接口：taobao.inventory.merchant.adjust ，该接口会逐步停用。
// 商家仓库存初始化接口，直接按照商家指定的商品库存数进行填充，没有单据核对，不参与库存对账。
func TaobaoInventoryInitial(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoInventoryInitialAPIRequest, resp *fenxiao.TaobaoInventoryInitialAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
