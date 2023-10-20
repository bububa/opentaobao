package tanx

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTanxAuditCreativeAddAPIResponse 创意预审新增接口 API返回值
// taobao.tanx.audit.creative.add
//
// 创意预审新增接口
type TaobaoTanxAuditCreativeAddAPIResponse struct {
	model.CommonResponse
	TaobaoTanxAuditCreativeAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTanxAuditCreativeAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTanxAuditCreativeAddAPIResponseModel).Reset()
}

// TaobaoTanxAuditCreativeAddAPIResponseModel is 创意预审新增接口 成功返回结果
type TaobaoTanxAuditCreativeAddAPIResponseModel struct {
	XMLName xml.Name `xml:"tanx_audit_creative_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用的成功信息或失败信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用返回码
	TanxErrorCode int64 `json:"tanx_error_code,omitempty" xml:"tanx_error_code,omitempty"`
	// 是否成功
	IsOk bool `json:"is_ok,omitempty" xml:"is_ok,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTanxAuditCreativeAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.TanxErrorCode = 0
	m.IsOk = false
}

var poolTaobaoTanxAuditCreativeAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTanxAuditCreativeAddAPIResponse)
	},
}

// GetTaobaoTanxAuditCreativeAddAPIResponse 从 sync.Pool 获取 TaobaoTanxAuditCreativeAddAPIResponse
func GetTaobaoTanxAuditCreativeAddAPIResponse() *TaobaoTanxAuditCreativeAddAPIResponse {
	return poolTaobaoTanxAuditCreativeAddAPIResponse.Get().(*TaobaoTanxAuditCreativeAddAPIResponse)
}

// ReleaseTaobaoTanxAuditCreativeAddAPIResponse 将 TaobaoTanxAuditCreativeAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTanxAuditCreativeAddAPIResponse(v *TaobaoTanxAuditCreativeAddAPIResponse) {
	v.Reset()
	poolTaobaoTanxAuditCreativeAddAPIResponse.Put(v)
}
