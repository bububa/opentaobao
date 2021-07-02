package admarket

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

// YunosAdmarketMaterialAuditAPIRequest 广告平台创意审核 API请求
// yunos.admarket.material.audit
//
// 用于厂商上报广告平台审核结果
type YunosAdmarketMaterialAuditAPIRequest struct {
	model.Params
	// 创意审核结果
	_sspMaterialAuditResult *SspMaterialAuditResult
}

// NewYunosAdmarketMaterialAuditRequest 初始化YunosAdmarketMaterialAuditAPIRequest对象
func NewYunosAdmarketMaterialAuditRequest() *YunosAdmarketMaterialAuditAPIRequest {
	return &YunosAdmarketMaterialAuditAPIRequest{
		Params: model.NewParams(),
	}
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosAdmarketMaterialAuditAPIRequest) GetApiMethodName() string {
	return "yunos.admarket.material.audit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosAdmarketMaterialAuditAPIRequest) GetApiParams() url.Values {
	params := url.Values{}
	for k, v := range r.GetRawParams() {
		params.Set(k, v.String())
	}
	return params
}

// Set is SspMaterialAuditResult Setter
// 创意审核结果
func (r *YunosAdmarketMaterialAuditAPIRequest) SetSspMaterialAuditResult(_sspMaterialAuditResult *SspMaterialAuditResult) error {
	r._sspMaterialAuditResult = _sspMaterialAuditResult
	r.Set("ssp_material_audit_result", _sspMaterialAuditResult)
	return nil
}

// Get SspMaterialAuditResult Getter
func (r YunosAdmarketMaterialAuditAPIRequest) GetSspMaterialAuditResult() *SspMaterialAuditResult {
	return r._sspMaterialAuditResult
}
