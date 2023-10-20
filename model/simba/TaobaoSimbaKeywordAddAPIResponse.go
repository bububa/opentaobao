package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordAddAPIResponse （新）关键词新增接口 API返回值
// taobao.simba.keyword.add
//
// （新）关键词更新相关接口
type TaobaoSimbaKeywordAddAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaKeywordAddAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaKeywordAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaKeywordAddAPIResponseModel).Reset()
}

// TaobaoSimbaKeywordAddAPIResponseModel is （新）关键词新增接口 成功返回结果
type TaobaoSimbaKeywordAddAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_keyword_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 整体的返回值
	Results []SiriusBidwordDto `json:"results,omitempty" xml:"results>sirius_bidword_dto,omitempty"`
	// 错误原因
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaKeywordAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
	m.ErrorMsg = ""
}

var poolTaobaoSimbaKeywordAddAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaKeywordAddAPIResponse)
	},
}

// GetTaobaoSimbaKeywordAddAPIResponse 从 sync.Pool 获取 TaobaoSimbaKeywordAddAPIResponse
func GetTaobaoSimbaKeywordAddAPIResponse() *TaobaoSimbaKeywordAddAPIResponse {
	return poolTaobaoSimbaKeywordAddAPIResponse.Get().(*TaobaoSimbaKeywordAddAPIResponse)
}

// ReleaseTaobaoSimbaKeywordAddAPIResponse 将 TaobaoSimbaKeywordAddAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaKeywordAddAPIResponse(v *TaobaoSimbaKeywordAddAPIResponse) {
	v.Reset()
	poolTaobaoSimbaKeywordAddAPIResponse.Put(v)
}
