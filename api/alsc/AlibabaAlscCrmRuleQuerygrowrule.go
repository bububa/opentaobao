package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

/* AlibabaAlscCrmRuleQuerygrowrule
查询品牌下的会员成长规则
alibaba.alsc.crm.rule.querygrowrule

查询品牌下的会员成长规则 */
func AlibabaAlscCrmRuleQuerygrowrule(clt *core.SDKClient, req *alsc.AlibabaAlscCrmRuleQuerygrowruleAPIRequest, session string) (*alsc.AlibabaAlscCrmRuleQuerygrowruleAPIResponse, error) {
	var resp alsc.AlibabaAlscCrmRuleQuerygrowruleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
