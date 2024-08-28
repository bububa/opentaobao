package bus

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/bus"
)

// TaobaoBusRefundSet B2B退票申请接口
// taobao.bus.refund.set
//
// B2B业务支持退票
func TaobaoBusRefundSet(ctx context.Context, clt *core.SDKClient, req *bus.TaobaoBusRefundSetAPIRequest, resp *bus.TaobaoBusRefundSetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
