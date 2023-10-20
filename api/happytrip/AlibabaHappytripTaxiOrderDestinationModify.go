package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// Alibabahappytriptaxiorderdestinationmodify 修改目的地
// alibaba.happytrip.taxi.order.destination.modify
//
// 通知ISV修改订单信息
func Alibabahappytriptaxiorderdestinationmodify(clt *core.SDKClient, req *happytrip.AlibabahappytriptaxiorderdestinationmodifyAPIRequest, session string) (*happytrip.AlibabahappytriptaxiorderdestinationmodifyAPIResponse, error) {
	var resp happytrip.AlibabahappytriptaxiorderdestinationmodifyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
