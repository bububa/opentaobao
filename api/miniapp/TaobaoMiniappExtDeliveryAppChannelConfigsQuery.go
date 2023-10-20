package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// TaobaoMiniappExtDeliveryAppChannelConfigsQuery ISV查询应用的渠道信息
// taobao.miniapp.ext.delivery.app.channel.configs.query
//
// ISV查询应用的渠道信息
func TaobaoMiniappExtDeliveryAppChannelConfigsQuery(clt *core.SDKClient, req *miniapp.TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIRequest, resp *miniapp.TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
