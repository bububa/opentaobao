package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// TaobaoMiniappExtDeliverySellChannelConfigSync 写入商家配置信息
// taobao.miniapp.ext.delivery.sell.channel.config.sync
//
// 写入商家配置信息
func TaobaoMiniappExtDeliverySellChannelConfigSync(clt *core.SDKClient, req *miniapp.TaobaoMiniappExtDeliverySellChannelConfigSyncAPIRequest, resp *miniapp.TaobaoMiniappExtDeliverySellChannelConfigSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
