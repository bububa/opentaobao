package jym

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/jym"
)

// AlibabaJymSteamRefundAudit 交易猫steam逆向单审核
// alibaba.jym.steam.refund.audit
//
// 交易猫steam逆向单审核
func AlibabaJymSteamRefundAudit(clt *core.SDKClient, req *jym.AlibabaJymSteamRefundAuditAPIRequest, resp *jym.AlibabaJymSteamRefundAuditAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
