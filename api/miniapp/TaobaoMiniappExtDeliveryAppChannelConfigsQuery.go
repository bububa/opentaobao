package miniapp

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// TaobaoMiniappExtDeliveryAppChannelConfigsQuery ISV查询应用的渠道信息
// taobao.miniapp.ext.delivery.app.channel.configs.query
//
// ISV查询应用的渠道信息
func TaobaoMiniappExtDeliveryAppChannelConfigsQuery(ctx context.Context, clt *core.SDKClient, req *miniapp.TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIRequest, resp *miniapp.TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
