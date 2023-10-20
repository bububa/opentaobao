package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// Alibabainteractisvlotteryisvdraw 天猫奖池鉴权接口
// alibaba.interact.isvlottery.isvdraw
//
// 鉴权接口，为tida.isvdraw接口鉴权
func Alibabainteractisvlotteryisvdraw(clt *core.SDKClient, req *interact.AlibabainteractisvlotteryisvdrawAPIRequest, session string) (*interact.AlibabainteractisvlotteryisvdrawAPIResponse, error) {
	var resp interact.AlibabainteractisvlotteryisvdrawAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
