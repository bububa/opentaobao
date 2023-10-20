package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaKeywordUpdateAPIResponse （新）关键词更新相关接口 API返回值
// taobao.simba.keyword.update
//
// （新）关键词更新相关接口
type TaobaoSimbaKeywordUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaKeywordUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaKeywordUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaKeywordUpdateAPIResponseModel).Reset()
}

// TaobaoSimbaKeywordUpdateAPIResponseModel is （新）关键词更新相关接口 成功返回结果
type TaobaoSimbaKeywordUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_keyword_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 整体的返回值
	Results []SiriusBidwordDto `json:"results,omitempty" xml:"results>sirius_bidword_dto,omitempty"`
	// 错误原因
	ErrorMsg string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaKeywordUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
	m.ErrorMsg = ""
}

var poolTaobaoSimbaKeywordUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaKeywordUpdateAPIResponse)
	},
}

// GetTaobaoSimbaKeywordUpdateAPIResponse 从 sync.Pool 获取 TaobaoSimbaKeywordUpdateAPIResponse
func GetTaobaoSimbaKeywordUpdateAPIResponse() *TaobaoSimbaKeywordUpdateAPIResponse {
	return poolTaobaoSimbaKeywordUpdateAPIResponse.Get().(*TaobaoSimbaKeywordUpdateAPIResponse)
}

// ReleaseTaobaoSimbaKeywordUpdateAPIResponse 将 TaobaoSimbaKeywordUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaKeywordUpdateAPIResponse(v *TaobaoSimbaKeywordUpdateAPIResponse) {
	v.Reset()
	poolTaobaoSimbaKeywordUpdateAPIResponse.Put(v)
}
