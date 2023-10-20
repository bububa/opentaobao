package paimai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/paimai"
)

// TaobaoAuctionVehicleDetectReportUpdate 检测服务-服务单报告信息更新
// taobao.auction.vehicle.detect.report.update
//
// 检测服务-服务单报告信息更新
func TaobaoAuctionVehicleDetectReportUpdate(clt *core.SDKClient, req *paimai.TaobaoAuctionVehicleDetectReportUpdateAPIRequest, session string) (*paimai.TaobaoAuctionVehicleDetectReportUpdateAPIResponse, error) {
	var resp paimai.TaobaoAuctionVehicleDetectReportUpdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
