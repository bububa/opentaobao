package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterSpserviceorderQuery 服务单列表查询
// tmall.servicecenter.spserviceorder.query
//
// 查询服务单列表
func TmallServicecenterSpserviceorderQuery(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterSpserviceorderQueryAPIRequest, resp *tmallservice.TmallServicecenterSpserviceorderQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
