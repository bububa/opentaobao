package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkcardevaluate 服务商反馈鉴定结果
// tmall.servicecenter.workcard.evaluate
//
// 服务商反馈鉴定结果
func Tmallservicecenterworkcardevaluate(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkcardevaluateAPIRequest, session string) (*tmallservice.TmallservicecenterworkcardevaluateAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkcardevaluateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
