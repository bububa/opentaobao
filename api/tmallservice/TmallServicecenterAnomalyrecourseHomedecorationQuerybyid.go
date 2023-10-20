package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterAnomalyrecourseHomedecorationQuerybyid 天猫服务平台服务商查询商家投诉单
// tmall.servicecenter.anomalyrecourse.homedecoration.querybyid
//
// 天猫服务平台服务商查询商家投诉单
func TmallServicecenterAnomalyrecourseHomedecorationQuerybyid(clt *core.SDKClient, req *tmallservice.TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIRequest, resp *tmallservice.TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
