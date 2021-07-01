package sungari

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoSungariInspectionSubmitAPIResponse
抽检指令录入 API返回值
taobao.sungari.inspection.submit

抽检指令录入 */
type TaobaoSungariInspectionSubmitAPIResponse struct {
	model.CommonResponse
	TaobaoSungariInspectionSubmitAPIResponseModel
}

// TaobaoSungariInspectionSubmitAPIResponseModel is 抽检指令录入 成功返回结果
type TaobaoSungariInspectionSubmitAPIResponseModel struct {
	XMLName xml.Name `xml:"sungari_inspection_submit_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	Data *InspectionResultInfo `json:"data,omitempty" xml:"data,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// 状态码
	ReturnCode int64 `json:"return_code,omitempty" xml:"return_code,omitempty"`
}
