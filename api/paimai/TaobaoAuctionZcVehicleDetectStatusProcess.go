package paimai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/paimai"
)

// TaobaoAuctionZcVehicleDetectStatusProcess 检测服务-服务单状态流转
// taobao.auction.zc.vehicle.detect.status.process
//
// 检测服务-服务单状态流转
func TaobaoAuctionZcVehicleDetectStatusProcess(clt *core.SDKClient, req *paimai.TaobaoAuctionZcVehicleDetectStatusProcessAPIRequest, resp *paimai.TaobaoAuctionZcVehicleDetectStatusProcessAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
