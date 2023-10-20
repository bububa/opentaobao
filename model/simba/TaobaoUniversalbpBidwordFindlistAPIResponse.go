package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoUniversalbpBidwordFindlistAPIResponse 词列表查询 API返回值
// taobao.universalbp.bidword.findlist
//
// 根据计划+单元id，查绑定的词列表
type TaobaoUniversalbpBidwordFindlistAPIResponse struct {
	model.CommonResponse
	TaobaoUniversalbpBidwordFindlistAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoUniversalbpBidwordFindlistAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoUniversalbpBidwordFindlistAPIResponseModel).Reset()
}

// TaobaoUniversalbpBidwordFindlistAPIResponseModel is 词列表查询 成功返回结果
type TaobaoUniversalbpBidwordFindlistAPIResponseModel struct {
	XMLName xml.Name `xml:"universalbp_bidword_findlist_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果体
	Result *TaobaoUniversalbpBidwordFindlistTopResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoUniversalbpBidwordFindlistAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolTaobaoUniversalbpBidwordFindlistAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoUniversalbpBidwordFindlistAPIResponse)
	},
}

// GetTaobaoUniversalbpBidwordFindlistAPIResponse 从 sync.Pool 获取 TaobaoUniversalbpBidwordFindlistAPIResponse
func GetTaobaoUniversalbpBidwordFindlistAPIResponse() *TaobaoUniversalbpBidwordFindlistAPIResponse {
	return poolTaobaoUniversalbpBidwordFindlistAPIResponse.Get().(*TaobaoUniversalbpBidwordFindlistAPIResponse)
}

// ReleaseTaobaoUniversalbpBidwordFindlistAPIResponse 将 TaobaoUniversalbpBidwordFindlistAPIResponse 保存到 sync.Pool
func ReleaseTaobaoUniversalbpBidwordFindlistAPIResponse(v *TaobaoUniversalbpBidwordFindlistAPIResponse) {
	v.Reset()
	poolTaobaoUniversalbpBidwordFindlistAPIResponse.Put(v)
}
