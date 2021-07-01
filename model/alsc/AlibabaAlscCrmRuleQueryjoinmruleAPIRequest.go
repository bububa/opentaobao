package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmRuleQueryjoinmruleAPIRequest
查询品牌下的成为会员规则 API请求
alibaba.alsc.crm.rule.queryjoinmrule

查询品牌下的成为会员规则 */
type AlibabaAlscCrmRuleQueryjoinmruleAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq
}

// New
