package tmallnr

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallnr"
)

// TmallNrFulfillLogisticsSync 同城配物流信息回传
// tmall.nr.fulfill.logistics.sync
//
// 同城配业务物流信息回传，通过接口将物流信息同步给天猫
func TmallNrFulfillLogisticsSync(ctx context.Context, clt *core.SDKClient, req *tmallnr.TmallNrFulfillLogisticsSyncAPIRequest, resp *tmallnr.TmallNrFulfillLogisticsSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
