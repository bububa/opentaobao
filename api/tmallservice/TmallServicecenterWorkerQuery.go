package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkerquery 工人信息查询
// tmall.servicecenter.worker.query
//
// 查询服务商对应的工人信息
func Tmallservicecenterworkerquery(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkerqueryAPIRequest, session string) (*tmallservice.TmallservicecenterworkerqueryAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkerqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
