package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

/* AlibabaAlscCrmRuleLevelQuerylevelrule
查询会员等级规则
alibaba.alsc.crm.rule.level.querylevelrule

查询会员等级规则 */
func AlibabaAlscCrmRuleLevelQuerylevelrule(clt *core.SDKClient, req *alsc.AlibabaAlscCrmRuleLevelQuerylevelruleAPIRequest, session string) (*alsc.AlibabaAlscCrmRuleLevelQuerylevelruleAPIResponse, error) {
	var resp alsc.AlibabaAlscCrmRuleLevelQuerylevelruleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
