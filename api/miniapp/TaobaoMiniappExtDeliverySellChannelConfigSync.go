package miniapp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// TaobaoMiniappExtDeliverySellChannelConfigSync 写入商家配置信息
// taobao.miniapp.ext.delivery.sell.channel.config.sync
//
// 写入商家配置信息
func TaobaoMiniappExtDeliverySellChannelConfigSync(ctx context.Context, clt *core.SDKClient, req *miniapp.TaobaoMiniappExtDeliverySellChannelConfigSyncAPIRequest, resp *miniapp.TaobaoMiniappExtDeliverySellChannelConfigSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
