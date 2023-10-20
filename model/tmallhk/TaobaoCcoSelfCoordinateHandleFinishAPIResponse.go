package tmallhk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCcoSelfCoordinateHandleFinishAPIResponse 天猫国际直购供应商处理完结回复通知 API返回值
// taobao.cco.self.coordinate.handle.finish
//
// 天猫国际直购供应商处理完结回复通知
type TaobaoCcoSelfCoordinateHandleFinishAPIResponse struct {
	model.CommonResponse
	TaobaoCcoSelfCoordinateHandleFinishAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCcoSelfCoordinateHandleFinishAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCcoSelfCoordinateHandleFinishAPIResponseModel).Reset()
}

// TaobaoCcoSelfCoordinateHandleFinishAPIResponseModel is 天猫国际直购供应商处理完结回复通知 成功返回结果
type TaobaoCcoSelfCoordinateHandleFinishAPIResponseModel struct {
	XMLName xml.Name `xml:"cco_self_coordinate_handle_finish_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误码信息
	ApiCode string `json:"api_code,omitempty" xml:"api_code,omitempty"`
	// 错误信息
	ApiMessage string `json:"api_message,omitempty" xml:"api_message,omitempty"`
	// 处理结果
	Data bool `json:"data,omitempty" xml:"data,omitempty"`
	// 接口是否调用成功
	ApiSuccess bool `json:"api_success,omitempty" xml:"api_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCcoSelfCoordinateHandleFinishAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiCode = ""
	m.ApiMessage = ""
	m.Data = false
	m.ApiSuccess = false
}

var poolTaobaoCcoSelfCoordinateHandleFinishAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCcoSelfCoordinateHandleFinishAPIResponse)
	},
}

// GetTaobaoCcoSelfCoordinateHandleFinishAPIResponse 从 sync.Pool 获取 TaobaoCcoSelfCoordinateHandleFinishAPIResponse
func GetTaobaoCcoSelfCoordinateHandleFinishAPIResponse() *TaobaoCcoSelfCoordinateHandleFinishAPIResponse {
	return poolTaobaoCcoSelfCoordinateHandleFinishAPIResponse.Get().(*TaobaoCcoSelfCoordinateHandleFinishAPIResponse)
}

// ReleaseTaobaoCcoSelfCoordinateHandleFinishAPIResponse 将 TaobaoCcoSelfCoordinateHandleFinishAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCcoSelfCoordinateHandleFinishAPIResponse(v *TaobaoCcoSelfCoordinateHandleFinishAPIResponse) {
	v.Reset()
	poolTaobaoCcoSelfCoordinateHandleFinishAPIResponse.Put(v)
}
