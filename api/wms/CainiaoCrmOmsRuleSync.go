package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// CainiaoCrmOmsRuleSync 商家ERP订单处理规则同步
// cainiao.crm.oms.rule.sync
//
// 将商家ERP订单处理规则同步到菜鸟CRM系统
func CainiaoCrmOmsRuleSync(clt *core.SDKClient, req *wms.CainiaoCrmOmsRuleSyncAPIRequest, resp *wms.CainiaoCrmOmsRuleSyncAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
