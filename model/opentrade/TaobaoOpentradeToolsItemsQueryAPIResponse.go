package opentrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeToolsItemsQueryAPIResponse 交易开放获取商品绑定信息 API返回值
// taobao.opentrade.tools.items.query
//
// 交易开放获取商品绑定信息
type TaobaoOpentradeToolsItemsQueryAPIResponse struct {
	model.CommonResponse
	TaobaoOpentradeToolsItemsQueryAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpentradeToolsItemsQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpentradeToolsItemsQueryAPIResponseModel).Reset()
}

// TaobaoOpentradeToolsItemsQueryAPIResponseModel is 交易开放获取商品绑定信息 成功返回结果
type TaobaoOpentradeToolsItemsQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"opentrade_tools_items_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 已绑定的商品ID列表
	ItemIds []int64 `json:"item_ids,omitempty" xml:"item_ids>int64,omitempty"`
	// 总商品数
	TotalCount int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpentradeToolsItemsQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.ItemIds = m.ItemIds[:0]
	m.TotalCount = 0
}

var poolTaobaoOpentradeToolsItemsQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpentradeToolsItemsQueryAPIResponse)
	},
}

// GetTaobaoOpentradeToolsItemsQueryAPIResponse 从 sync.Pool 获取 TaobaoOpentradeToolsItemsQueryAPIResponse
func GetTaobaoOpentradeToolsItemsQueryAPIResponse() *TaobaoOpentradeToolsItemsQueryAPIResponse {
	return poolTaobaoOpentradeToolsItemsQueryAPIResponse.Get().(*TaobaoOpentradeToolsItemsQueryAPIResponse)
}

// ReleaseTaobaoOpentradeToolsItemsQueryAPIResponse 将 TaobaoOpentradeToolsItemsQueryAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpentradeToolsItemsQueryAPIResponse(v *TaobaoOpentradeToolsItemsQueryAPIResponse) {
	v.Reset()
	poolTaobaoOpentradeToolsItemsQueryAPIResponse.Put(v)
}
