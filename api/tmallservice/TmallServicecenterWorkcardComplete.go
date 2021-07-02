package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardComplete 工单完结
// tmall.servicecenter.workcard.complete
//
// 工单完结
func TmallServicecenterWorkcardComplete(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardCompleteAPIRequest, session string) (*tmallservice.TmallServicecenterWorkcardCompleteAPIResponse, error) {
	var resp tmallservice.TmallServicecenterWorkcardCompleteAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
