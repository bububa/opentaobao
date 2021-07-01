package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoMobilePromotionBenefitActivitySendShareAPIResponse
手淘专用单用户发放接口 API返回值
taobao.mobile.promotion.benefit.activity.send.share

卖家活动中需要通过该API来发放对应的权益。手淘专用、验证分享链路。 */
type TaobaoMobilePromotionBenefitActivitySendShareAPIResponse struct {
	model.CommonResponse
	TaobaoMobilePromotionBenefitActivitySendShareAPIResponseModel
}

// TaobaoMobilePromotionBenefitActivitySendShareAPIResponseModel is 手淘专用单用户发放接口 成功返回结果
type TaobaoMobilePromotionBenefitActivitySendShareAPIResponseModel struct {
	XMLName xml.Name `xml:"mobile_promotion_benefit_activity_send_share_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 权益发放结果
	SendResult *ShareBenefitSendResult `json:"send_result,omitempty" xml:"send_result,omitempty"`
}
