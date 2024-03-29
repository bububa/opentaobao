package btrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/btrip"
)

// AlitripBtripVehicleOrderSearch 用车订单查询接口
// alitrip.btrip.vehicle.order.search
//
// 企业获取商旅用车订单数据
func AlitripBtripVehicleOrderSearch(clt *core.SDKClient, req *btrip.AlitripBtripVehicleOrderSearchAPIRequest, resp *btrip.AlitripBtripVehicleOrderSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
