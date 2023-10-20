package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxycardinfoAPIResponse 获取会员体系 API返回值
// alitrip.merchant.galaxy.card.info
//
// 星河=根据卡类型获取当前的会员体系
type AlitripmerchantgalaxycardinfoAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxycardinfoAPIResponseModel
}

// AlitripmerchantgalaxycardinfoAPIResponseModel is 获取会员体系 成功返回结果
type AlitripmerchantgalaxycardinfoAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_card_info_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripmerchantgalaxycardinfoResponse `json:"result,omitempty" xml:"result,omitempty"`
}
