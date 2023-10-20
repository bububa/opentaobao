package travel

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelItemSkuPackageModifyAPIResponse 【API3.0】套餐级别日历价格库存增删操作 API返回值
// taobao.alitrip.travel.item.sku.package.modify
//
// 【API3.0】套餐级别日历价格库存增删操作
type TaobaoAlitripTravelItemSkuPackageModifyAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelItemSkuPackageModifyAPIResponseModel
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelItemSkuPackageModifyAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.TaobaoAlitripTravelItemSkuPackageModifyAPIResponseModel).Reset()
}

// TaobaoAlitripTravelItemSkuPackageModifyAPIResponseModel is 【API3.0】套餐级别日历价格库存增删操作 成功返回结果
type TaobaoAlitripTravelItemSkuPackageModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_item_sku_package_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品sku修改结果
	TravelItem *TopTravelItem `json:"travel_item,omitempty" xml:"travel_item,omitempty"`
}

// Reset 清空结构体
func (m *TaobaoAlitripTravelItemSkuPackageModifyAPIResponseModel) Reset() {
	m.RequestId = ""
	m.TravelItem = nil
}

var poolTaobaoAlitripTravelItemSkuPackageModifyAPIResponse = sync.Pool{
	New: func() any {
		return new(TaobaoAlitripTravelItemSkuPackageModifyAPIResponse)
	},
}

// GetTaobaoAlitripTravelItemSkuPackageModifyAPIResponse 从 sync.Pool 获取 TaobaoAlitripTravelItemSkuPackageModifyAPIResponse
func GetTaobaoAlitripTravelItemSkuPackageModifyAPIResponse() *TaobaoAlitripTravelItemSkuPackageModifyAPIResponse {
	return poolTaobaoAlitripTravelItemSkuPackageModifyAPIResponse.Get().(*TaobaoAlitripTravelItemSkuPackageModifyAPIResponse)
}

// ReleaseTaobaoAlitripTravelItemSkuPackageModifyAPIResponse 将 TaobaoAlitripTravelItemSkuPackageModifyAPIResponse 保存到 sync.Pool
func ReleaseTaobaoAlitripTravelItemSkuPackageModifyAPIResponse(v *TaobaoAlitripTravelItemSkuPackageModifyAPIResponse) {
	v.Reset()
	poolTaobaoAlitripTravelItemSkuPackageModifyAPIResponse.Put(v)
}
