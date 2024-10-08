package omniorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/omniorder"
)

// TaobaoOmniorderAllocatedinfoSync 分单结果同步给星盘
// taobao.omniorder.allocatedinfo.sync
//
// ISV分单完成，将分单结果同步给星盘
func TaobaoOmniorderAllocatedinfoSync(ctx context.Context, clt *core.SDKClient, req *omniorder.TaobaoOmniorderAllocatedinfoSyncAPIRequest, resp *omniorder.TaobaoOmniorderAllocatedinfoSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
