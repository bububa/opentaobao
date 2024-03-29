package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaInsightCatstopwordnewGetAPIResponse 获取类目下最热门的词 API返回值
// taobao.simba.insight.catstopwordnew.get
//
// 按照某个维度，查询某个类目下最热门的词，维度有点击，展现，花费，点击率等，具体可以按哪些字段进行排序，参考参数说明，比如选择了impression，则返回该类目下展现量最高那几个词。
type TaobaoSimbaInsightCatstopwordnewGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaInsightCatstopwordnewGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaInsightCatstopwordnewGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaInsightCatstopwordnewGetAPIResponseModel).Reset()
}

// TaobaoSimbaInsightCatstopwordnewGetAPIResponseModel is 获取类目下最热门的词 成功返回结果
type TaobaoSimbaInsightCatstopwordnewGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_insight_catstopwordnew_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 类目下热门词详细数据
	TopwordDataList []InsightWordDataUnderCatDto `json:"topword_data_list,omitempty" xml:"topword_data_list>insight_word_data_under_cat_dto,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaInsightCatstopwordnewGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TopwordDataList = m.TopwordDataList[:0]
}

var poolTaobaoSimbaInsightCatstopwordnewGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaInsightCatstopwordnewGetAPIResponse)
	},
}

// GetTaobaoSimbaInsightCatstopwordnewGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaInsightCatstopwordnewGetAPIResponse
func GetTaobaoSimbaInsightCatstopwordnewGetAPIResponse() *TaobaoSimbaInsightCatstopwordnewGetAPIResponse {
	return poolTaobaoSimbaInsightCatstopwordnewGetAPIResponse.Get().(*TaobaoSimbaInsightCatstopwordnewGetAPIResponse)
}

// ReleaseTaobaoSimbaInsightCatstopwordnewGetAPIResponse 将 TaobaoSimbaInsightCatstopwordnewGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaInsightCatstopwordnewGetAPIResponse(v *TaobaoSimbaInsightCatstopwordnewGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaInsightCatstopwordnewGetAPIResponse.Put(v)
}
