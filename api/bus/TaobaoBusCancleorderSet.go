package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusCancleorderSet 取消订单
// taobao.bus.cancleorder.set
//
// 取消订单
func TaobaoBusCancleorderSet(clt *core.SDKClient, req *bus.TaobaoBusCancleorderSetAPIRequest, resp *bus.TaobaoBusCancleorderSetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
