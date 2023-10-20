package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// Taobaobuscancleorderset 取消订单
// taobao.bus.cancleorder.set
//
// 取消订单
func Taobaobuscancleorderset(clt *core.SDKClient, req *bus.TaobaobuscancleordersetAPIRequest, session string) (*bus.TaobaobuscancleordersetAPIResponse, error) {
	var resp bus.TaobaobuscancleordersetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
