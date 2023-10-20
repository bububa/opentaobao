package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterAnomalyrecourseHomedecorationClose 天猫服务平台商家投诉单服务商完结接口
// tmall.servicecenter.anomalyrecourse.homedecoration.close
//
// 天猫服务平台商家投诉单服务商完结接口
func TmallServicecenterAnomalyrecourseHomedecorationClose(clt *core.SDKClient, req *tmallsc.TmallServicecenterAnomalyrecourseHomedecorationCloseAPIRequest, resp *tmallsc.TmallServicecenterAnomalyrecourseHomedecorationCloseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
