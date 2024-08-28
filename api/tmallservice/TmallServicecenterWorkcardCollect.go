package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardCollect 工单揽件
// tmall.servicecenter.workcard.collect
//
// 服务工单揽件接口
func TmallServicecenterWorkcardCollect(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardCollectAPIRequest, resp *tmallservice.TmallServicecenterWorkcardCollectAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
