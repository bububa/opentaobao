package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterAnomalyrecourseHomedecorationQuerybyid 天猫服务平台服务商查询商家投诉单
// tmall.servicecenter.anomalyrecourse.homedecoration.querybyid
//
// 天猫服务平台服务商查询商家投诉单
func TmallServicecenterAnomalyrecourseHomedecorationQuerybyid(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIRequest, resp *tmallservice.TmallServicecenterAnomalyrecourseHomedecorationQuerybyidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
