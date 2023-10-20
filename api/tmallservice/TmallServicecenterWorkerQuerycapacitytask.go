package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkerquerycapacitytask 查询需求容量
// tmall.servicecenter.worker.querycapacitytask
//
// 查询需求容量
func Tmallservicecenterworkerquerycapacitytask(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkerquerycapacitytaskAPIRequest, session string) (*tmallservice.TmallservicecenterworkerquerycapacitytaskAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkerquerycapacitytaskAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
