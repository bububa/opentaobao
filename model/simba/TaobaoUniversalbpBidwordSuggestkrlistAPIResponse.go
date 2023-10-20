package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpBidwordSuggestkrlistAPIResponse 关键词建议 API返回值
// taobao.universalbp.bidword.suggestkrlist
//
// 入参推广信息，出参建议的全部关键词
type TaobaoUniversalbpBidwordSuggestkrlistAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpBidwordSuggestkrlistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpBidwordSuggestkrlistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpBidwordSuggestkrlistAPIResponseModel).Reset()
}

// TaobaoUniversalbpBidwordSuggestkrlistAPIResponseModel is 关键词建议 成功返回结果
type TaobaoUniversalbpBidwordSuggestkrlistAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_bidword_suggestkrlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpBidwordSuggestkrlistTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpBidwordSuggestkrlistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpBidwordSuggestkrlistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpBidwordSuggestkrlistAPIResponse)
	},
}

// GetTaobaoUniversalbpBidwordSuggestkrlistAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpBidwordSuggestkrlistAPIResponse
func GetTaobaoUniversalbpBidwordSuggestkrlistAPIResponse() *TaobaoUniversalbpBidwordSuggestkrlistAPIResponse {
	return poolTaobaoUniversalbpBidwordSuggestkrlistAPIResponse.Get().(*TaobaoUniversalbpBidwordSuggestkrlistAPIResponse)
}

// ReleaseTaobaoUniversalbpBidwordSuggestkrlistAPIResponse 将 TaobaoUniversalbpBidwordSuggestkrlistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpBidwordSuggestkrlistAPIResponse(v *TaobaoUniversalbpBidwordSuggestkrlistAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpBidwordSuggestkrlistAPIResponse.Put(v)
}
