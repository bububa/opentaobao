package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmRuleQuerytaglist 查询标签列表
// alibaba.alsc.crm.rule.querytaglist
//
// 查询标签列表
func AlibabaAlscCrmRuleQuerytaglist(clt *core.SDKClient, req *alsc.AlibabaAlscCrmRuleQuerytaglistAPIRequest, session string) (*alsc.AlibabaAlscCrmRuleQuerytaglistAPIResponse, error) {
	var resp alsc.AlibabaAlscCrmRuleQuerytaglistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
