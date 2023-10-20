package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecodeconsume 天猫服务平台服务核销
// tmall.service.code.consume
//
// 天猫服务平台－服务核销
func Tmallservicecodeconsume(clt *core.SDKClient, req *tmallservice.TmallservicecodeconsumeAPIRequest, session string) (*tmallservice.TmallservicecodeconsumeAPIResponse, error) {
	var resp tmallservice.TmallservicecodeconsumeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
