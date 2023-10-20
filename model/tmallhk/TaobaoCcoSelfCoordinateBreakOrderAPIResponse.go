package tmallhk

import (
	"encoding/xml"

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
