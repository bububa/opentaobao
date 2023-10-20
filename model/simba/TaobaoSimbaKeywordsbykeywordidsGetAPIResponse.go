package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordsbykeywordidsGetAPIResponse 根据一个关键词Id列表取得一组关键词 API返回值
// taobao.simba.keywordsbykeywordids.get
//
// 根据一个关键词Id列表取得一组关键词
type TaobaoSimbaKeywordsbykeywordidsGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaKeywordsbykeywordidsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaKeywordsbykeywordidsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaKeywordsbykeywordidsGetAPIResponseModel).Reset()
}

// TaobaoSimbaKeywordsbykeywordidsGetAPIResponseModel is 根据一个关键词Id列表取得一组关键词 成功返回结果
type TaobaoSimbaKeywordsbykeywordidsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_keywordsbykeywordids_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 取得的关键词列表
	Keywords []Keyword `json:"keywords,omitempty" xml:"keywords>keyword,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaKeywordsbykeywordidsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Keywords = m.Keywords[:0]
}

var poolTaobaoSimbaKeywordsbykeywordidsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaKeywordsbykeywordidsGetAPIResponse)
	},
}

// GetTaobaoSimbaKeywordsbykeywordidsGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaKeywordsbykeywordidsGetAPIResponse
func GetTaobaoSimbaKeywordsbykeywordidsGetAPIResponse() *TaobaoSimbaKeywordsbykeywordidsGetAPIResponse {
	return poolTaobaoSimbaKeywordsbykeywordidsGetAPIResponse.Get().(*TaobaoSimbaKeywordsbykeywordidsGetAPIResponse)
}

// ReleaseTaobaoSimbaKeywordsbykeywordidsGetAPIResponse 将 TaobaoSimbaKeywordsbykeywordidsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaKeywordsbykeywordidsGetAPIResponse(v *TaobaoSimbaKeywordsbykeywordidsGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaKeywordsbykeywordidsGetAPIResponse.Put(v)
}
