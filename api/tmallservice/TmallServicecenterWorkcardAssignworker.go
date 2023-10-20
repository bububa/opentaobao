package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkcardassignworker 服务商分派工人
// tmall.servicecenter.workcard.assignworker
//
// 服务商调用该接口分派工人给具体的工单
func Tmallservicecenterworkcardassignworker(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkcardassignworkerAPIRequest, session string) (*tmallservice.TmallservicecenterworkcardassignworkerAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkcardassignworkerAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
