package jst

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jst"
)

/* TaobaoOcApRuleCreate
创建分账规则
taobao.oc.ap.rule.create

OC分账业务功能支持，用于创建分账规则 */
func TaobaoOcApRuleCreate(clt *core.SDKClient, req *jst.TaobaoOcApRuleCreateAPIRequest, session string) (*jst.TaobaoOcApRuleCreateAPIResponse, error) {
	var resp jst.TaobaoOcApRuleCreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
