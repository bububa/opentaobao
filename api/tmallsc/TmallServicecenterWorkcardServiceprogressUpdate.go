package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterWorkcardServiceprogressUpdate 回传工单服务进度
// tmall.servicecenter.workcard.serviceprogress.update
//
// 回传工单服务进度
func TmallServicecenterWorkcardServiceprogressUpdate(clt *core.SDKClient, req *tmallsc.TmallServicecenterWorkcardServiceprogressUpdateAPIRequest, resp *tmallsc.TmallServicecenterWorkcardServiceprogressUpdateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
