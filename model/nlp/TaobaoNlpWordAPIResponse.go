package nlp

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaonlpwordAPIResponse 文本语言词法分析 API返回值
// taobao.nlp.word
//
// 提供文本语言处理中的词法分析功能,开放中文分词和词权重计算功能。
type TaobaonlpwordAPIResponse struct {
	model.CommonResponse
	TaobaonlpwordAPIResponseModel
}

// TaobaonlpwordAPIResponseModel is 文本语言词法分析 成功返回结果
type TaobaonlpwordAPIResponseModel struct {
	XMLName xml.Name `xml:"nlp_word_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回词法分析的结果
	Wordresult *WordResult `json:"wordresult,omitempty" xml:"wordresult,omitempty"`
}
