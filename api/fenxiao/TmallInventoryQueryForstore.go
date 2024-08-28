package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TmallInventoryQueryForstore 查询后端商品仓库库存
// tmall.inventory.query.forstore
//
// 商家查询后端商品仓库库存
func TmallInventoryQueryForstore(ctx context.Context, clt *core.SDKClient, req *fenxiao.TmallInventoryQueryForstoreAPIRequest, resp *fenxiao.TmallInventoryQueryForstoreAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
