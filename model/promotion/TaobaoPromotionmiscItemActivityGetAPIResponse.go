package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscItemActivityGetAPIResponse 查询无条件单品优惠活动 API返回值
// taobao.promotionmisc.item.activity.get
//
// 查询无条件单品优惠活动
type TaobaoPromotionmiscItemActivityGetAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionmiscItemActivityGetAPIResponseModel
}

// TaobaoPromotionmiscItemActivityGetAPIResponseModel is 查询无条件单品优惠活动 成功返回结果
type TaobaoPromotionmiscItemActivityGetAPIResponseModel struct {
	XMLName xml.Name `xml:"promotionmisc_item_activity_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 单品优惠活动信息。
	ItemPromotion *ItemPromotion `json:"item_promotion,omitempty" xml:"item_promotion,omitempty"`
}
