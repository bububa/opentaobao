package paimai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/paimai"
)

// Taobaoauctionzcvehicledetectstatusprocess 检测服务-服务单状态流转
// taobao.auction.zc.vehicle.detect.status.process
//
// 检测服务-服务单状态流转
func Taobaoauctionzcvehicledetectstatusprocess(clt *core.SDKClient, req *paimai.TaobaoauctionzcvehicledetectstatusprocessAPIRequest, session string) (*paimai.TaobaoauctionzcvehicledetectstatusprocessAPIResponse, error) {
	var resp paimai.TaobaoauctionzcvehicledetectstatusprocessAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
