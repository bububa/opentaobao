package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmRuleQuerygrowrule 查询品牌下的会员成长规则
// alibaba.alsc.crm.rule.querygrowrule
//
// 查询品牌下的会员成长规则
func AlibabaAlscCrmRuleQuerygrowrule(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmRuleQuerygrowruleAPIRequest, resp *alsc.AlibabaAlscCrmRuleQuerygrowruleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
