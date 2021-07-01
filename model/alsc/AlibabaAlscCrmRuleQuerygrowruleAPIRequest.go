package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmRuleQuerygrowruleAPIRequest
查询品牌下的会员成长规则 API请求
alibaba.alsc.crm.rule.querygrowrule

查询品牌下的会员成长规则 */
type AlibabaAlscCrmRuleQuerygrowruleAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq
}

// New
