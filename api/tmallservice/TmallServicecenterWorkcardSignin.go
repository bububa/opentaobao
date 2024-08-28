package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardSignin 服务商确认工人签到成功
// tmall.servicecenter.workcard.signin
//
// 服务商确认工人签到成功。需要服务商自己保证工人是在现场服务中。否则虚假回传签到而引起的后续问题全部由服务商自己承担
func TmallServicecenterWorkcardSignin(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardSigninAPIRequest, resp *tmallservice.TmallServicecenterWorkcardSigninAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
