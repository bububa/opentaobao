package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterTaskQueryrefund 查询任务类工单是否退款
// tmall.servicecenter.task.queryrefund
//
// 查询任务类工单是否退款
func TmallServicecenterTaskQueryrefund(clt *core.SDKClient, req *tmallservice.TmallServicecenterTaskQueryrefundAPIRequest, resp *tmallservice.TmallServicecenterTaskQueryrefundAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
