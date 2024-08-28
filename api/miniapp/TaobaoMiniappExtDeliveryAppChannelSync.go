package miniapp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// TaobaoMiniappExtDeliveryAppChannelSync ISV写入应用的渠道信息
// taobao.miniapp.ext.delivery.app.channel.sync
//
// ISV写入应用的渠道信息
func TaobaoMiniappExtDeliveryAppChannelSync(ctx context.Context, clt *core.SDKClient, req *miniapp.TaobaoMiniappExtDeliveryAppChannelSyncAPIRequest, resp *miniapp.TaobaoMiniappExtDeliveryAppChannelSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
