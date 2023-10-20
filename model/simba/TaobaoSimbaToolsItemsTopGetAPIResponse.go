package simba

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoSimbaToolsItemsTopGetAPIResponse 取得一个关键词的推广组排名列表 API返回值
// taobao.simba.tools.items.top.get
//
// 取得一个关键词的推广组排名列表
type TaobaoSimbaToolsItemsTopGetAPIResponse struct {
	model.CommonResponse
	TaobaoSimbaToolsItemsTopGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoSimbaToolsItemsTopGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoSimbaToolsItemsTopGetAPIResponseModel).Reset()
}

// TaobaoSimbaToolsItemsTopGetAPIResponseModel is 取得一个关键词的推广组排名列表 成功返回结果
type TaobaoSimbaToolsItemsTopGetAPIResponseModel struct {
	XMLName xml.Name `xml:"simba_tools_items_top_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推广组信息列表
	Rankeditems []RankedItem `json:"rankeditems,omitempty" xml:"rankeditems>ranked_item,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoSimbaToolsItemsTopGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Rankeditems = m.Rankeditems[:0]
}

var poolTaobaoSimbaToolsItemsTopGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoSimbaToolsItemsTopGetAPIResponse)
	},
}

// GetTaobaoSimbaToolsItemsTopGetAPIResponse 从 sync.Pool 获取 TaobaoSimbaToolsItemsTopGetAPIResponse
func GetTaobaoSimbaToolsItemsTopGetAPIResponse() *TaobaoSimbaToolsItemsTopGetAPIResponse {
	return poolTaobaoSimbaToolsItemsTopGetAPIResponse.Get().(*TaobaoSimbaToolsItemsTopGetAPIResponse)
}

// ReleaseTaobaoSimbaToolsItemsTopGetAPIResponse 将 TaobaoSimbaToolsItemsTopGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoSimbaToolsItemsTopGetAPIResponse(v *TaobaoSimbaToolsItemsTopGetAPIResponse) {
	v.Reset()
	poolTaobaoSimbaToolsItemsTopGetAPIResponse.Put(v)
}
