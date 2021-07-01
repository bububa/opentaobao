package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmRuleLevelQuerylevelruleAPIRequest
查询会员等级规则 API请求
alibaba.alsc.crm.rule.level.querylevelrule

查询会员等级规则 */
type AlibabaAlscCrmRuleLevelQuerylevelruleAPIRequest struct {
	model.Params
	// 请求参数
	_planRuleQueryRequest *PlanRuleQueryOpenReq
}

// New
