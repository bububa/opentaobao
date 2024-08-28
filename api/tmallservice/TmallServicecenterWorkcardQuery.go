package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardQuery 工单查询接口
// tmall.servicecenter.workcard.query
//
// 工单查询接口
func TmallServicecenterWorkcardQuery(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardQueryAPIRequest, resp *tmallservice.TmallServicecenterWorkcardQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
