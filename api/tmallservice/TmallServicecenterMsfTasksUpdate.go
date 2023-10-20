package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterMsfTasksUpdate 喵师傅工人任务批量完成接口
// tmall.servicecenter.msf.tasks.update
//
// 喵师傅工人任务批量完成接口
func TmallServicecenterMsfTasksUpdate(clt *core.SDKClient, req *tmallservice.TmallServicecenterMsfTasksUpdateAPIRequest, resp *tmallservice.TmallServicecenterMsfTasksUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
