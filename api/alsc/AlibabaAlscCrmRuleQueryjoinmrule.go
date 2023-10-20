package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmRuleQueryjoinmrule 查询品牌下的成为会员规则
// alibaba.alsc.crm.rule.queryjoinmrule
//
// 查询品牌下的成为会员规则
func AlibabaAlscCrmRuleQueryjoinmrule(clt *core.SDKClient, req *alsc.AlibabaAlscCrmRuleQueryjoinmruleAPIRequest, resp *alsc.AlibabaAlscCrmRuleQueryjoinmruleAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
