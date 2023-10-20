package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxymembercardAPIResponse 星河-获取会员卡信息 API返回值
// alitrip.merchant.galaxy.member.card
//
// 星河=根据会员等级获取会员的权益
type AlitripmerchantgalaxymembercardAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxymembercardAPIResponseModel
}

// AlitripmerchantgalaxymembercardAPIResponseModel is 星河-获取会员卡信息 成功返回结果
type AlitripmerchantgalaxymembercardAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_member_card_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripmerchantgalaxymembercardResponse `json:"result,omitempty" xml:"result,omitempty"`
}
