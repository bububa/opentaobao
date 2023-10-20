package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// Alibabahappytriptaxiordercreate 用户叫车
// alibaba.happytrip.taxi.order.create
//
// 用户根据需要发起叫车请求，在发起请求之前必须事先获得order id.
func Alibabahappytriptaxiordercreate(clt *core.SDKClient, req *happytrip.AlibabahappytriptaxiordercreateAPIRequest, session string) (*happytrip.AlibabahappytriptaxiordercreateAPIResponse, error) {
	var resp happytrip.AlibabahappytriptaxiordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
