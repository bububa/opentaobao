package nlp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoNlpPreprocessAPIResponse 文本语言预处理 API返回值
// taobao.nlp.preprocess
//
// 实现文本语言处理中的预处理功能，如实现文字繁简转换、文字转拼音、文字拆分、谐音同音字替换和形似字替换。
type TaobaoNlpPreprocessAPIResponse struct {
	model.CommonResponse
	TaobaoNlpPreprocessAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoNlpPreprocessAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoNlpPreprocessAPIResponseModel).Reset()
}

// TaobaoNlpPreprocessAPIResponseModel is 文本语言预处理 成功返回结果
type TaobaoNlpPreprocessAPIResponseModel struct {
	XMLName xml.Name `xml:"nlp_preprocess_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Processresult *ProcessResult `json:"processresult,omitempty" xml:"processresult,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoNlpPreprocessAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Processresult = nil
}

var poolTaobaoNlpPreprocessAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoNlpPreprocessAPIResponse)
	},
}

// GetTaobaoNlpPreprocessAPIResponse 从 sync.Pool 获取 TaobaoNlpPreprocessAPIResponse
func GetTaobaoNlpPreprocessAPIResponse() *TaobaoNlpPreprocessAPIResponse {
	return poolTaobaoNlpPreprocessAPIResponse.Get().(*TaobaoNlpPreprocessAPIResponse)
}

// ReleaseTaobaoNlpPreprocessAPIResponse 将 TaobaoNlpPreprocessAPIResponse 保存到 sync.Pool
func ReleaseTaobaoNlpPreprocessAPIResponse(v *TaobaoNlpPreprocessAPIResponse) {
	v.Reset()
	poolTaobaoNlpPreprocessAPIResponse.Put(v)
}
