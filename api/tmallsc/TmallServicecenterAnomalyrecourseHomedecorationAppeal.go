package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterAnomalyrecourseHomedecorationAppeal 天猫服务平台商家投诉单服务商申诉接口
// tmall.servicecenter.anomalyrecourse.homedecoration.appeal
//
// 天猫服务平台商家投诉单服务商申诉接口
func TmallServicecenterAnomalyrecourseHomedecorationAppeal(clt *core.SDKClient, req *tmallsc.TmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest, resp *tmallsc.TmallServicecenterAnomalyrecourseHomedecorationAppealAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
