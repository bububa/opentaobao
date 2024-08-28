package bus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusLastplaceGet 获取目的地数据
// taobao.bus.lastplace.get
//
// 传入城市 获取对应的目的地
func TaobaoBusLastplaceGet(ctx context.Context, clt *core.SDKClient, req *bus.TaobaoBusLastplaceGetAPIRequest, resp *bus.TaobaoBusLastplaceGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
