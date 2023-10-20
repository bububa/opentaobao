package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterWorkcardEvaluate 服务商反馈鉴定结果
// tmall.servicecenter.workcard.evaluate
//
// 服务商反馈鉴定结果
func TmallServicecenterWorkcardEvaluate(clt *core.SDKClient, req *tmallservice.TmallServicecenterWorkcardEvaluateAPIRequest, resp *tmallservice.TmallServicecenterWorkcardEvaluateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
