package openmall

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/openmall"
)

// TaobaoOpenmallTraceSearch 获取Openmall订单物流流转信息
// taobao.openmall.trace.search
//
// 获取Openmall订单物流流转信息
func TaobaoOpenmallTraceSearch(ctx context.Context, clt *core.SDKClient, req *openmall.TaobaoOpenmallTraceSearchAPIRequest, resp *openmall.TaobaoOpenmallTraceSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
