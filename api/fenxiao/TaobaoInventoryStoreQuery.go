package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TaobaoInventoryStoreQuery 查询仓库信息
// taobao.inventory.store.query
//
// 查询商家仓信息
func TaobaoInventoryStoreQuery(ctx context.Context, clt *core.SDKClient, req *fenxiao.TaobaoInventoryStoreQueryAPIRequest, resp *fenxiao.TaobaoInventoryStoreQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
