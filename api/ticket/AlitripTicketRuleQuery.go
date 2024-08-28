package ticket

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/ticket"
)

// AlitripTicketRuleQuery 【门票API2.0】门票规则信息查询接口
// alitrip.ticket.rule.query
//
// 门票规则信息查询接口：返回商家上传的门票规则信息
func AlitripTicketRuleQuery(ctx context.Context, clt *core.SDKClient, req *ticket.AlitripTicketRuleQueryAPIRequest, resp *ticket.AlitripTicketRuleQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
