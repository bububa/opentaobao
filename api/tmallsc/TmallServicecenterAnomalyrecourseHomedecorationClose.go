package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterAnomalyrecourseHomedecorationClose 天猫服务平台商家投诉单服务商完结接口
// tmall.servicecenter.anomalyrecourse.homedecoration.close
//
// 天猫服务平台商家投诉单服务商完结接口
func TmallServicecenterAnomalyrecourseHomedecorationClose(clt *core.SDKClient, req *tmallsc.TmallServicecenterAnomalyrecourseHomedecorationCloseAPIRequest, session string) (*tmallsc.TmallServicecenterAnomalyrecourseHomedecorationCloseAPIResponse, error) {
	var resp tmallsc.TmallServicecenterAnomalyrecourseHomedecorationCloseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
