package tanx

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTanxAuditDepositcreativeAddAPIResponse dsp托管创意新增接口 API返回值
// taobao.tanx.audit.depositcreative.add
//
// dsp托管创意新增接口
type TaobaoTanxAuditDepositcreativeAddAPIResponse struct {
	model.CommonResponse
	TaobaoTanxAuditDepositcreativeAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTanxAuditDepositcreativeAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTanxAuditDepositcreativeAddAPIResponseModel).Reset()
}

// TaobaoTanxAuditDepositcreativeAddAPIResponseModel is dsp托管创意新增接口 成功返回结果
type TaobaoTanxAuditDepositcreativeAddAPIResponseModel struct {
	XMLName xml.Name `xml:"tanx_audit_depositcreative_add_response"`
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
func (m *TaobaoTanxAuditDepositcreativeAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.TanxErrorCode = 0
	m.IsOk = false
}

var poolTaobaoTanxAuditDepositcreativeAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTanxAuditDepositcreativeAddAPIResponse)
	},
}

// GetTaobaoTanxAuditDepositcreativeAddAPIResponse 从 sync.Pool 获取 TaobaoTanxAuditDepositcreativeAddAPIResponse
func GetTaobaoTanxAuditDepositcreativeAddAPIResponse() *TaobaoTanxAuditDepositcreativeAddAPIResponse {
	return poolTaobaoTanxAuditDepositcreativeAddAPIResponse.Get().(*TaobaoTanxAuditDepositcreativeAddAPIResponse)
}

// ReleaseTaobaoTanxAuditDepositcreativeAddAPIResponse 将 TaobaoTanxAuditDepositcreativeAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTanxAuditDepositcreativeAddAPIResponse(v *TaobaoTanxAuditDepositcreativeAddAPIResponse) {
	v.Reset()
	poolTaobaoTanxAuditDepositcreativeAddAPIResponse.Put(v)
}
