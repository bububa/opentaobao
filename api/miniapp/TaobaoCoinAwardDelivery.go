package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// TaobaoCoinAwardDelivery 淘金币奖励投放
// taobao.coin.award.delivery
//
// 淘金币奖励投放
func TaobaoCoinAwardDelivery(clt *core.SDKClient, req *miniapp.TaobaoCoinAwardDeliveryAPIRequest, session string) (*miniapp.TaobaoCoinAwardDeliveryAPIResponse, error) {
	var resp miniapp.TaobaoCoinAwardDeliveryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
