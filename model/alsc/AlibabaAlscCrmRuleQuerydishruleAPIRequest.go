package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmRuleQuerydishruleAPIRequest
查询品牌下的入会菜品规则 API请求
alibaba.alsc.crm.rule.querydishrule

查询品牌下的入会菜品规则 */
type AlibabaAlscCrmRuleQuerydishruleAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq
}

// New
