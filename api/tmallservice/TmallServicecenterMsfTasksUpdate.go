package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterMsfTasksUpdate 喵师傅工人任务批量完成接口
// tmall.servicecenter.msf.tasks.update
//
// 喵师傅工人任务批量完成接口
func TmallServicecenterMsfTasksUpdate(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterMsfTasksUpdateAPIRequest, resp *tmallservice.TmallServicecenterMsfTasksUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
