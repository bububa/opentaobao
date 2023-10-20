package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// Alibabahappytriptaxiidget 获取请求id
// alibaba.happytrip.taxi.id.get
//
// 获取订单号
func Alibabahappytriptaxiidget(clt *core.SDKClient, req *happytrip.AlibabahappytriptaxiidgetAPIRequest, session string) (*happytrip.AlibabahappytriptaxiidgetAPIResponse, error) {
	var resp happytrip.AlibabahappytriptaxiidgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
