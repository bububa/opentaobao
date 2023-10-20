package util

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoKfcKeywordSearchAPIResponse 关键词过滤匹配 API返回值
// taobao.kfc.keyword.search
//
// 对输入的文本信息进行禁忌关键词匹配，返回匹配的结果
type TaobaoKfcKeywordSearchAPIResponse struct {
	model.CommonResponse
	TaobaoKfcKeywordSearchAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoKfcKeywordSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoKfcKeywordSearchAPIResponseModel).Reset()
}

// TaobaoKfcKeywordSearchAPIResponseModel is 关键词过滤匹配 成功返回结果
type TaobaoKfcKeywordSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"kfc_keyword_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// KFC 关键词过滤匹配结果
	KfcSearchResult *KfcSearchResult `json:"kfc_search_result,omitempty" xml:"kfc_search_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoKfcKeywordSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.KfcSearchResult = nil
}

var poolTaobaoKfcKeywordSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoKfcKeywordSearchAPIResponse)
	},
}

// GetTaobaoKfcKeywordSearchAPIResponse 从 sync.Pool 获取 TaobaoKfcKeywordSearchAPIResponse
func GetTaobaoKfcKeywordSearchAPIResponse() *TaobaoKfcKeywordSearchAPIResponse {
	return poolTaobaoKfcKeywordSearchAPIResponse.Get().(*TaobaoKfcKeywordSearchAPIResponse)
}

// ReleaseTaobaoKfcKeywordSearchAPIResponse 将 TaobaoKfcKeywordSearchAPIResponse 保存到 sync.Pool
func ReleaseTaobaoKfcKeywordSearchAPIResponse(v *TaobaoKfcKeywordSearchAPIResponse) {
	v.Reset()
	poolTaobaoKfcKeywordSearchAPIResponse.Put(v)
}
