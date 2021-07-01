package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmRuleQuerympriceruleAPIRequest
查询品牌下的会员价规则 API请求
alibaba.alsc.crm.rule.querympricerule

查询品牌下的会员价规则 */
type AlibabaAlscCrmRuleQuerympriceruleAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq
}

// New
