package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyactivityaddressaddAPIResponse 星河-营销抽奖奖品邮寄地址添加 API返回值
// alitrip.merchant.galaxy.activity.address.add
//
// 营销抽奖活动，奖品邮寄地址填写
type AlitripmerchantgalaxyactivityaddressaddAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyactivityaddressaddAPIResponseModel
}

// AlitripmerchantgalaxyactivityaddressaddAPIResponseModel is 星河-营销抽奖奖品邮寄地址添加 成功返回结果
type AlitripmerchantgalaxyactivityaddressaddAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_activity_address_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripmerchantgalaxyactivityaddressaddResponse `json:"result,omitempty" xml:"result,omitempty"`
}
