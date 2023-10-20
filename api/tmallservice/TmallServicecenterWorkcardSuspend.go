package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkcardsuspend 工单挂起
// tmall.servicecenter.workcard.suspend
//
// 工单挂起
func Tmallservicecenterworkcardsuspend(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkcardsuspendAPIRequest, session string) (*tmallservice.TmallservicecenterworkcardsuspendAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkcardsuspendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
