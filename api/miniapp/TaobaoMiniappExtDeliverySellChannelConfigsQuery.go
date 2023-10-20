package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// TaobaoMiniappExtDeliverySellChannelConfigsQuery 查询商家配置的信息
// taobao.miniapp.ext.delivery.sell.channel.configs.query
//
// 查询商家配置的信息
func TaobaoMiniappExtDeliverySellChannelConfigsQuery(clt *core.SDKClient, req *miniapp.TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIRequest, resp *miniapp.TaobaoMiniappExtDeliverySellChannelConfigsQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
