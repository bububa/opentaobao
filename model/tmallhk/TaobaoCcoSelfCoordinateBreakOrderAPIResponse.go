package tmallhk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoCcoSelfCoordinateBreakOrderAPIResponse 天猫国际直购供应商毁单通知 API返回值
// taobao.cco.self.coordinate.break.order
//
// 天猫国际直购供应商毁单通知
type TaobaoCcoSelfCoordinateBreakOrderAPIResponse struct {
	model.CommonResponse
	TaobaoCcoSelfCoordinateBreakOrderAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoCcoSelfCoordinateBreakOrderAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoCcoSelfCoordinateBreakOrderAPIResponseModel).Reset()
}

// TaobaoCcoSelfCoordinateBreakOrderAPIResponseModel is 天猫国际直购供应商毁单通知 成功返回结果
type TaobaoCcoSelfCoordinateBreakOrderAPIResponseModel struct {
	XMLName xml.Name `xml:"cco_self_coordinate_break_order_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// api错误码
	ApiCode string `json:"api_code,omitempty" xml:"api_code,omitempty"`
	// api错误信息
	ApiMessage string `json:"api_message,omitempty" xml:"api_message,omitempty"`
	// api调用结果
	ApiData bool `json:"api_data,omitempty" xml:"api_data,omitempty"`
	// api调用是否成功
	ApiSuccess bool `json:"api_success,omitempty" xml:"api_success,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoCcoSelfCoordinateBreakOrderAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ApiCode = ""
	m.ApiMessage = ""
	m.ApiData = false
	m.ApiSuccess = false
}

var poolTaobaoCcoSelfCoordinateBreakOrderAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoCcoSelfCoordinateBreakOrderAPIResponse)
	},
}

// GetTaobaoCcoSelfCoordinateBreakOrderAPIResponse 从 sync.Pool 获取 TaobaoCcoSelfCoordinateBreakOrderAPIResponse
func GetTaobaoCcoSelfCoordinateBreakOrderAPIResponse() *TaobaoCcoSelfCoordinateBreakOrderAPIResponse {
	return poolTaobaoCcoSelfCoordinateBreakOrderAPIResponse.Get().(*TaobaoCcoSelfCoordinateBreakOrderAPIResponse)
}

// ReleaseTaobaoCcoSelfCoordinateBreakOrderAPIResponse 将 TaobaoCcoSelfCoordinateBreakOrderAPIResponse 保存到 sync.Pool
func ReleaseTaobaoCcoSelfCoordinateBreakOrderAPIResponse(v *TaobaoCcoSelfCoordinateBreakOrderAPIResponse) {
	v.Reset()
	poolTaobaoCcoSelfCoordinateBreakOrderAPIResponse.Put(v)
}
