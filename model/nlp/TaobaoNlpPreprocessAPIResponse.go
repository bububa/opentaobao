package nlp

import (
	"encoding/xml"

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

// TaobaoNlpPreprocessAPIResponseModel is 文本语言预处理 成功返回结果
type TaobaoNlpPreprocessAPIResponseModel struct {
	XMLName xml.Name `xml:"nlp_preprocess_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Processresult *ProcessResult `json:"processresult,omitempty" xml:"processresult,omitempty"`
}
