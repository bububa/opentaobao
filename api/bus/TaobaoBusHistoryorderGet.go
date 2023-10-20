package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusHistoryorderGet 历史订单查询（对账）
// taobao.bus.historyorder.get
//
// 历史订单查询，对账接口
func TaobaoBusHistoryorderGet(clt *core.SDKClient, req *bus.TaobaoBusHistoryorderGetAPIRequest, resp *bus.TaobaoBusHistoryorderGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
