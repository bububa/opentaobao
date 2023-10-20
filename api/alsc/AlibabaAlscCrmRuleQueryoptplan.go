package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmRuleQueryoptplan 查询运营计划
// alibaba.alsc.crm.rule.queryoptplan
//
// 查询运营计划
func AlibabaAlscCrmRuleQueryoptplan(clt *core.SDKClient, req *alsc.AlibabaAlscCrmRuleQueryoptplanAPIRequest, resp *alsc.AlibabaAlscCrmRuleQueryoptplanAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
