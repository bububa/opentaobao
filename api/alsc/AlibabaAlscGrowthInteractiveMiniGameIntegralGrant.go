package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscGrowthInteractiveMiniGameIntegralGrant 本地生活用户增长互动小游戏积分发放
// alibaba.alsc.growth.interactive.mini.game.integral.grant
//
// 本地生活用户增长互动小游戏积分发放
func AlibabaAlscGrowthInteractiveMiniGameIntegralGrant(clt *core.SDKClient, req *alsc.AlibabaAlscGrowthInteractiveMiniGameIntegralGrantAPIRequest, session string) (*alsc.AlibabaAlscGrowthInteractiveMiniGameIntegralGrantAPIResponse, error) {
	var resp alsc.AlibabaAlscGrowthInteractiveMiniGameIntegralGrantAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
