package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkcardextrachargecreate 创建工单额外收费项
// tmall.servicecenter.workcard.extracharge.create
//
// 创建额外收费项
func Tmallservicecenterworkcardextrachargecreate(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkcardextrachargecreateAPIRequest, session string) (*tmallservice.TmallservicecenterworkcardextrachargecreateAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkcardextrachargecreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
