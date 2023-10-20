package opentrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeToolsItemsUnbindAPIResponse 交易开放商品解绑 API返回值
// taobao.opentrade.tools.items.unbind
//
// 交易开放商品解绑
type TaobaoOpentradeToolsItemsUnbindAPIResponse struct {
	model.CommonResponse
	TaobaoOpentradeToolsItemsUnbindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpentradeToolsItemsUnbindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpentradeToolsItemsUnbindAPIResponseModel).Reset()
}

// TaobaoOpentradeToolsItemsUnbindAPIResponseModel is 交易开放商品解绑 成功返回结果
type TaobaoOpentradeToolsItemsUnbindAPIResponseModel struct {
	XMLName xml.Name `xml:"opentrade_tools_items_unbind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 解绑返回结构
	Results []ItemUnBindResult `json:"results,omitempty" xml:"results>item_un_bind_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpentradeToolsItemsUnbindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolTaobaoOpentradeToolsItemsUnbindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpentradeToolsItemsUnbindAPIResponse)
	},
}

// GetTaobaoOpentradeToolsItemsUnbindAPIResponse 从 sync.Pool 获取 TaobaoOpentradeToolsItemsUnbindAPIResponse
func GetTaobaoOpentradeToolsItemsUnbindAPIResponse() *TaobaoOpentradeToolsItemsUnbindAPIResponse {
	return poolTaobaoOpentradeToolsItemsUnbindAPIResponse.Get().(*TaobaoOpentradeToolsItemsUnbindAPIResponse)
}

// ReleaseTaobaoOpentradeToolsItemsUnbindAPIResponse 将 TaobaoOpentradeToolsItemsUnbindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpentradeToolsItemsUnbindAPIResponse(v *TaobaoOpentradeToolsItemsUnbindAPIResponse) {
	v.Reset()
	poolTaobaoOpentradeToolsItemsUnbindAPIResponse.Put(v)
}
