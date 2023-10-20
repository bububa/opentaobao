package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpBidwordSuggestdefaultlistAPIResponse 建议默认关键词 API返回值
// taobao.universalbp.bidword.suggestdefaultlist
//
// 入参推广信息，出参建议的默认关键词
type TaobaoUniversalbpBidwordSuggestdefaultlistAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpBidwordSuggestdefaultlistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpBidwordSuggestdefaultlistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpBidwordSuggestdefaultlistAPIResponseModel).Reset()
}

// TaobaoUniversalbpBidwordSuggestdefaultlistAPIResponseModel is 建议默认关键词 成功返回结果
type TaobaoUniversalbpBidwordSuggestdefaultlistAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_bidword_suggestdefaultlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpBidwordSuggestdefaultlistTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpBidwordSuggestdefaultlistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpBidwordSuggestdefaultlistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpBidwordSuggestdefaultlistAPIResponse)
	},
}

// GetTaobaoUniversalbpBidwordSuggestdefaultlistAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpBidwordSuggestdefaultlistAPIResponse
func GetTaobaoUniversalbpBidwordSuggestdefaultlistAPIResponse() *TaobaoUniversalbpBidwordSuggestdefaultlistAPIResponse {
	return poolTaobaoUniversalbpBidwordSuggestdefaultlistAPIResponse.Get().(*TaobaoUniversalbpBidwordSuggestdefaultlistAPIResponse)
}

// ReleaseTaobaoUniversalbpBidwordSuggestdefaultlistAPIResponse 将 TaobaoUniversalbpBidwordSuggestdefaultlistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpBidwordSuggestdefaultlistAPIResponse(v *TaobaoUniversalbpBidwordSuggestdefaultlistAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpBidwordSuggestdefaultlistAPIResponse.Put(v)
}
