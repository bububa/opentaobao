package pur

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/pur"
)

// AlibabaPurCmallGoodsSync 第三方商家接入采购商城-商品同步
// alibaba.pur.cmall.goods.sync
//
// 第三方商家接入采购商城-商品同步
func AlibabaPurCmallGoodsSync(ctx context.Context, clt *core.SDKClient, req *pur.AlibabaPurCmallGoodsSyncAPIRequest, resp *pur.AlibabaPurCmallGoodsSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
