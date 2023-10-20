package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSalestarKeywordsRecommendGetAPIResponse 销量明星api相关接口 API返回值
// taobao.simba.salestar.keywords.recommend.get
//
// 取得一个推广组的推荐关键词列表
type TaobaoSimbaSalestarKeywordsRecommendGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaSalestarKeywordsRecommendGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaSalestarKeywordsRecommendGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaSalestarKeywordsRecommendGetAPIResponseModel).Reset()
}

// TaobaoSimbaSalestarKeywordsRecommendGetAPIResponseModel is 销量明星api相关接口 成功返回结果
type TaobaoSimbaSalestarKeywordsRecommendGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_salestar_keywords_recommend_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推荐词分页对象，当输入的页码大于最大数值时，将返回最大的page_no值，并且结果中的数据列表为空值
	RecommendWords *RecommendWordPage `json:"recommend_words,omitempty" xml:"recommend_words,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaSalestarKeywordsRecommendGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.RecommendWords = nil
}

var poolTaobaoSimbaSalestarKeywordsRecommendGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaSalestarKeywordsRecommendGetAPIResponse)
	},
}

// GetTaobaoSimbaSalestarKeywordsRecommendGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaSalestarKeywordsRecommendGetAPIResponse
func GetTaobaoSimbaSalestarKeywordsRecommendGetAPIResponse() *TaobaoSimbaSalestarKeywordsRecommendGetAPIResponse {
	return poolTaobaoSimbaSalestarKeywordsRecommendGetAPIResponse.Get().(*TaobaoSimbaSalestarKeywordsRecommendGetAPIResponse)
}

// ReleaseTaobaoSimbaSalestarKeywordsRecommendGetAPIResponse 将 TaobaoSimbaSalestarKeywordsRecommendGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaSalestarKeywordsRecommendGetAPIResponse(v *TaobaoSimbaSalestarKeywordsRecommendGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaSalestarKeywordsRecommendGetAPIResponse.Put(v)
}
