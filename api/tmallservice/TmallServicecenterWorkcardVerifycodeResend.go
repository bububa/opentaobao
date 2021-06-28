package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
重发核销码 
tmall.servicecenter.workcard.verifycode.resend

重发核销码
*/
func TmallServicecenterWorkcardVerifycodeResend(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardVerifycodeResendRequest, session string) (*tmallservice.TmallServicecenterWorkcardVerifycodeResendAPIResponse, error) {
    var resp tmallservice.TmallServicecenterWorkcardVerifycodeResendAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
