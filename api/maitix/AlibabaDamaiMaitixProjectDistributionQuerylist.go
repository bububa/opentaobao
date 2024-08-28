package maitix

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// AlibabaDamaiMaitixProjectDistributionQuerylist 分销项目列表查询（已过时，不推荐使用）
// alibaba.damai.maitix.project.distribution.querylist
//
// 分销项目列表查询接口（已过时，不推荐使用）
func AlibabaDamaiMaitixProjectDistributionQuerylist(ctx context.Context, clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixProjectDistributionQuerylistAPIRequest, resp *maitix.AlibabaDamaiMaitixProjectDistributionQuerylistAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
