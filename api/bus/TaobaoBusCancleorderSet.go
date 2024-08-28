package bus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusCancleorderSet 取消订单
// taobao.bus.cancleorder.set
//
// 取消订单
func TaobaoBusCancleorderSet(ctx context.Context, clt *core.SDKClient, req *bus.TaobaoBusCancleorderSetAPIRequest, resp *bus.TaobaoBusCancleorderSetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
