package auction

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/auction"
)

// Taobaoauctionvehiclereportrecieve 机动车报告回调数据接收
// taobao.auction.vehicle.report.recieve
//
// 机动车报告同步接收接口
func Taobaoauctionvehiclereportrecieve(clt *core.SDKClient, req *auction.TaobaoauctionvehiclereportrecieveAPIRequest, session string) (*auction.TaobaoauctionvehiclereportrecieveAPIResponse, error) {
	var resp auction.TaobaoauctionvehiclereportrecieveAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
