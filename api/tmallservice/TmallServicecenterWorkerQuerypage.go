package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkerquerypage 查询工人列表
// tmall.servicecenter.worker.querypage
//
// 服务商查询工人列表
func Tmallservicecenterworkerquerypage(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkerquerypageAPIRequest, session string) (*tmallservice.TmallservicecenterworkerquerypageAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkerquerypageAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
