package tmallsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallsc"
)

// TmallServicecenterAnomalyrecourseHomedecorationAdmit 天猫服务平台商家投诉单服务商认责接口
// tmall.servicecenter.anomalyrecourse.homedecoration.admit
//
// 天猫服务平台商家投诉单服务商认责接口
func TmallServicecenterAnomalyrecourseHomedecorationAdmit(ctx context.Context, clt *core.SDKClient, req *tmallsc.TmallServicecenterAnomalyrecourseHomedecorationAdmitAPIRequest, resp *tmallsc.TmallServicecenterAnomalyrecourseHomedecorationAdmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
