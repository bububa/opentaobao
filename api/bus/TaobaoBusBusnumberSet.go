package bus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusBusnumberSet 商家汽车票车次更新通知接口
// taobao.bus.busnumber.set
//
// 商家汽车票车次更新后，调用该接口通知平台。
func TaobaoBusBusnumberSet(ctx context.Context, clt *core.SDKClient, req *bus.TaobaoBusBusnumberSetAPIRequest, resp *bus.TaobaoBusBusnumberSetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
