package bus

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusTvmbookorderSet 线下自助机通知出票接口
// taobao.bus.tvmbookorder.set
//
// 出票，当成功的时候告知出票；当失败的时候告知出票失败，飞猪退款给用户。
func TaobaoBusTvmbookorderSet(clt *core.SDKClient, req *bus.TaobaoBusTvmbookorderSetAPIRequest, resp *bus.TaobaoBusTvmbookorderSetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
