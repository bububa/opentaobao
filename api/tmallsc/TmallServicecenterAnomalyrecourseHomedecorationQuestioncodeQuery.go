package tmallsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQuery 天猫服务平台商家投诉单问题列表查询
// tmall.servicecenter.anomalyrecourse.homedecoration.questioncode.query
//
// 天猫服务平台商家投诉单问题列表查询
func TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQuery(clt *core.SDKClient, req *tmallsc.TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIRequest, session string) (*tmallsc.TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIResponse, error) {
	var resp tmallsc.TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
