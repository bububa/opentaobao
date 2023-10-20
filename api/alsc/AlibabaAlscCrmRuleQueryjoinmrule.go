package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmRuleQueryjoinmrule 查询品牌下的成为会员规则
// alibaba.alsc.crm.rule.queryjoinmrule
//
// 查询品牌下的成为会员规则
func AlibabaAlscCrmRuleQueryjoinmrule(clt *core.SDKClient, req *alsc.AlibabaAlscCrmRuleQueryjoinmruleAPIRequest, session string) (*alsc.AlibabaAlscCrmRuleQueryjoinmruleAPIResponse, error) {
	var resp alsc.AlibabaAlscCrmRuleQueryjoinmruleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
