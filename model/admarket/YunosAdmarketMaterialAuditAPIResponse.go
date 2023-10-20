package admarket

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// YunosAdmarketMaterialAuditAPIResponse 广告平台创意审核 API返回值
// yunos.admarket.material.audit
//
// 用于厂商上报广告平台审核结果
type YunosAdmarketMaterialAuditAPIResponse struct {
	model.CommonResponse
	YunosAdmarketMaterialAuditAPIResponseModel
}

// Reset 清空结构体
func (m *YunosAdmarketMaterialAuditAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.YunosAdmarketMaterialAuditAPIResponseModel).Reset()
}

// YunosAdmarketMaterialAuditAPIResponseModel is 广告平台创意审核 成功返回结果
type YunosAdmarketMaterialAuditAPIResponseModel struct {
	XMLName xml.Name `xml:"yunos_admarket_material_audit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *YunosAdmarketMaterialAuditResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *YunosAdmarketMaterialAuditAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolYunosAdmarketMaterialAuditAPIResponse = sync.Pool{
	New: func() any {
		return new(YunosAdmarketMaterialAuditAPIResponse)
	},
}

// GetYunosAdmarketMaterialAuditAPIResponse 从 sync.Pool 获取 YunosAdmarketMaterialAuditAPIResponse
func GetYunosAdmarketMaterialAuditAPIResponse() *YunosAdmarketMaterialAuditAPIResponse {
	return poolYunosAdmarketMaterialAuditAPIResponse.Get().(*YunosAdmarketMaterialAuditAPIResponse)
}

// ReleaseYunosAdmarketMaterialAuditAPIResponse 将 YunosAdmarketMaterialAuditAPIResponse 保存到 sync.Pool
func ReleaseYunosAdmarketMaterialAuditAPIResponse(v *YunosAdmarketMaterialAuditAPIResponse) {
	v.Reset()
	poolYunosAdmarketMaterialAuditAPIResponse.Put(v)
}
