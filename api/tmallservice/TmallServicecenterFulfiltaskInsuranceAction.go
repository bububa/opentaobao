package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterfulfiltaskinsuranceaction 供应链保险链路动作
// tmall.servicecenter.fulfiltask.insurance.action
//
// 服务供应链履约链路 保险类业务履约接口
func Tmallservicecenterfulfiltaskinsuranceaction(clt *core.SDKClient, req *tmallservice.TmallservicecenterfulfiltaskinsuranceactionAPIRequest, session string) (*tmallservice.TmallservicecenterfulfiltaskinsuranceactionAPIResponse, error) {
	var resp tmallservice.TmallservicecenterfulfiltaskinsuranceactionAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
