package travel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelItemSkuPriceModifyAPIResponse 【API3.0】日期级别日历价格库存修改，增量维护 API返回值
// taobao.alitrip.travel.item.sku.price.modify
//
// 【API3.0】日期级别日历价格库存增量维护
type TaobaoAlitripTravelItemSkuPriceModifyAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelItemSkuPriceModifyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelItemSkuPriceModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelItemSkuPriceModifyAPIResponseModel).Reset()
}

// TaobaoAlitripTravelItemSkuPriceModifyAPIResponseModel is 【API3.0】日期级别日历价格库存修改，增量维护 成功返回结果
type TaobaoAlitripTravelItemSkuPriceModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_item_sku_price_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 日期级别日历价格库存增量维护
	TravelItem *TopTravelItem `json:"travel_item,omitempty" xml:"travel_item,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelItemSkuPriceModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TravelItem = nil
}

var poolTaobaoAlitripTravelItemSkuPriceModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelItemSkuPriceModifyAPIResponse)
	},
}

// GetTaobaoAlitripTravelItemSkuPriceModifyAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelItemSkuPriceModifyAPIResponse
func GetTaobaoAlitripTravelItemSkuPriceModifyAPIResponse() *TaobaoAlitripTravelItemSkuPriceModifyAPIResponse {
	return poolTaobaoAlitripTravelItemSkuPriceModifyAPIResponse.Get().(*TaobaoAlitripTravelItemSkuPriceModifyAPIResponse)
}

// ReleaseTaobaoAlitripTravelItemSkuPriceModifyAPIResponse 将 TaobaoAlitripTravelItemSkuPriceModifyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelItemSkuPriceModifyAPIResponse(v *TaobaoAlitripTravelItemSkuPriceModifyAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelItemSkuPriceModifyAPIResponse.Put(v)
}
