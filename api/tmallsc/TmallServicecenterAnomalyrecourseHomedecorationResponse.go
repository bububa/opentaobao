package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterAnomalyrecourseHomedecorationResponse 天猫服务平台商家投诉单服务商响应接口
// tmall.servicecenter.anomalyrecourse.homedecoration.response
//
// 天猫服务平台商家投诉单服务商响应接口
func TmallServicecenterAnomalyrecourseHomedecorationResponse(clt *core.SDKClient, req *tmallsc.TmallServicecenterAnomalyrecourseHomedecorationResponseAPIRequest, session string) (*tmallsc.TmallServicecenterAnomalyrecourseHomedecorationResponseAPIResponse, error) {
	var resp tmallsc.TmallServicecenterAnomalyrecourseHomedecorationResponseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
