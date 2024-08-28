package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterSpserviceorderCreate 服务单创建
// tmall.servicecenter.spserviceorder.create
//
// 服务单创建
func TmallServicecenterSpserviceorderCreate(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterSpserviceorderCreateAPIRequest, resp *tmallservice.TmallServicecenterSpserviceorderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
