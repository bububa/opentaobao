package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardComplete 工单完结
// tmall.servicecenter.workcard.complete
//
// 工单完结
func TmallServicecenterWorkcardComplete(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardCompleteAPIRequest, resp *tmallservice.TmallServicecenterWorkcardCompleteAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
