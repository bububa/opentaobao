package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardReassign 工单改派门店
// tmall.servicecenter.workcard.reassign
//
// 工单改派门店
func TmallServicecenterWorkcardReassign(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardReassignAPIRequest, session string) (*tmallservice.TmallServicecenterWorkcardReassignAPIResponse, error) {
	var resp tmallservice.TmallServicecenterWorkcardReassignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
