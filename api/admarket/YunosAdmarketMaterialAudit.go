package admarket

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/admarket"
)

// YunosAdmarketMaterialAudit 广告平台创意审核
// yunos.admarket.material.audit
//
// 用于厂商上报广告平台审核结果
func YunosAdmarketMaterialAudit(clt *core.SDKClient, req *admarket.YunosAdmarketMaterialAuditAPIRequest, resp *admarket.YunosAdmarketMaterialAuditAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
