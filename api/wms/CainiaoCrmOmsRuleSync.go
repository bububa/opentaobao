package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// CainiaoCrmOmsRuleSync 商家ERP订单处理规则同步
// cainiao.crm.oms.rule.sync
//
// 将商家ERP订单处理规则同步到菜鸟CRM系统
func CainiaoCrmOmsRuleSync(clt *core.SDKClient, req *wms.CainiaoCrmOmsRuleSyncAPIRequest, session string) (*wms.CainiaoCrmOmsRuleSyncAPIResponse, error) {
	var resp wms.CainiaoCrmOmsRuleSyncAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
