package charity

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/charity"
)

// AlibabaCharityUseractionSync 用户公益行为同步
// alibaba.charity.useraction.sync
//
// 外部公益活动，用户公益行为同步
func AlibabaCharityUseractionSync(ctx context.Context, clt *core.SDKClient, req *charity.AlibabaCharityUseractionSyncAPIRequest, resp *charity.AlibabaCharityUseractionSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
