package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterAnomalyrecourseClose 服务投诉问题单关单
// tmall.servicecenter.anomalyrecourse.close
//
// 提供给服务商在投诉单完结时调用，关闭投诉问题工单。
func TmallServicecenterAnomalyrecourseClose(clt *core.SDKClient, req *tmallsc.TmallServicecenterAnomalyrecourseCloseAPIRequest, resp *tmallsc.TmallServicecenterAnomalyrecourseCloseAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
