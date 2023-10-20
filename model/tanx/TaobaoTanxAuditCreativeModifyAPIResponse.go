package tanx

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTanxAuditCreativeModifyAPIResponse 创意修改接口 API返回值
// taobao.tanx.audit.creative.modify
//
// 创意修改接口
type TaobaoTanxAuditCreativeModifyAPIResponse struct {
	model.CommonResponse
	TaobaoTanxAuditCreativeModifyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTanxAuditCreativeModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTanxAuditCreativeModifyAPIResponseModel).Reset()
}

// TaobaoTanxAuditCreativeModifyAPIResponseModel is 创意修改接口 成功返回结果
type TaobaoTanxAuditCreativeModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"tanx_audit_creative_modify_response"`
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
func (m *TaobaoTanxAuditCreativeModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.TanxErrorCode = 0
	m.IsOk = false
}

var poolTaobaoTanxAuditCreativeModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTanxAuditCreativeModifyAPIResponse)
	},
}

// GetTaobaoTanxAuditCreativeModifyAPIResponse 从 sync.Pool 获取 TaobaoTanxAuditCreativeModifyAPIResponse
func GetTaobaoTanxAuditCreativeModifyAPIResponse() *TaobaoTanxAuditCreativeModifyAPIResponse {
	return poolTaobaoTanxAuditCreativeModifyAPIResponse.Get().(*TaobaoTanxAuditCreativeModifyAPIResponse)
}

// ReleaseTaobaoTanxAuditCreativeModifyAPIResponse 将 TaobaoTanxAuditCreativeModifyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTanxAuditCreativeModifyAPIResponse(v *TaobaoTanxAuditCreativeModifyAPIResponse) {
	v.Reset()
	poolTaobaoTanxAuditCreativeModifyAPIResponse.Put(v)
}
