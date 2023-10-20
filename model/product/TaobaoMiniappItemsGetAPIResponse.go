package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoMiniappItemsGetAPIResponse 批量获取商品信息 API返回值
// taobao.miniapp.items.get
//
// 获取商品公开属性，只允许在商家应用环境中使用
type TaobaoMiniappItemsGetAPIResponse struct {
	model.CommonResponse
	TaobaoMiniappItemsGetAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoMiniappItemsGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoMiniappItemsGetAPIResponseModel).Reset()
}

// TaobaoMiniappItemsGetAPIResponseModel is 批量获取商品信息 成功返回结果
type TaobaoMiniappItemsGetAPIResponseModel struct {
	XMLName xml.Name `xml:"miniapp_items_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// Item(商品)结构
	Items []Item `json:"items,omitempty" xml:"items>item,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoMiniappItemsGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Items = m.Items[:0]
}

var poolTaobaoMiniappItemsGetAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoMiniappItemsGetAPIResponse)
	},
}

// GetTaobaoMiniappItemsGetAPIResponse 从 sync.Pool 获取 TaobaoMiniappItemsGetAPIResponse
func GetTaobaoMiniappItemsGetAPIResponse() *TaobaoMiniappItemsGetAPIResponse {
	return poolTaobaoMiniappItemsGetAPIResponse.Get().(*TaobaoMiniappItemsGetAPIResponse)
}

// ReleaseTaobaoMiniappItemsGetAPIResponse 将 TaobaoMiniappItemsGetAPIResponse 保存到 sync.Pool
func ReleaseTaobaoMiniappItemsGetAPIResponse(v *TaobaoMiniappItemsGetAPIResponse) {
	v.Reset()
	poolTaobaoMiniappItemsGetAPIResponse.Put(v)
}
