package maitix

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// AlibabaDamaiMaitixOrderQuery 大麦-查询分销单
// alibaba.damai.maitix.order.query
//
// 查询分销单
func AlibabaDamaiMaitixOrderQuery(ctx context.Context, clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixOrderQueryAPIRequest, resp *maitix.AlibabaDamaiMaitixOrderQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
