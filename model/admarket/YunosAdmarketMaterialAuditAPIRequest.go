package admarket

import (
	"net/url"
	"sync"

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
		Params: model.NewParams(1),
	}
}

// Reset IRequest interface 方法, 清空结构体
func (r *YunosAdmarketMaterialAuditAPIRequest) Reset() {
	r._sspMaterialAuditResult = nil
	r.Params.ToZero()
}

// GetApiMethodName IRequest interface 方法, 获取Api method
func (r YunosAdmarketMaterialAuditAPIRequest) GetApiMethodName() string {
	return "yunos.admarket.material.audit"
}

// GetApiParams IRequest interface 方法, 获取API参数
func (r YunosAdmarketMaterialAuditAPIRequest) GetApiParams(params url.Values) {
	for k, v := range r.Params {
		params.Set(k, v.String())
	}
}

// GetRawParams IRequest interface 方法, 获取API原始参数
func (r YunosAdmarketMaterialAuditAPIRequest) GetRawParams() model.Params {
	return r.Params
}

// SetSspMaterialAuditResult is SspMaterialAuditResult Setter
// 创意审核结果
func (r *YunosAdmarketMaterialAuditAPIRequest) SetSspMaterialAuditResult(_sspMaterialAuditResult *SspMaterialAuditResult) error {
	r._sspMaterialAuditResult = _sspMaterialAuditResult
	r.Set("ssp_material_audit_result", _sspMaterialAuditResult)
	return nil
}

// GetSspMaterialAuditResult SspMaterialAuditResult Getter
func (r YunosAdmarketMaterialAuditAPIRequest) GetSspMaterialAuditResult() *SspMaterialAuditResult {
	return r._sspMaterialAuditResult
}

var poolYunosAdmarketMaterialAuditAPIRequest = sync.Pool{
	New: func() any {
		return NewYunosAdmarketMaterialAuditRequest()
	},
}

// GetYunosAdmarketMaterialAuditRequest 从 sync.Pool 获取 YunosAdmarketMaterialAuditAPIRequest
func GetYunosAdmarketMaterialAuditAPIRequest() *YunosAdmarketMaterialAuditAPIRequest {
	return poolYunosAdmarketMaterialAuditAPIRequest.Get().(*YunosAdmarketMaterialAuditAPIRequest)
}

// ReleaseYunosAdmarketMaterialAuditAPIRequest 将 YunosAdmarketMaterialAuditAPIRequest 放入 sync.Pool
func ReleaseYunosAdmarketMaterialAuditAPIRequest(v *YunosAdmarketMaterialAuditAPIRequest) {
	v.Reset()
	poolYunosAdmarketMaterialAuditAPIRequest.Put(v)
}
