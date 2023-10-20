package opentrade

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoOpentradeSpecialItemsBindAPIResponse 专属下单场景商品绑定 API返回值
// taobao.opentrade.special.items.bind
//
// 专属下单场景商品绑定
type TaobaoOpentradeSpecialItemsBindAPIResponse struct {
	model.CommonResponse
	TaobaoOpentradeSpecialItemsBindAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoOpentradeSpecialItemsBindAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoOpentradeSpecialItemsBindAPIResponseModel).Reset()
}

// TaobaoOpentradeSpecialItemsBindAPIResponseModel is 专属下单场景商品绑定 成功返回结果
type TaobaoOpentradeSpecialItemsBindAPIResponseModel struct {
	XMLName xml.Name `xml:"opentrade_special_items_bind_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 绑定返回结构
	Results []ItemBindResult `json:"results,omitempty" xml:"results>item_bind_result,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoOpentradeSpecialItemsBindAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Results = m.Results[:0]
}

var poolTaobaoOpentradeSpecialItemsBindAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoOpentradeSpecialItemsBindAPIResponse)
	},
}

// GetTaobaoOpentradeSpecialItemsBindAPIResponse 从 sync.Pool 获取 TaobaoOpentradeSpecialItemsBindAPIResponse
func GetTaobaoOpentradeSpecialItemsBindAPIResponse() *TaobaoOpentradeSpecialItemsBindAPIResponse {
	return poolTaobaoOpentradeSpecialItemsBindAPIResponse.Get().(*TaobaoOpentradeSpecialItemsBindAPIResponse)
}

// ReleaseTaobaoOpentradeSpecialItemsBindAPIResponse 将 TaobaoOpentradeSpecialItemsBindAPIResponse 保存到 sync.Pool
func ReleaseTaobaoOpentradeSpecialItemsBindAPIResponse(v *TaobaoOpentradeSpecialItemsBindAPIResponse) {
	v.Reset()
	poolTaobaoOpentradeSpecialItemsBindAPIResponse.Put(v)
}
