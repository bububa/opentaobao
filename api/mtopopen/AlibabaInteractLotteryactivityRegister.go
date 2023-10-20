package mtopopen

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mtopopen"
)

// AlibabaInteractLotteryactivityRegister 回传抽奖相关参数
// alibaba.interact.lotteryactivity.register
//
// 提供接口供三方应用将数据回传到平台
func AlibabaInteractLotteryactivityRegister(clt *core.SDKClient, req *mtopopen.AlibabaInteractLotteryactivityRegisterAPIRequest, resp *mtopopen.AlibabaInteractLotteryactivityRegisterAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
