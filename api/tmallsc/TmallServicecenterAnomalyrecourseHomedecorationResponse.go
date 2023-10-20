package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterAnomalyrecourseHomedecorationResponse 天猫服务平台商家投诉单服务商响应接口
// tmall.servicecenter.anomalyrecourse.homedecoration.response
//
// 天猫服务平台商家投诉单服务商响应接口
func TmallServicecenterAnomalyrecourseHomedecorationResponse(clt *core.SDKClient, req *tmallsc.TmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest, resp *tmallsc.TmallServicecenterAnomalyrecourseHomedecorationResponseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
