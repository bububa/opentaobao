package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusNumbersUpdate 汽车票车次更新服务
// taobao.bus.numbers.update
//
// 用于汽车票车次信息的新增、更新和逻辑删除
func TaobaoBusNumbersUpdate(clt *core.SDKClient, req *bus.TaobaoBusNumbersUpdateAPIRequest, resp *bus.TaobaoBusNumbersUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
