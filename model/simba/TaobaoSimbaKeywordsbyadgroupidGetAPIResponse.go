package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordsbyadgroupidGetAPIResponse 取得一个推广组的所有关键词 API返回值
// taobao.simba.keywordsbyadgroupid.get
//
// 取得一个推广组的所有关键词
type TaobaoSimbaKeywordsbyadgroupidGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaKeywordsbyadgroupidGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaKeywordsbyadgroupidGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaKeywordsbyadgroupidGetAPIResponseModel).Reset()
}

// TaobaoSimbaKeywordsbyadgroupidGetAPIResponseModel is 取得一个推广组的所有关键词 成功返回结果
type TaobaoSimbaKeywordsbyadgroupidGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_keywordsbyadgroupid_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 取得的关键词列表
	Keywords []Keyword `json:"keywords,omitempty" xml:"keywords>keyword,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaKeywordsbyadgroupidGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Keywords = m.Keywords[:0]
}

var poolTaobaoSimbaKeywordsbyadgroupidGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaKeywordsbyadgroupidGetAPIResponse)
	},
}

// GetTaobaoSimbaKeywordsbyadgroupidGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaKeywordsbyadgroupidGetAPIResponse
func GetTaobaoSimbaKeywordsbyadgroupidGetAPIResponse() *TaobaoSimbaKeywordsbyadgroupidGetAPIResponse {
	return poolTaobaoSimbaKeywordsbyadgroupidGetAPIResponse.Get().(*TaobaoSimbaKeywordsbyadgroupidGetAPIResponse)
}

// ReleaseTaobaoSimbaKeywordsbyadgroupidGetAPIResponse 将 TaobaoSimbaKeywordsbyadgroupidGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaKeywordsbyadgroupidGetAPIResponse(v *TaobaoSimbaKeywordsbyadgroupidGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaKeywordsbyadgroupidGetAPIResponse.Put(v)
}
