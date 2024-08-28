package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterAnomalyrecourseQuerybyid 根据一键求助id查询指定服务商的一键求助单
// tmall.servicecenter.anomalyrecourse.querybyid
//
// 根据一键求助id查询指定服务商的一键求助单
func TmallServicecenterAnomalyrecourseQuerybyid(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterAnomalyrecourseQuerybyidAPIRequest, resp *tmallservice.TmallServicecenterAnomalyrecourseQuerybyidAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
