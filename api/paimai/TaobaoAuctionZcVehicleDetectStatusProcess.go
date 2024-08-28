package paimai

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/paimai"
)

// TaobaoAuctionZcVehicleDetectStatusProcess 检测服务-服务单状态流转
// taobao.auction.zc.vehicle.detect.status.process
//
// 检测服务-服务单状态流转
func TaobaoAuctionZcVehicleDetectStatusProcess(ctx context.Context, clt *core.SDKClient, req *paimai.TaobaoAuctionZcVehicleDetectStatusProcessAPIRequest, resp *paimai.TaobaoAuctionZcVehicleDetectStatusProcessAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
