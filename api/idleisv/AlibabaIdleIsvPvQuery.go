package idleisv

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/idleisv"
)

// AlibabaIdleIsvPvQuery 查询pv属性
// alibaba.idle.isv.pv.query
//
// 查询pv属性
func AlibabaIdleIsvPvQuery(ctx context.Context, clt *core.SDKClient, req *idleisv.AlibabaIdleIsvPvQueryAPIRequest, resp *idleisv.AlibabaIdleIsvPvQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
