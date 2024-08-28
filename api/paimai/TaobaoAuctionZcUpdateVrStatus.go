package paimai

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/paimai"
)

// TaobaoAuctionZcUpdateVrStatus 如视VR更新活跃状态
// taobao.auction.zc.update.vr.status
//
// 如视VR更新活跃状态
func TaobaoAuctionZcUpdateVrStatus(ctx context.Context, clt *core.SDKClient, req *paimai.TaobaoAuctionZcUpdateVrStatusAPIRequest, resp *paimai.TaobaoAuctionZcUpdateVrStatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
