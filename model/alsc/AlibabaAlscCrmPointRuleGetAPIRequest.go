package alsc

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaAlscCrmPointRuleGetAPIRequest
查询积分规则 API请求
alibaba.alsc.crm.point.rule.get

新增积分规则查询接口,传入includeLogicalDelete和maxUpdateTime时走同步下行逻辑不然则走普通积分规则查询接口 */
type AlibabaAlscCrmPointRuleGetAPIRequest struct {
	model.Params
	// 入参
	_paramQueryPointRuleOpenReq *QueryPointRuleOpenReq
}

// New
