package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// TaobaoCoinAwardDelivery 淘金币奖励投放
// taobao.coin.award.delivery
//
// 淘金币奖励投放
func TaobaoCoinAwardDelivery(clt *core.SDKClient, req *miniapp.TaobaoCoinAwardDeliveryAPIRequest, resp *miniapp.TaobaoCoinAwardDeliveryAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
