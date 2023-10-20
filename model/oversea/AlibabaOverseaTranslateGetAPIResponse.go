package oversea

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaOverseaTranslateGetAPIResponse 获取文本翻译信息 API返回值
// alibaba.oversea.translate.get
//
// 根据传入的文本信息，获取其目标语言的翻译结果
type AlibabaOverseaTranslateGetAPIResponse struct {
	model.CommonResponse
	AlibabaOverseaTranslateGetAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaOverseaTranslateGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaOverseaTranslateGetAPIResponseModel).Reset()
}

// AlibabaOverseaTranslateGetAPIResponseModel is 获取文本翻译信息 成功返回结果
type AlibabaOverseaTranslateGetAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_oversea_translate_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *SimpleTransResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaOverseaTranslateGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlibabaOverseaTranslateGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaOverseaTranslateGetAPIResponse)
	},
}

// GetAlibabaOverseaTranslateGetAPIResponse 从 sync.Pool 获取 AlibabaOverseaTranslateGetAPIResponse
func GetAlibabaOverseaTranslateGetAPIResponse() *AlibabaOverseaTranslateGetAPIResponse {
	return poolAlibabaOverseaTranslateGetAPIResponse.Get().(*AlibabaOverseaTranslateGetAPIResponse)
}

// ReleaseAlibabaOverseaTranslateGetAPIResponse 将 AlibabaOverseaTranslateGetAPIResponse 保存到 sync.Pool
func ReleaseAlibabaOverseaTranslateGetAPIResponse(v *AlibabaOverseaTranslateGetAPIResponse) {
	v.Reset()
	poolAlibabaOverseaTranslateGetAPIResponse.Put(v)
}
