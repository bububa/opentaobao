package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusBusnumberGet 汽车票车次查询
// taobao.bus.busnumber.get
//
// 提供汽车票车次查询服务
func TaobaoBusBusnumberGet(clt *core.SDKClient, req *bus.TaobaoBusBusnumberGetAPIRequest, resp *bus.TaobaoBusBusnumberGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
