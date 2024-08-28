package paimai

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/paimai"
)

// TaobaoAuctionVehicleDetectReportUpdate 检测服务-服务单报告信息更新
// taobao.auction.vehicle.detect.report.update
//
// 检测服务-服务单报告信息更新
func TaobaoAuctionVehicleDetectReportUpdate(ctx context.Context, clt *core.SDKClient, req *paimai.TaobaoAuctionVehicleDetectReportUpdateAPIRequest, resp *paimai.TaobaoAuctionVehicleDetectReportUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
