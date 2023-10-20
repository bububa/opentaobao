package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// TaobaoMiniappExtDeliveryAppChannelSync ISV写入应用的渠道信息
// taobao.miniapp.ext.delivery.app.channel.sync
//
// ISV写入应用的渠道信息
func TaobaoMiniappExtDeliveryAppChannelSync(clt *core.SDKClient, req *miniapp.TaobaoMiniappExtDeliveryAppChannelSyncAPIRequest, session string) (*miniapp.TaobaoMiniappExtDeliveryAppChannelSyncAPIResponse, error) {
	var resp miniapp.TaobaoMiniappExtDeliveryAppChannelSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
