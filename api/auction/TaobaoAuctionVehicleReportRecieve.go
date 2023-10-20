package auction

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/auction"
)

// TaobaoAuctionVehicleReportRecieve 机动车报告回调数据接收
// taobao.auction.vehicle.report.recieve
//
// 机动车报告同步接收接口
func TaobaoAuctionVehicleReportRecieve(clt *core.SDKClient, req *auction.TaobaoAuctionVehicleReportRecieveAPIRequest, resp *auction.TaobaoAuctionVehicleReportRecieveAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
