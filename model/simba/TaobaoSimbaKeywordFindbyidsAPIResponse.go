package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordFindbyidsAPIResponse （新）根据一堆关键词ids获取关键词 API返回值
// taobao.simba.keyword.findbyids
//
// 根据一个关键词Id列表取得一组关键词
type TaobaoSimbaKeywordFindbyidsAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaKeywordFindbyidsAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaKeywordFindbyidsAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaKeywordFindbyidsAPIResponseModel).Reset()
}

// TaobaoSimbaKeywordFindbyidsAPIResponseModel is （新）根据一堆关键词ids获取关键词 成功返回结果
type TaobaoSimbaKeywordFindbyidsAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_keyword_findbyids_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 整体的返回值
	Results []SiriusBidwordDto `json:"results,omitempty" xml:"results>sirius_bidword_dto,omitempty"`
	// 错误原因
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaKeywordFindbyidsAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
	m.ErrorMsg = ""
}

var poolTaobaoSimbaKeywordFindbyidsAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaKeywordFindbyidsAPIResponse)
	},
}

// GetTaobaoSimbaKeywordFindbyidsAPIResponse 从 sync.Pool 获取 TaobaoSimbaKeywordFindbyidsAPIResponse
func GetTaobaoSimbaKeywordFindbyidsAPIResponse() *TaobaoSimbaKeywordFindbyidsAPIResponse {
	return poolTaobaoSimbaKeywordFindbyidsAPIResponse.Get().(*TaobaoSimbaKeywordFindbyidsAPIResponse)
}

// ReleaseTaobaoSimbaKeywordFindbyidsAPIResponse 将 TaobaoSimbaKeywordFindbyidsAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaKeywordFindbyidsAPIResponse(v *TaobaoSimbaKeywordFindbyidsAPIResponse) {
	v.Reset()
	poolTaobaoSimbaKeywordFindbyidsAPIResponse.Put(v)
}
