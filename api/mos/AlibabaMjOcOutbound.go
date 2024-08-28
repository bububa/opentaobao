package mos

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMjOcOutbound 零售商品发货
// alibaba.mj.oc.outbound
//
// 用于接收发货的数据
func AlibabaMjOcOutbound(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMjOcOutboundAPIRequest, resp *mos.AlibabaMjOcOutboundAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
