package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaopromotionbenefitactivityupdateAPIResponse 修改关联的活动权益 API返回值
// taobao.promotion.benefit.activity.update
//
// 修改卖家活动中关联的对应的权益。
type TaobaopromotionbenefitactivityupdateAPIResponse struct {
	model.CommonResponse
	TaobaopromotionbenefitactivityupdateAPIResponseModel
}

// TaobaopromotionbenefitactivityupdateAPIResponseModel is 修改关联的活动权益 成功返回结果
type TaobaopromotionbenefitactivityupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"promotion_benefit_activity_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 更新是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
