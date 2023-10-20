package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusTicketSet 出票接口
// taobao.bus.ticket.set
//
// 提供给汽车票商家出票使用
func TaobaoBusTicketSet(clt *core.SDKClient, req *bus.TaobaoBusTicketSetAPIRequest, resp *bus.TaobaoBusTicketSetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
