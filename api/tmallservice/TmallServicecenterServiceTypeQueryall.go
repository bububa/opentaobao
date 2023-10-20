package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterservicetypequeryall 服务供应链服务类型
// tmall.servicecenter.service.type.queryall
//
// 查询天猫服务类型列表
func Tmallservicecenterservicetypequeryall(clt *core.SDKClient, req *tmallservice.TmallservicecenterservicetypequeryallAPIRequest, session string) (*tmallservice.TmallservicecenterservicetypequeryallAPIResponse, error) {
	var resp tmallservice.TmallservicecenterservicetypequeryallAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
