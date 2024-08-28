package bus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusRefundfeeGet 查询退票费用明细
// taobao.bus.refundfee.get
//
// 查询退票的费用信息
func TaobaoBusRefundfeeGet(ctx context.Context, clt *core.SDKClient, req *bus.TaobaoBusRefundfeeGetAPIRequest, resp *bus.TaobaoBusRefundfeeGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
