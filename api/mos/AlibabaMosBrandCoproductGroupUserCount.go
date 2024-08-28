package mos

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMosBrandCoproductGroupUserCount 按照查询条件统计总数
// alibaba.mos.brand.coproduct.group.user.count
//
// 按照查询条件统计总数
func AlibabaMosBrandCoproductGroupUserCount(ctx context.Context, clt *core.SDKClient, req *mos.AlibabaMosBrandCoproductGroupUserCountAPIRequest, resp *mos.AlibabaMosBrandCoproductGroupUserCountAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
