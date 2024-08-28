package mos

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMjOcOfflineMaxticketnoGet pos机获取线下最大小票号
// alibaba.mj.oc.offline.maxticketno.get
//
// 给pos机提供线下最大小票号查询
func AlibabaMjOcOfflineMaxticketnoGet(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMjOcOfflineMaxticketnoGetAPIRequest, resp *mos.AlibabaMjOcOfflineMaxticketnoGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
