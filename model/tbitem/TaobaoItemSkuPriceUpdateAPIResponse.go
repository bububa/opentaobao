package tbitem

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoItemSkuPriceUpdateAPIResponse 更新商品SKU的价格 API返回值
// taobao.item.sku.price.update
//
// 更新商品SKU的价格
type TaobaoItemSkuPriceUpdateAPIResponse struct {
	model.CommonResponse
	TaobaoItemSkuPriceUpdateAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoItemSkuPriceUpdateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoItemSkuPriceUpdateAPIResponseModel).Reset()
}

// TaobaoItemSkuPriceUpdateAPIResponseModel is 更新商品SKU的价格 成功返回结果
type TaobaoItemSkuPriceUpdateAPIResponseModel struct {
	XMLName xml.Name `xml:"item_sku_price_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品SKU信息（只包含num_iid和modified）
	Sku *Sku `json:"sku,omitempty" xml:"sku,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoItemSkuPriceUpdateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Sku = nil
}

var poolTaobaoItemSkuPriceUpdateAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoItemSkuPriceUpdateAPIResponse)
	},
}

// GetTaobaoItemSkuPriceUpdateAPIResponse 从 sync.Pool 获取 TaobaoItemSkuPriceUpdateAPIResponse
func GetTaobaoItemSkuPriceUpdateAPIResponse() *TaobaoItemSkuPriceUpdateAPIResponse {
	return poolTaobaoItemSkuPriceUpdateAPIResponse.Get().(*TaobaoItemSkuPriceUpdateAPIResponse)
}

// ReleaseTaobaoItemSkuPriceUpdateAPIResponse 将 TaobaoItemSkuPriceUpdateAPIResponse 保存到 sync.Pool
func ReleaseTaobaoItemSkuPriceUpdateAPIResponse(v *TaobaoItemSkuPriceUpdateAPIResponse) {
	v.Reset()
	poolTaobaoItemSkuPriceUpdateAPIResponse.Put(v)
}
