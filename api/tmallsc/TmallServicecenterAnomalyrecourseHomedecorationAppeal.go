package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterAnomalyrecourseHomedecorationAppeal 天猫服务平台商家投诉单服务商申诉接口
// tmall.servicecenter.anomalyrecourse.homedecoration.appeal
//
// 天猫服务平台商家投诉单服务商申诉接口
func TmallServicecenterAnomalyrecourseHomedecorationAppeal(clt *core.SDKClient, req *tmallsc.TmallServicecenterAnomalyrecourseHomedecorationAppealAPIRequest, session string) (*tmallsc.TmallServicecenterAnomalyrecourseHomedecorationAppealAPIResponse, error) {
	var resp tmallsc.TmallServicecenterAnomalyrecourseHomedecorationAppealAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
