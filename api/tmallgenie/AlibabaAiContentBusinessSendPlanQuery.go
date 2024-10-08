package tmallgenie

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tmallgenie"
)

// AlibabaAiContentBusinessSendPlanQuery 内容商业化发放权益查询
// alibaba.ai.content.business.send.plan.query
//
// 天猫精灵内容生态的权益查询
func AlibabaAiContentBusinessSendPlanQuery(ctx context.Context, clt *core.SDKClient, req *tmallgenie.AlibabaAiContentBusinessSendPlanQueryAPIRequest, resp *tmallgenie.AlibabaAiContentBusinessSendPlanQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
