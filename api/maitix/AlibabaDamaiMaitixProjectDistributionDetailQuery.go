package maitix

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// AlibabaDamaiMaitixProjectDistributionDetailQuery 大麦分销项目内容详情查询
// alibaba.damai.maitix.project.distribution.detail.query
//
// 大麦分销项目内容详情查询，提供项目的内容详情
func AlibabaDamaiMaitixProjectDistributionDetailQuery(ctx context.Context, clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixProjectDistributionDetailQueryAPIRequest, resp *maitix.AlibabaDamaiMaitixProjectDistributionDetailQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
