package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// Tmallservicecenterworkcardverifycoderesend 重发核销码
// tmall.servicecenter.workcard.verifycode.resend
//
// 重发核销码
func Tmallservicecenterworkcardverifycoderesend(clt *core.SDKClient, req *tmallservice.TmallservicecenterworkcardverifycoderesendAPIRequest, session string) (*tmallservice.TmallservicecenterworkcardverifycoderesendAPIResponse, error) {
	var resp tmallservice.TmallservicecenterworkcardverifycoderesendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
