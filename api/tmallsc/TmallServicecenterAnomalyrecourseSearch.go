package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterAnomalyrecourseSearch 天猫服务平台服务商一键求助单查询
// tmall.servicecenter.anomalyrecourse.search
//
// 天猫服务平台服务商一键求助单查询
func TmallServicecenterAnomalyrecourseSearch(clt *core.SDKClient, req *tmallsc.TmallServicecenterAnomalyrecourseSearchAPIRequest, resp *tmallsc.TmallServicecenterAnomalyrecourseSearchAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
