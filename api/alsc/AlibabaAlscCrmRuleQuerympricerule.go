package alsc

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alsc"
)

// AlibabaAlscCrmRuleQuerympricerule 查询品牌下的会员价规则
// alibaba.alsc.crm.rule.querympricerule
//
// 查询品牌下的会员价规则
func AlibabaAlscCrmRuleQuerympricerule(clt *core.SDKClient, req *alsc.AlibabaAlscCrmRuleQuerympriceruleAPIRequest, session string) (*alsc.AlibabaAlscCrmRuleQuerympriceruleAPIResponse, error) {
	var resp alsc.AlibabaAlscCrmRuleQuerympriceruleAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
