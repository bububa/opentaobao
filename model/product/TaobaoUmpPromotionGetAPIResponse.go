package product

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoumppromotiongetAPIResponse 商品优惠详情查询 API返回值
// taobao.ump.promotion.get
//
// 商品优惠详情查询，可查询商品设置的详细优惠。包括限时折扣，满就送等官方优惠以及第三方优惠。
type TaobaoumppromotiongetAPIResponse struct {
	model.CommonResponse
	TaobaoumppromotiongetAPIResponseModel
}

// TaobaoumppromotiongetAPIResponseModel is 商品优惠详情查询 成功返回结果
type TaobaoumppromotiongetAPIResponseModel struct {
	XMLName xml.Name `xml:"ump_promotion_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 优惠详细信息
	Promotions *PromotionDisplayTop `json:"promotions,omitempty" xml:"promotions,omitempty"`
}
