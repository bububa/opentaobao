package tanx

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTanxCreativeGetAPIResponse 获取单个审核创意状态 API返回值
// taobao.tanx.creative.get
//
// 获取单个审核创意状态
type TaobaoTanxCreativeGetAPIResponse struct {
	model.CommonResponse
	TaobaoTanxCreativeGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTanxCreativeGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTanxCreativeGetAPIResponseModel).Reset()
}

// TaobaoTanxCreativeGetAPIResponseModel is 获取单个审核创意状态 成功返回结果
type TaobaoTanxCreativeGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tanx_creative_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 调用的成功信息或失败信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用返回码
	TanxErrorCode int64 `json:"tanx_error_code,omitempty" xml:"tanx_error_code,omitempty"`
	// 创意查询返回结果列表
	Result *CreativeAuditDto `json:"result,omitempty" xml:"result,omitempty"`
	// 是否成功
	IsOk bool `json:"is_ok,omitempty" xml:"is_ok,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTanxCreativeGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.TanxErrorCode = 0
	m.Result = nil
	m.IsOk = false
}

var poolTaobaoTanxCreativeGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTanxCreativeGetAPIResponse)
	},
}

// GetTaobaoTanxCreativeGetAPIResponse 从 sync.Pool 获取 TaobaoTanxCreativeGetAPIResponse
func GetTaobaoTanxCreativeGetAPIResponse() *TaobaoTanxCreativeGetAPIResponse {
	return poolTaobaoTanxCreativeGetAPIResponse.Get().(*TaobaoTanxCreativeGetAPIResponse)
}

// ReleaseTaobaoTanxCreativeGetAPIResponse 将 TaobaoTanxCreativeGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTanxCreativeGetAPIResponse(v *TaobaoTanxCreativeGetAPIResponse) {
	v.Reset()
	poolTaobaoTanxCreativeGetAPIResponse.Put(v)
}
