package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaoPromotionmiscMjsActivityGetAPIResponse 查询满就送活动 API返回值
// taobao.promotionmisc.mjs.activity.get
//
// 查询满就送活动
type TaobaoPromotionmiscMjsActivityGetAPIResponse struct {
	model.CommonResponse
	TaobaoPromotionmiscMjsActivityGetAPIResponseModel
}

// TaobaoPromotionmiscMjsActivityGetAPIResponseModel is 查询满就送活动 成功返回结果
type TaobaoPromotionmiscMjsActivityGetAPIResponseModel struct {
	XMLName xml.Name `xml:"promotionmisc_mjs_activity_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 满就送活动信息。
	MjsPromotion *MjsPromotion `json:"mjs_promotion,omitempty" xml:"mjs_promotion,omitempty"`
}
