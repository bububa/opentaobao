package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordsvonAddAPIResponse 创建一批关键词 API返回值
// taobao.simba.keywordsvon.add
//
// 创建一批关键词
type TaobaoSimbaKeywordsvonAddAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaKeywordsvonAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaKeywordsvonAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaKeywordsvonAddAPIResponseModel).Reset()
}

// TaobaoSimbaKeywordsvonAddAPIResponseModel is 创建一批关键词 成功返回结果
type TaobaoSimbaKeywordsvonAddAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_keywordsvon_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 关键词列表
	Keywords []Keyword `json:"keywords,omitempty" xml:"keywords>keyword,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaKeywordsvonAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Keywords = m.Keywords[:0]
}

var poolTaobaoSimbaKeywordsvonAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaKeywordsvonAddAPIResponse)
	},
}

// GetTaobaoSimbaKeywordsvonAddAPIResponse 从 sync.Pool 获取 TaobaoSimbaKeywordsvonAddAPIResponse
func GetTaobaoSimbaKeywordsvonAddAPIResponse() *TaobaoSimbaKeywordsvonAddAPIResponse {
	return poolTaobaoSimbaKeywordsvonAddAPIResponse.Get().(*TaobaoSimbaKeywordsvonAddAPIResponse)
}

// ReleaseTaobaoSimbaKeywordsvonAddAPIResponse 将 TaobaoSimbaKeywordsvonAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaKeywordsvonAddAPIResponse(v *TaobaoSimbaKeywordsvonAddAPIResponse) {
	v.Reset()
	poolTaobaoSimbaKeywordsvonAddAPIResponse.Put(v)
}
