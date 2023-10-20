package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmOpenRuleGet 查询规则
// alibaba.alsc.crm.open.rule.get
//
// 查询会员规则
func AlibabaAlscCrmOpenRuleGet(clt *core.SDKClient, req *alsc.AlibabaAlscCrmOpenRuleGetAPIRequest, resp *alsc.AlibabaAlscCrmOpenRuleGetAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
