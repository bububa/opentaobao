package alicom

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaaliqintanumbersinglecallbyvoiceAPIResponse 根据号码tts单呼 API返回值
// alibaba.aliqin.ta.number.singlecallbyvoice
//
// 根据号码语音单呼
type AlibabaaliqintanumbersinglecallbyvoiceAPIResponse struct {
	model.CommonResponse
	AlibabaaliqintanumbersinglecallbyvoiceAPIResponseModel
}

// AlibabaaliqintanumbersinglecallbyvoiceAPIResponseModel is 根据号码tts单呼 成功返回结果
type AlibabaaliqintanumbersinglecallbyvoiceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_ta_number_singlecallbyvoice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaaliqintanumbersinglecallbyvoiceResult `json:"result,omitempty" xml:"result,omitempty"`
}
