package tmallservice

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterAnomalyrecourseHomedecorationQuerybyid 天猫服务平台服务商查询商家投诉单
// tmall.servicecenter.anomalyrecourse.homedecoration.querybyid
//
// 天猫服务平台服务商查询商家投诉单
func TmallServicecenterAnomalyrecourseHomedecorationQuerybyid(clt *core.SDKClient, req *tmallservice.TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIRequest, session string) (*tmallservice.TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIResponse, error) {
	var resp tmallservice.TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
