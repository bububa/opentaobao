package tmallservice

import (
    "github.com/bububa/opentaobao/model"
)

/* 
服务商确认工人签到成功 APIResponse
tmall.servicecenter.workcard.signin

服务商确认工人签到成功。需要服务商自己保证工人是在现场服务中。否则虚假回传签到而引起的后续问题全部由服务商自己承担
*/
type TmallServicecenterWorkcardSigninAPIResponse struct {
    model.CommonResponse
    Response *TmallServicecenterWorkcardSigninResponse `json:"tmall_servicecenter_workcard_signin_response,omitempty"`
}

type TmallServicecenterWorkcardSigninResponse struct {

    // .
    Result   *TmallServicecenterWorkcardSigninResult `json:"result,omitempty"`

}
