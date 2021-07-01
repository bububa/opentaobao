package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmRuleQueryoptplanAPIRequest
查询运营计划 API请求
alibaba.alsc.crm.rule.queryoptplan

查询运营计划 */
type AlibabaAlscCrmRuleQueryoptplanAPIRequest struct {
	model.Params
	// 请求参数
	_planRuleQueryOpenRequest *PlanRuleQueryOpenReq
}

// New
