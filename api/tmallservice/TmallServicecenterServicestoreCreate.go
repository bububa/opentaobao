package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterServicestoreCreate 创建门店
// tmall.servicecenter.servicestore.create
//
// 用于创建门店/网点。多个业务共用
func TmallServicecenterServicestoreCreate(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterServicestoreCreateAPIRequest, resp *tmallservice.TmallServicecenterServicestoreCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
