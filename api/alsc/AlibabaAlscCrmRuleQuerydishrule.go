package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmRuleQuerydishrule 查询品牌下的入会菜品规则
// alibaba.alsc.crm.rule.querydishrule
//
// 查询品牌下的入会菜品规则
func AlibabaAlscCrmRuleQuerydishrule(clt *core.SDKClient, req *alsc.AlibabaAlscCrmRuleQuerydishruleAPIRequest, resp *alsc.AlibabaAlscCrmRuleQuerydishruleAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
