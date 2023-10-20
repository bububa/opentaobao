package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobuscityget 城市接口
// taobao.bus.city.get
//
// 汽车票出发城市获取接口，获取所有出发城市
func Taobaobuscityget(clt *core.SDKClient, req *bus.TaobaobuscitygetAPIRequest, session string) (*bus.TaobaobuscitygetAPIResponse, error) {
	var resp bus.TaobaobuscitygetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
