package travel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoAlitripTravelProductSkuOverrideAPIResponse （供销）产品级别日历价格库存修改，全量覆盖 API返回值
// taobao.alitrip.travel.product.sku.override
//
// （供销）产品级别日历价格库存修改，全量覆盖
type TaobaoAlitripTravelProductSkuOverrideAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelProductSkuOverrideAPIResponseModel
}

// TaobaoAlitripTravelProductSkuOverrideAPIResponseModel is （供销）产品级别日历价格库存修改，全量覆盖 成功返回结果
type TaobaoAlitripTravelProductSkuOverrideAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_product_sku_override_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品sku修改结果
	TravelItem *TopTravelItem `json:"travel_item,omitempty" xml:"travel_item,omitempty"`
}
