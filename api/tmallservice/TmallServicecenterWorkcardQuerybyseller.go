package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardQuerybyseller 工单查询接口（面向商家）
// tmall.servicecenter.workcard.querybyseller
//
// 查询工单
func TmallServicecenterWorkcardQuerybyseller(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardQuerybysellerAPIRequest, resp *tmallservice.TmallServicecenterWorkcardQuerybysellerAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
