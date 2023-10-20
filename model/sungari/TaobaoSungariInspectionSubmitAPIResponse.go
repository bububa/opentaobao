package sungari

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSungariInspectionSubmitAPIResponse 抽检指令录入 API返回值
// taobao.sungari.inspection.submit
//
// 抽检指令录入
type TaobaoSungariInspectionSubmitAPIResponse struct {
	model.CommonResponse
	TaobaoSungariInspectionSubmitAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSungariInspectionSubmitAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSungariInspectionSubmitAPIResponseModel).Reset()
}

// TaobaoSungariInspectionSubmitAPIResponseModel is 抽检指令录入 成功返回结果
type TaobaoSungariInspectionSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"sungari_inspection_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 状态码
	ReturnCode int64 `json:"return_code,omitempty" xml:"return_code,omitempty"`
	// 是否成功
	Data *InspectionResultInfo `json:"data,omitempty" xml:"data,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSungariInspectionSubmitAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Message = ""
	m.ReturnCode = 0
	m.Data = nil
}

var poolTaobaoSungariInspectionSubmitAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSungariInspectionSubmitAPIResponse)
	},
}

// GetTaobaoSungariInspectionSubmitAPIResponse 从 sync.Pool 获取 TaobaoSungariInspectionSubmitAPIResponse
func GetTaobaoSungariInspectionSubmitAPIResponse() *TaobaoSungariInspectionSubmitAPIResponse {
	return poolTaobaoSungariInspectionSubmitAPIResponse.Get().(*TaobaoSungariInspectionSubmitAPIResponse)
}

// ReleaseTaobaoSungariInspectionSubmitAPIResponse 将 TaobaoSungariInspectionSubmitAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSungariInspectionSubmitAPIResponse(v *TaobaoSungariInspectionSubmitAPIResponse) {
	v.Reset()
	poolTaobaoSungariInspectionSubmitAPIResponse.Put(v)
}
