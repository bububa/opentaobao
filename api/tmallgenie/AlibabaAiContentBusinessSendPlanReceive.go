package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAiContentBusinessSendPlanReceive 天猫精灵商业化采销发放计划领取
// alibaba.ai.content.business.send.plan.receive
//
// 天猫精灵商业化采销发放计划领取
func AlibabaAiContentBusinessSendPlanReceive(ctx context.Context, clt *core.SDKClient, req *tmallgenie.AlibabaAiContentBusinessSendPlanReceiveAPIRequest, resp *tmallgenie.AlibabaAiContentBusinessSendPlanReceiveAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
