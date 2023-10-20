package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaInsightWordsdataGetAPIResponse 获取关键词的大盘数据 API返回值
// taobao.simba.insight.wordsdata.get
//
// 获取关键词的详细数据
type TaobaoSimbaInsightWordsdataGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaInsightWordsdataGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaInsightWordsdataGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaInsightWordsdataGetAPIResponseModel).Reset()
}

// TaobaoSimbaInsightWordsdataGetAPIResponseModel is 获取关键词的大盘数据 成功返回结果
type TaobaoSimbaInsightWordsdataGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_insight_wordsdata_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 关键词大盘数据列表
	WordDataList []InsightWordDataDto `json:"word_data_list,omitempty" xml:"word_data_list>insight_word_data_dto,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaInsightWordsdataGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.WordDataList = m.WordDataList[:0]
}

var poolTaobaoSimbaInsightWordsdataGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaInsightWordsdataGetAPIResponse)
	},
}

// GetTaobaoSimbaInsightWordsdataGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaInsightWordsdataGetAPIResponse
func GetTaobaoSimbaInsightWordsdataGetAPIResponse() *TaobaoSimbaInsightWordsdataGetAPIResponse {
	return poolTaobaoSimbaInsightWordsdataGetAPIResponse.Get().(*TaobaoSimbaInsightWordsdataGetAPIResponse)
}

// ReleaseTaobaoSimbaInsightWordsdataGetAPIResponse 将 TaobaoSimbaInsightWordsdataGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaInsightWordsdataGetAPIResponse(v *TaobaoSimbaInsightWordsdataGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaInsightWordsdataGetAPIResponse.Put(v)
}
