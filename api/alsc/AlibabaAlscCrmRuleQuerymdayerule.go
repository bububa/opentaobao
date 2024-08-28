package alsc

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmRuleQuerymdayerule 查询品牌下的会员日规则
// alibaba.alsc.crm.rule.querymdayerule
//
// 查询品牌下的会员日规则
func AlibabaAlscCrmRuleQuerymdayerule(ctx context.Context, clt *core.SDKClient, req *alsc.AlibabaAlscCrmRuleQuerymdayeruleAPIRequest, resp *alsc.AlibabaAlscCrmRuleQuerymdayeruleAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
