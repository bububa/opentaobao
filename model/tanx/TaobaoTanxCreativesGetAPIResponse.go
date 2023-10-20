package tanx

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoTanxCreativesGetAPIResponse 批量获取DSP用户的创意审核结果 API返回值
// taobao.tanx.creatives.get
//
// 批量获取DSP用户的创意审核结果
type TaobaoTanxCreativesGetAPIResponse struct {
	model.CommonResponse
	TaobaoTanxCreativesGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoTanxCreativesGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoTanxCreativesGetAPIResponseModel).Reset()
}

// TaobaoTanxCreativesGetAPIResponseModel is 批量获取DSP用户的创意审核结果 成功返回结果
type TaobaoTanxCreativesGetAPIResponseModel struct {
	XMLName xml.Name `xml:"tanx_creatives_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回的创意列表
	Results []CreativeDto `json:"results,omitempty" xml:"results>creative_dto,omitempty"`
	// 调用的成功信息或失败信息
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 调用返回码
	TanxErrorCode int64 `json:"tanx_error_code,omitempty" xml:"tanx_error_code,omitempty"`
	// 调用是否成功
	IsOk bool `json:"is_ok,omitempty" xml:"is_ok,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoTanxCreativesGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
	m.Message = ""
	m.TanxErrorCode = 0
	m.IsOk = false
}

var poolTaobaoTanxCreativesGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoTanxCreativesGetAPIResponse)
	},
}

// GetTaobaoTanxCreativesGetAPIResponse 从 sync.Pool 获取 TaobaoTanxCreativesGetAPIResponse
func GetTaobaoTanxCreativesGetAPIResponse() *TaobaoTanxCreativesGetAPIResponse {
	return poolTaobaoTanxCreativesGetAPIResponse.Get().(*TaobaoTanxCreativesGetAPIResponse)
}

// ReleaseTaobaoTanxCreativesGetAPIResponse 将 TaobaoTanxCreativesGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoTanxCreativesGetAPIResponse(v *TaobaoTanxCreativesGetAPIResponse) {
	v.Reset()
	poolTaobaoTanxCreativesGetAPIResponse.Put(v)
}
