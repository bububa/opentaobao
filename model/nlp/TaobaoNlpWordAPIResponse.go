package nlp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoNlpWordAPIResponse 文本语言词法分析 API返回值
// taobao.nlp.word
//
// 提供文本语言处理中的词法分析功能,开放中文分词和词权重计算功能。
type TaobaoNlpWordAPIResponse struct {
	model.CommonResponse
	TaobaoNlpWordAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoNlpWordAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoNlpWordAPIResponseModel).Reset()
}

// TaobaoNlpWordAPIResponseModel is 文本语言词法分析 成功返回结果
type TaobaoNlpWordAPIResponseModel struct {
	XMLName xml.Name `xml:"nlp_word_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回词法分析的结果
	Wordresult *WordResult `json:"wordresult,omitempty" xml:"wordresult,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoNlpWordAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Wordresult = nil
}

var poolTaobaoNlpWordAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoNlpWordAPIResponse)
	},
}

// GetTaobaoNlpWordAPIResponse 从 sync.Pool 获取 TaobaoNlpWordAPIResponse
func GetTaobaoNlpWordAPIResponse() *TaobaoNlpWordAPIResponse {
	return poolTaobaoNlpWordAPIResponse.Get().(*TaobaoNlpWordAPIResponse)
}

// ReleaseTaobaoNlpWordAPIResponse 将 TaobaoNlpWordAPIResponse 保存到 sync.Pool
func ReleaseTaobaoNlpWordAPIResponse(v *TaobaoNlpWordAPIResponse) {
	v.Reset()
	poolTaobaoNlpWordAPIResponse.Put(v)
}
