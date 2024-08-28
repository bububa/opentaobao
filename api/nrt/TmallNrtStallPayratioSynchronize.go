package nrt

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtStallPayratioSynchronize 同步摊位收银比例
// tmall.nrt.stall.payratio.synchronize
//
// ISV同步摊位收银比例到阿里
func TmallNrtStallPayratioSynchronize(ctx context.Context, clt *core.SDKClient, req *nrt.TmallNrtStallPayratioSynchronizeAPIRequest, resp *nrt.TmallNrtStallPayratioSynchronizeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
