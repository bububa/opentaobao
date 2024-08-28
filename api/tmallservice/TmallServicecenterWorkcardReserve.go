package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardReserve 工单预约
// tmall.servicecenter.workcard.reserve
//
// 服务工单更新通用接口
func TmallServicecenterWorkcardReserve(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardReserveAPIRequest, resp *tmallservice.TmallServicecenterWorkcardReserveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
