package bus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusHistoryorderGet 历史订单查询（对账）
// taobao.bus.historyorder.get
//
// 历史订单查询，对账接口
func TaobaoBusHistoryorderGet(ctx context.Context, clt *core.SDKClient, req *bus.TaobaoBusHistoryorderGetAPIRequest, resp *bus.TaobaoBusHistoryorderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
