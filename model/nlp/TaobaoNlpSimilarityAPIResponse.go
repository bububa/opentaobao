package nlp

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoNlpSimilarityAPIResponse 文本语言相似度 API返回值
// taobao.nlp.similarity
//
// 文本语言相似度计算，提供余弦距离、编辑距离和simHash三种相似度计算。返回文本相似度区间为0-1之间，0为完全不相似，1为完全相似。
type TaobaoNlpSimilarityAPIResponse struct {
	model.CommonResponse
	TaobaoNlpSimilarityAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoNlpSimilarityAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoNlpSimilarityAPIResponseModel).Reset()
}

// TaobaoNlpSimilarityAPIResponseModel is 文本语言相似度 成功返回结果
type TaobaoNlpSimilarityAPIResponseModel struct {
	XMLName xml.Name `xml:"nlp_similarity_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回结果
	Simresult *SimResult `json:"simresult,omitempty" xml:"simresult,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoNlpSimilarityAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Simresult = nil
}

var poolTaobaoNlpSimilarityAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoNlpSimilarityAPIResponse)
	},
}

// GetTaobaoNlpSimilarityAPIResponse 从 sync.Pool 获取 TaobaoNlpSimilarityAPIResponse
func GetTaobaoNlpSimilarityAPIResponse() *TaobaoNlpSimilarityAPIResponse {
	return poolTaobaoNlpSimilarityAPIResponse.Get().(*TaobaoNlpSimilarityAPIResponse)
}

// ReleaseTaobaoNlpSimilarityAPIResponse 将 TaobaoNlpSimilarityAPIResponse 保存到 sync.Pool
func ReleaseTaobaoNlpSimilarityAPIResponse(v *TaobaoNlpSimilarityAPIResponse) {
	v.Reset()
	poolTaobaoNlpSimilarityAPIResponse.Put(v)
}
