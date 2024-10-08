package pur

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// AlibabaPurCmallGoodsStatusSync 第三方商城接入采购商城-商品状态同步
// alibaba.pur.cmall.goods.status.sync
//
// 第三方商城接入采购商城-商品状态同步
func AlibabaPurCmallGoodsStatusSync(ctx context.Context, clt *core.SDKClient, req *pur.AlibabaPurCmallGoodsStatusSyncAPIRequest, resp *pur.AlibabaPurCmallGoodsStatusSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
