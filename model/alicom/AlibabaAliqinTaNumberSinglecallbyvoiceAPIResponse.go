package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinTaNumberSinglecallbyvoiceAPIResponse 根据号码tts单呼 API返回值
// alibaba.aliqin.ta.number.singlecallbyvoice
//
// 根据号码语音单呼
type AlibabaAliqinTaNumberSinglecallbyvoiceAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinTaNumberSinglecallbyvoiceAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinTaNumberSinglecallbyvoiceAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinTaNumberSinglecallbyvoiceAPIResponseModel).Reset()
}

// AlibabaAliqinTaNumberSinglecallbyvoiceAPIResponseModel is 根据号码tts单呼 成功返回结果
type AlibabaAliqinTaNumberSinglecallbyvoiceAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_ta_number_singlecallbyvoice_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAliqinTaNumberSinglecallbyvoiceResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinTaNumberSinglecallbyvoiceAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAliqinTaNumberSinglecallbyvoiceAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinTaNumberSinglecallbyvoiceAPIResponse)
	},
}

// GetAlibabaAliqinTaNumberSinglecallbyvoiceAPIResponse 从 sync.Pool 获取 AlibabaAliqinTaNumberSinglecallbyvoiceAPIResponse
func GetAlibabaAliqinTaNumberSinglecallbyvoiceAPIResponse() *AlibabaAliqinTaNumberSinglecallbyvoiceAPIResponse {
	return poolAlibabaAliqinTaNumberSinglecallbyvoiceAPIResponse.Get().(*AlibabaAliqinTaNumberSinglecallbyvoiceAPIResponse)
}

// ReleaseAlibabaAliqinTaNumberSinglecallbyvoiceAPIResponse 将 AlibabaAliqinTaNumberSinglecallbyvoiceAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinTaNumberSinglecallbyvoiceAPIResponse(v *AlibabaAliqinTaNumberSinglecallbyvoiceAPIResponse) {
	v.Reset()
	poolAlibabaAliqinTaNumberSinglecallbyvoiceAPIResponse.Put(v)
}
