package travel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelProductBaseModifyAPIResponse 供应商编辑产品API API返回值
// taobao.alitrip.travel.product.base.modify
//
// 飞猪供销平台供应商可通过该API编辑产品
type TaobaoAlitripTravelProductBaseModifyAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelProductBaseModifyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelProductBaseModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelProductBaseModifyAPIResponseModel).Reset()
}

// TaobaoAlitripTravelProductBaseModifyAPIResponseModel is 供应商编辑产品API 成功返回结果
type TaobaoAlitripTravelProductBaseModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_product_base_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品修改结果
	TravelItem *TopTravelItem `json:"travel_item,omitempty" xml:"travel_item,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelProductBaseModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TravelItem = nil
}

var poolTaobaoAlitripTravelProductBaseModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelProductBaseModifyAPIResponse)
	},
}

// GetTaobaoAlitripTravelProductBaseModifyAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelProductBaseModifyAPIResponse
func GetTaobaoAlitripTravelProductBaseModifyAPIResponse() *TaobaoAlitripTravelProductBaseModifyAPIResponse {
	return poolTaobaoAlitripTravelProductBaseModifyAPIResponse.Get().(*TaobaoAlitripTravelProductBaseModifyAPIResponse)
}

// ReleaseTaobaoAlitripTravelProductBaseModifyAPIResponse 将 TaobaoAlitripTravelProductBaseModifyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelProductBaseModifyAPIResponse(v *TaobaoAlitripTravelProductBaseModifyAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelProductBaseModifyAPIResponse.Put(v)
}
