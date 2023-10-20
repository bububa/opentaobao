package miniapp

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/miniapp"
)

// Taobaocoinawarddelivery 淘金币奖励投放
// taobao.coin.award.delivery
//
// 淘金币奖励投放
func Taobaocoinawarddelivery(clt *core.SDKClient, req *miniapp.TaobaocoinawarddeliveryAPIRequest, session string) (*miniapp.TaobaocoinawarddeliveryAPIResponse, error) {
	var resp miniapp.TaobaocoinawarddeliveryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
