package maitix

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// AlibabaDamaiMaitixProjectDistributionQuerybypage 分销项目分页查询项目列表服务
// alibaba.damai.maitix.project.distribution.querybypage
//
// 分销项目分页查询项目列表服务
func AlibabaDamaiMaitixProjectDistributionQuerybypage(ctx context.Context, clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixProjectDistributionQuerybypageAPIRequest, resp *maitix.AlibabaDamaiMaitixProjectDistributionQuerybypageAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
