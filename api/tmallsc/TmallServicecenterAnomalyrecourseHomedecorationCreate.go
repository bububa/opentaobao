package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterAnomalyrecourseHomedecorationCreate 天猫服务平台服务商代商家发起投诉单
// tmall.servicecenter.anomalyrecourse.homedecoration.create
//
// 天猫服务平台服务商代商家发起投诉单
func TmallServicecenterAnomalyrecourseHomedecorationCreate(clt *core.SDKClient, req *tmallsc.TmallServicecenterAnomalyrecourseHomedecorationCreateAPIRequest, resp *tmallsc.TmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
