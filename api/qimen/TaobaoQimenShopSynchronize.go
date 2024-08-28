package qimen

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/qimen"
)

// TaobaoQimenShopSynchronize 店铺同步接口
// taobao.qimen.shop.synchronize
//
// 店铺同步接口描述
func TaobaoQimenShopSynchronize(ctx context.Context, clt *core.SDKClient, req *qimen.TaobaoQimenShopSynchronizeAPIRequest, resp *qimen.TaobaoQimenShopSynchronizeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
