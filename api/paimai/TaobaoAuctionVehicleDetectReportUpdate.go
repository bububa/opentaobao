package paimai

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/paimai"
)

// Taobaoauctionvehicledetectreportupdate 检测服务-服务单报告信息更新
// taobao.auction.vehicle.detect.report.update
//
// 检测服务-服务单报告信息更新
func Taobaoauctionvehicledetectreportupdate(clt *core.SDKClient, req *paimai.TaobaoauctionvehicledetectreportupdateAPIRequest, session string) (*paimai.TaobaoauctionvehicledetectreportupdateAPIResponse, error) {
	var resp paimai.TaobaoauctionvehicledetectreportupdateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
