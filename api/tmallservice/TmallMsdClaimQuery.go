package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallMsdClaimQuery 查询待理赔工单数据接口
// tmall.msd.claim.query
//
// 查询待理赔工单数据接口
func TmallMsdClaimQuery(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallMsdClaimQueryAPIRequest, resp *tmallservice.TmallMsdClaimQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
