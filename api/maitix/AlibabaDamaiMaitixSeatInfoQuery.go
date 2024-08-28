package maitix

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/maitix"
)

// AlibabaDamaiMaitixSeatInfoQuery 分销商查询座位信息
// alibaba.damai.maitix.seat.info.query
//
// 分销查询座位文案信息
func AlibabaDamaiMaitixSeatInfoQuery(ctx context.Context, clt *core.SDKClient, req *maitix.AlibabaDamaiMaitixSeatInfoQueryAPIRequest, resp *maitix.AlibabaDamaiMaitixSeatInfoQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
