package admarket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosadmarketmaterialauditAPIRequest 广告平台创意审核 API请求
// yunos.admarket.material.audit
//
// 用于厂商上报广告平台审核结果
type YunosadmarketmaterialauditAPIRequest struct {
	model.Params
	// 创意审核结果
	_sspMaterialAuditResult *SspMaterialAuditResult
}

// NewYunosadmarketmaterialauditRequest 初始化YunosadmarketmaterialauditAPIRequest对象
func NewYunosadmarketmaterialauditRequest() *YunosadmarketmaterialauditAPIRequest {
	return &YunosadmarketmaterialauditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosadmarketmaterialauditAPIRequest) GetApiMethodName() string {
	return "yunos.admarket.material.audit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosadmarketmaterialauditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosadmarketmaterialauditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSspMaterialAuditResult is SspMaterialAuditResult Setter
// 创意审核结果
func (r *YunosadmarketmaterialauditAPIRequest) SetSspMaterialAuditResult(_sspMaterialAuditResult *SspMaterialAuditResult) error {
	r._sspMaterialAuditResult = _sspMaterialAuditResult
	r.Set("ssp_material_audit_result", _sspMaterialAuditResult)
	return nil
}

// GetSspMaterialAuditResult SspMaterialAuditResult Getter
func (r YunosadmarketmaterialauditAPIRequest) GetSspMaterialAuditResult() *SspMaterialAuditResult {
	return r._sspMaterialAuditResult
}
