package tmallservice

import (
    "github.com/bububa/opentaobao/core"
    "github.com/bububa/opentaobao/model/tmallservice"
)

/* 
服务商确认工人签到成功 
tmall.servicecenter.workcard.signin

服务商确认工人签到成功。需要服务商自己保证工人是在现场服务中。否则虚假回传签到而引起的后续问题全部由服务商自己承担
*/
func TmallServicecenterWorkcardSignin(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardSigninAPIRequest, session string) (*tmallservice.TmallServicecenterWorkcardSigninAPIResponse, error) {
    var resp tmallservice.TmallServicecenterWorkcardSigninAPIResponse
    err := clt.Post(req, &resp, session)
    if err != nil {
        return nil, err
    }
    return &resp, nil
}
