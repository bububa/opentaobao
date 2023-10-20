package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardSuspend 工单挂起
// tmall.servicecenter.workcard.suspend
//
// 工单挂起
func TmallServicecenterWorkcardSuspend(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardSuspendAPIRequest, resp *tmallservice.TmallServicecenterWorkcardSuspendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
