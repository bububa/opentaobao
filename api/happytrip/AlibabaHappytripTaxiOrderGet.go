package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// Alibabahappytriptaxiorderget 订单详情
// alibaba.happytrip.taxi.order.get
//
// 获取订单状态及详情
func Alibabahappytriptaxiorderget(clt *core.SDKClient, req *happytrip.AlibabahappytriptaxiordergetAPIRequest, session string) (*happytrip.AlibabahappytriptaxiordergetAPIResponse, error) {
	var resp happytrip.AlibabahappytriptaxiordergetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
