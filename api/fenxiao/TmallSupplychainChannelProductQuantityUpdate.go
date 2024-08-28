package fenxiao

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/fenxiao"
)

// TmallSupplychainChannelProductQuantityUpdate 渠道无仓库存更新接口
// tmall.supplychain.channel.product.quantity.update
//
// 渠道无仓库存更新接口
func TmallSupplychainChannelProductQuantityUpdate(ctx context.Context, clt *core.SDKClient, req *fenxiao.TmallSupplychainChannelProductQuantityUpdateAPIRequest, resp *fenxiao.TmallSupplychainChannelProductQuantityUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
