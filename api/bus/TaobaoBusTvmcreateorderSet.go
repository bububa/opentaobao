package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusTvmcreateorderSet 线下自助机创建订单
// taobao.bus.tvmcreateorder.set
//
// 提供给汽车票线下自助机的创建订单使用
func TaobaoBusTvmcreateorderSet(clt *core.SDKClient, req *bus.TaobaoBusTvmcreateorderSetAPIRequest, resp *bus.TaobaoBusTvmcreateorderSetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
