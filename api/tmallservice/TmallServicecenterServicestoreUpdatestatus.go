package tmallservice

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallservice"
)

// TmallServicecenterServicestoreUpdatestatus 网点/门店状态修改
// tmall.servicecenter.servicestore.updatestatus
//
// 修改网点/门店状态
func TmallServicecenterServicestoreUpdatestatus(ctx context.Context, clt *core.SDKClient, req *tmallservice.TmallServicecenterServicestoreUpdatestatusAPIRequest, resp *tmallservice.TmallServicecenterServicestoreUpdatestatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
