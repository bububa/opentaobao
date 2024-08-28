package nrt

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/nrt"
)

// TmallNrtStallSynchronize 摊位信息同步
// tmall.nrt.stall.synchronize
//
// 摊位信息同步
func TmallNrtStallSynchronize(ctx context.Context, clt *core.SDKClient, req *nrt.TmallNrtStallSynchronizeAPIRequest, resp *nrt.TmallNrtStallSynchronizeAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
