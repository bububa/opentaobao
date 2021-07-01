package travel

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoAlitripTravelProductBaseModifyAPIResponse
供应商编辑产品API API返回值
taobao.alitrip.travel.product.base.modify

飞猪供销平台供应商可通过该API编辑产品 */
type TaobaoAlitripTravelProductBaseModifyAPIResponse struct {
	model.CommonResponse
	TaobaoAlitripTravelProductBaseModifyAPIResponseModel
}

// TaobaoAlitripTravelProductBaseModifyAPIResponseModel is 供应商编辑产品API 成功返回结果
type TaobaoAlitripTravelProductBaseModifyAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_travel_product_base_modify_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 商品修改结果
	TravelItem *TopTravelItem `json:"travel_item,omitempty" xml:"travel_item,omitempty"`
}
