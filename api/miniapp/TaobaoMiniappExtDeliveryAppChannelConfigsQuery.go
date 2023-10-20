package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// TaobaoMiniappExtDeliveryAppChannelConfigsQuery ISV查询应用的渠道信息
// taobao.miniapp.ext.delivery.app.channel.configs.query
//
// ISV查询应用的渠道信息
func TaobaoMiniappExtDeliveryAppChannelConfigsQuery(clt *core.SDKClient, req *miniapp.TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIRequest, session string) (*miniapp.TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponse, error) {
	var resp miniapp.TaobaoMiniappExtDeliveryAppChannelConfigsQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
