package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyMemberQueryAPIResponse
星河-获取登录用户的信息 API返回值
alitrip.merchant.galaxy.member.query

获取登录用户的信息 */
type AlitripMerchantGalaxyMemberQueryAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyMemberQueryAPIResponseModel
}

// AlitripMerchantGalaxyMemberQueryAPIResponseModel is 星河-获取登录用户的信息 成功返回结果
type AlitripMerchantGalaxyMemberQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_member_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyMemberQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}
