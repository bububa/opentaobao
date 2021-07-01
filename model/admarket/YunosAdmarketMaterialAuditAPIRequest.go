package admarket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* YunosAdmarketMaterialAuditAPIRequest
广告平台创意审核 API请求
yunos.admarket.material.audit

用于厂商上报广告平台审核结果 */
type YunosAdmarketMaterialAuditAPIRequest struct {
	model.Params
	// 创意审核结果
	_sspMaterialAuditResult *SspMaterialAuditResult
}

// New
