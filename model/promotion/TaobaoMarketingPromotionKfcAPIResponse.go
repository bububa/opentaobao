package promotion

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// TaobaomarketingpromotionkfcAPIResponse 定向优惠活动名称与描述违禁词检查 API返回值
// taobao.marketing.promotion.kfc
//
// 活动名称与描述违禁词检查
type TaobaomarketingpromotionkfcAPIResponse struct {
	model.CommonResponse
	TaobaomarketingpromotionkfcAPIResponseModel
}

// TaobaomarketingpromotionkfcAPIResponseModel is 定向优惠活动名称与描述违禁词检查 成功返回结果
type TaobaomarketingpromotionkfcAPIResponseModel struct {
	XMLName xml.Name `xml:"marketing_promotion_kfc_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
