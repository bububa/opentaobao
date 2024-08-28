package tmallsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQuery 天猫服务平台商家投诉单问题列表查询
// tmall.servicecenter.anomalyrecourse.homedecoration.questioncode.query
//
// 天猫服务平台商家投诉单问题列表查询
func TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQuery(ctx context.Context, clt *core.SDKClient, req *tmallsc.TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIRequest, resp *tmallsc.TmallServicecenterAnomalyrecourseHomedecorationQuestioncodeQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
