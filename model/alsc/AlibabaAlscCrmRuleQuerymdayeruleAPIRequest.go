package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmRuleQuerymdayeruleAPIRequest
查询品牌下的会员日规则 API请求
alibaba.alsc.crm.rule.querymdayerule

查询品牌下的会员日规则 */
type AlibabaAlscCrmRuleQuerymdayeruleAPIRequest struct {
	model.Params
	// 系统自动生成
	_paramPlanRuleQueryOpenReq *PlanRuleQueryOpenReq
}

// New
