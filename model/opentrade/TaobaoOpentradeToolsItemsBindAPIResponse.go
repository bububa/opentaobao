package opentrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeToolsItemsBindAPIResponse 交易开放商品绑定 API返回值
// taobao.opentrade.tools.items.bind
//
// 交易开放商品绑定
type TaobaoOpentradeToolsItemsBindAPIResponse struct {
	model.CommonResponse
	TaobaoOpentradeToolsItemsBindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpentradeToolsItemsBindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpentradeToolsItemsBindAPIResponseModel).Reset()
}

// TaobaoOpentradeToolsItemsBindAPIResponseModel is 交易开放商品绑定 成功返回结果
type TaobaoOpentradeToolsItemsBindAPIResponseModel struct {
	XMLName xml.Name `xml:"opentrade_tools_items_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 绑定返回结构
	Results []ItemBindResult `json:"results,omitempty" xml:"results>item_bind_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpentradeToolsItemsBindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolTaobaoOpentradeToolsItemsBindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpentradeToolsItemsBindAPIResponse)
	},
}

// GetTaobaoOpentradeToolsItemsBindAPIResponse 从 sync.Pool 获取 TaobaoOpentradeToolsItemsBindAPIResponse
func GetTaobaoOpentradeToolsItemsBindAPIResponse() *TaobaoOpentradeToolsItemsBindAPIResponse {
	return poolTaobaoOpentradeToolsItemsBindAPIResponse.Get().(*TaobaoOpentradeToolsItemsBindAPIResponse)
}

// ReleaseTaobaoOpentradeToolsItemsBindAPIResponse 将 TaobaoOpentradeToolsItemsBindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpentradeToolsItemsBindAPIResponse(v *TaobaoOpentradeToolsItemsBindAPIResponse) {
	v.Reset()
	poolTaobaoOpentradeToolsItemsBindAPIResponse.Put(v)
}
