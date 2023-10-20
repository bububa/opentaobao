package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordsQscoreSplitGetAPIResponse 新质量分服务 API返回值
// taobao.simba.keywords.qscore.split.get
//
// 获取关键词新的质量分
type TaobaoSimbaKeywordsQscoreSplitGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaKeywordsQscoreSplitGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaKeywordsQscoreSplitGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaKeywordsQscoreSplitGetAPIResponseModel).Reset()
}

// TaobaoSimbaKeywordsQscoreSplitGetAPIResponseModel is 新质量分服务 成功返回结果
type TaobaoSimbaKeywordsQscoreSplitGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_keywords_qscore_split_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoSimbaKeywordsQscoreSplitGetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaKeywordsQscoreSplitGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoSimbaKeywordsQscoreSplitGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaKeywordsQscoreSplitGetAPIResponse)
	},
}

// GetTaobaoSimbaKeywordsQscoreSplitGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaKeywordsQscoreSplitGetAPIResponse
func GetTaobaoSimbaKeywordsQscoreSplitGetAPIResponse() *TaobaoSimbaKeywordsQscoreSplitGetAPIResponse {
	return poolTaobaoSimbaKeywordsQscoreSplitGetAPIResponse.Get().(*TaobaoSimbaKeywordsQscoreSplitGetAPIResponse)
}

// ReleaseTaobaoSimbaKeywordsQscoreSplitGetAPIResponse 将 TaobaoSimbaKeywordsQscoreSplitGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaKeywordsQscoreSplitGetAPIResponse(v *TaobaoSimbaKeywordsQscoreSplitGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaKeywordsQscoreSplitGetAPIResponse.Put(v)
}
