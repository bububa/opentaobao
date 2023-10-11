package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterAnomalyrecourseHomedecorationAdmit 天猫服务平台商家投诉单服务商认责接口
// tmall.servicecenter.anomalyrecourse.homedecoration.admit
//
// 天猫服务平台商家投诉单服务商认责接口
func TmallServicecenterAnomalyrecourseHomedecorationAdmit(clt *core.SDKClient, req *tmallsc.TmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest, session string) (*tmallsc.TmallServicecenterAnomalyrecourseHomedecorationAdmitAPIResponse, error) {
	var resp tmallsc.TmallServicecenterAnomalyrecourseHomedecorationAdmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
