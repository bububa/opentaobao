package maitix

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// AlibabaDamaiMaitixEticketDistributionQuery 分销电子票查询接口
// alibaba.damai.maitix.eticket.distribution.query
//
// 分销电子票查询接口
func AlibabaDamaiMaitixEticketDistributionQuery(ctx context.Context, clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixEticketDistributionQueryAPIRequest, resp *maitix.AlibabaDamaiMaitixEticketDistributionQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
