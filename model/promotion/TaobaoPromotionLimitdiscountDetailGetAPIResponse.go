package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionLimitdiscountDetailGetAPIResponse 限时打折详情查询 API返回值
// taobao.promotion.limitdiscount.detail.get
//
// 限时打折详情查询。查询出指定限时打折的对应商品记录信息。
type TaobaoPromotionLimitdiscountDetailGetAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionLimitdiscountDetailGetAPIResponseModel
}

// TaobaoPromotionLimitdiscountDetailGetAPIResponseModel is 限时打折详情查询 成功返回结果
type TaobaoPromotionLimitdiscountDetailGetAPIResponseModel struct {
	XMLName xml.Name `xml:"promotion_limitdiscount_detail_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 限时打折对应的商品详情列表。
	ItemDiscountDetailList []LimitDiscountDetail `json:"item_discount_detail_list,omitempty" xml:"item_discount_detail_list>limit_discount_detail,omitempty"`
}
