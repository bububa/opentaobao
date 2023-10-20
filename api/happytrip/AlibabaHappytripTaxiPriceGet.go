package happytrip

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/happytrip"
)

// Alibabahappytriptaxipriceget 获取价格预估信息
// alibaba.happytrip.taxi.price.get
//
// 打车价格预估
func Alibabahappytriptaxipriceget(clt *core.SDKClient, req *happytrip.AlibabahappytriptaxipricegetAPIRequest, session string) (*happytrip.AlibabahappytriptaxipricegetAPIResponse, error) {
	var resp happytrip.AlibabahappytriptaxipricegetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
