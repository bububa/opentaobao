package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordsQscoreGetAPIResponse 取得一个推广组的所有关键词的质量得分或者根据关键词Id列表取得一组关键词的质量得分 API返回值
// taobao.simba.keywords.qscore.get
//
// 取得一个推广组的所有关键词的质量得分列表
type TaobaoSimbaKeywordsQscoreGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaKeywordsQscoreGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaKeywordsQscoreGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaKeywordsQscoreGetAPIResponseModel).Reset()
}

// TaobaoSimbaKeywordsQscoreGetAPIResponseModel is 取得一个推广组的所有关键词的质量得分或者根据关键词Id列表取得一组关键词的质量得分 成功返回结果
type TaobaoSimbaKeywordsQscoreGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_keywords_qscore_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 取得的关键词质量得分列表
	KeywordQscoreList []KeywordQscore `json:"keyword_qscore_list,omitempty" xml:"keyword_qscore_list>keyword_qscore,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaKeywordsQscoreGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.KeywordQscoreList = m.KeywordQscoreList[:0]
}

var poolTaobaoSimbaKeywordsQscoreGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaKeywordsQscoreGetAPIResponse)
	},
}

// GetTaobaoSimbaKeywordsQscoreGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaKeywordsQscoreGetAPIResponse
func GetTaobaoSimbaKeywordsQscoreGetAPIResponse() *TaobaoSimbaKeywordsQscoreGetAPIResponse {
	return poolTaobaoSimbaKeywordsQscoreGetAPIResponse.Get().(*TaobaoSimbaKeywordsQscoreGetAPIResponse)
}

// ReleaseTaobaoSimbaKeywordsQscoreGetAPIResponse 将 TaobaoSimbaKeywordsQscoreGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaKeywordsQscoreGetAPIResponse(v *TaobaoSimbaKeywordsQscoreGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaKeywordsQscoreGetAPIResponse.Put(v)
}
