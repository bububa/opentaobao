package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterAnomalyrecourseClose 服务投诉问题单关单
// tmall.servicecenter.anomalyrecourse.close
//
// 提供给服务商在投诉单完结时调用，关闭投诉问题工单。
func TmallServicecenterAnomalyrecourseClose(clt *core.SDKClient, req *tmallsc.TmallServicecenterAnomalyrecourseCloseAPIRequest, session string) (*tmallsc.TmallServicecenterAnomalyrecourseCloseAPIResponse, error) {
	var resp tmallsc.TmallServicecenterAnomalyrecourseCloseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
