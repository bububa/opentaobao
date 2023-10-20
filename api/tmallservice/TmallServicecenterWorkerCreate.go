package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkercreate 服务商工人信息创建
// tmall.servicecenter.worker.create
//
// 服务商工人信息创建
func Tmallservicecenterworkercreate(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkercreateAPIRequest, session string) (*tmallservice.TmallservicecenterworkercreateAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
