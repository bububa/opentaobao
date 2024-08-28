package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAiContentBusinessGetThirdCycleVipStatus 天猫精灵商业化获取三方连续包会员状态
// alibaba.ai.content.business.get.third.cycle.vip.status
//
// 天猫精灵商业化获取三方连续包会员状态
func AlibabaAiContentBusinessGetThirdCycleVipStatus(ctx context.Context, clt *core.SDKClient, req *tmallgenie.AlibabaAiContentBusinessGetThirdCycleVipStatusAPIRequest, resp *tmallgenie.AlibabaAiContentBusinessGetThirdCycleVipStatusAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
