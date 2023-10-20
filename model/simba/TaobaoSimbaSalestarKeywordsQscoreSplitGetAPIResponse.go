package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponse (新)销量明星质量分相关接口 API返回值
// taobao.simba.salestar.keywords.qscore.split.get
//
// 获取关键词新的质量分
type TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponseModel).Reset()
}

// TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponseModel is (新)销量明星质量分相关接口 成功返回结果
type TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_salestar_keywords_qscore_split_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// result
	Result *TaobaoSimbaSalestarKeywordsQscoreSplitGetResultDto `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponse)
	},
}

// GetTaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponse
func GetTaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponse() *TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponse {
	return poolTaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponse.Get().(*TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponse)
}

// ReleaseTaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponse 将 TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponse(v *TaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaSalestarKeywordsQscoreSplitGetAPIResponse.Put(v)
}
