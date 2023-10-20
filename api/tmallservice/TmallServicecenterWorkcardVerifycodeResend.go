package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardVerifycodeResend 重发核销码
// tmall.servicecenter.workcard.verifycode.resend
//
// 重发核销码
func TmallServicecenterWorkcardVerifycodeResend(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardVerifycodeResendAPIRequest, resp *tmallservice.TmallServicecenterWorkcardVerifycodeResendAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
