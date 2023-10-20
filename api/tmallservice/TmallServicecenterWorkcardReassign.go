package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardReassign 工单改派门店
// tmall.servicecenter.workcard.reassign
//
// 工单改派门店
func TmallServicecenterWorkcardReassign(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardReassignAPIRequest, resp *tmallservice.TmallServicecenterWorkcardReassignAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
