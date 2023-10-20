package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordsPricevonSetAPIResponse 设置一批关键词的信息 API返回值
// taobao.simba.keywords.pricevon.set
//
// 设置一批关键词的信息，包含无线出价、计算机出价和关键词匹配方式
type TaobaoSimbaKeywordsPricevonSetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaKeywordsPricevonSetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaKeywordsPricevonSetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaKeywordsPricevonSetAPIResponseModel).Reset()
}

// TaobaoSimbaKeywordsPricevonSetAPIResponseModel is 设置一批关键词的信息 成功返回结果
type TaobaoSimbaKeywordsPricevonSetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_keywords_pricevon_set_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 成功设置关键词价格的关键词列表
	Keywords []Keyword `json:"keywords,omitempty" xml:"keywords>keyword,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaKeywordsPricevonSetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Keywords = m.Keywords[:0]
}

var poolTaobaoSimbaKeywordsPricevonSetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaKeywordsPricevonSetAPIResponse)
	},
}

// GetTaobaoSimbaKeywordsPricevonSetAPIResponse 从 sync.Pool 获取 TaobaoSimbaKeywordsPricevonSetAPIResponse
func GetTaobaoSimbaKeywordsPricevonSetAPIResponse() *TaobaoSimbaKeywordsPricevonSetAPIResponse {
	return poolTaobaoSimbaKeywordsPricevonSetAPIResponse.Get().(*TaobaoSimbaKeywordsPricevonSetAPIResponse)
}

// ReleaseTaobaoSimbaKeywordsPricevonSetAPIResponse 将 TaobaoSimbaKeywordsPricevonSetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaKeywordsPricevonSetAPIResponse(v *TaobaoSimbaKeywordsPricevonSetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaKeywordsPricevonSetAPIResponse.Put(v)
}
