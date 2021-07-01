package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

/* TmallServicecenterWorkcardSuspend
工单挂起
tmall.servicecenter.workcard.suspend

工单挂起 */
func TmallServicecenterWorkcardSuspend(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardSuspendAPIRequest, session string) (*tmallservice.TmallServicecenterWorkcardSuspendAPIResponse, error) {
	var resp tmallservice.TmallServicecenterWorkcardSuspendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
