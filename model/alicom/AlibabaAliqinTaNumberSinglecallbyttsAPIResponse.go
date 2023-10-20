package alicom

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaAliqinTaNumberSinglecallbyttsAPIResponse 根据号码tts单呼 API返回值
// alibaba.aliqin.ta.number.singlecallbytts
//
// 将语音验证码和语音通知发布至聚石塔渠道
type AlibabaAliqinTaNumberSinglecallbyttsAPIResponse struct {
	model.CommonResponse
	AlibabaAliqinTaNumberSinglecallbyttsAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaAliqinTaNumberSinglecallbyttsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaAliqinTaNumberSinglecallbyttsAPIResponseModel).Reset()
}

// AlibabaAliqinTaNumberSinglecallbyttsAPIResponseModel is 根据号码tts单呼 成功返回结果
type AlibabaAliqinTaNumberSinglecallbyttsAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_aliqin_ta_number_singlecallbytts_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlibabaAliqinTaNumberSinglecallbyttsResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaAliqinTaNumberSinglecallbyttsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaAliqinTaNumberSinglecallbyttsAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaAliqinTaNumberSinglecallbyttsAPIResponse)
	},
}

// GetAlibabaAliqinTaNumberSinglecallbyttsAPIResponse 从 sync.Pool 获取 AlibabaAliqinTaNumberSinglecallbyttsAPIResponse
func GetAlibabaAliqinTaNumberSinglecallbyttsAPIResponse() *AlibabaAliqinTaNumberSinglecallbyttsAPIResponse {
	return poolAlibabaAliqinTaNumberSinglecallbyttsAPIResponse.Get().(*AlibabaAliqinTaNumberSinglecallbyttsAPIResponse)
}

// ReleaseAlibabaAliqinTaNumberSinglecallbyttsAPIResponse 将 AlibabaAliqinTaNumberSinglecallbyttsAPIResponse 保存到 sync.Pool
func ReleaseAlibabaAliqinTaNumberSinglecallbyttsAPIResponse(v *AlibabaAliqinTaNumberSinglecallbyttsAPIResponse) {
	v.Reset()
	poolAlibabaAliqinTaNumberSinglecallbyttsAPIResponse.Put(v)
}
