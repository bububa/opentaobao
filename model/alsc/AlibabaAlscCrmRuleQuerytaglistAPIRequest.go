package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmRuleQuerytaglistAPIRequest
查询标签列表 API请求
alibaba.alsc.crm.rule.querytaglist

查询标签列表 */
type AlibabaAlscCrmRuleQuerytaglistAPIRequest struct {
	model.Params
	// 请求参数
	_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq
}

// New
