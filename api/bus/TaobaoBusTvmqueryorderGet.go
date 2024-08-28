package bus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusTvmqueryorderGet 线下自助机查询订单信息
// taobao.bus.tvmqueryorder.get
//
// 查询订单详情
func TaobaoBusTvmqueryorderGet(ctx context.Context, clt *core.SDKClient, req *bus.TaobaoBusTvmqueryorderGetAPIRequest, resp *bus.TaobaoBusTvmqueryorderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
