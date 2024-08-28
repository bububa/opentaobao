package youkuott

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/youkuott"
)

// YoukuOttKittyCommonorderSync 运营商一般订单同步
// youku.ott.kitty.commonorder.sync
//
// 运营商一般订单同步
func YoukuOttKittyCommonorderSync(ctx context.Context, clt *core.SDKClient, req *youkuott.YoukuOttKittyCommonorderSyncAPIRequest, resp *youkuott.YoukuOttKittyCommonorderSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
