package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmOpenRuleGetAPIRequest
查询规则 API请求
alibaba.alsc.crm.open.rule.get

查询会员规则 */
type AlibabaAlscCrmOpenRuleGetAPIRequest struct {
	model.Params
	// 入参
	_paramRuleOpenReq *RuleOpenReq
}

// New
