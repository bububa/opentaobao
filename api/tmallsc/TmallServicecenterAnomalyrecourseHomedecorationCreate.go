package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterAnomalyrecourseHomedecorationCreate 天猫服务平台服务商代商家发起投诉单
// tmall.servicecenter.anomalyrecourse.homedecoration.create
//
// 天猫服务平台服务商代商家发起投诉单
func TmallServicecenterAnomalyrecourseHomedecorationCreate(clt *core.SDKClient, req *tmallsc.TmallServicecenterAnomalyrecourseHomedecorationCreateAPIRequest, session string) (*tmallsc.TmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponse, error) {
	var resp tmallsc.TmallServicecenterAnomalyrecourseHomedecorationCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
