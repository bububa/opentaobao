package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterTasksSearch 查询任务类工单信息
// tmall.servicecenter.tasks.search
//
// 查询任务类工单信息
func TmallServicecenterTasksSearch(clt *core.SDKClient, req *tmallservice.TmallServicecenterTasksSearchAPIRequest, resp *tmallservice.TmallServicecenterTasksSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
