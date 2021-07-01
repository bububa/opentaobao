package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlitripMerchantGalaxyProviderMemberQueryAPIResponse
提供会员查询接口 API返回值
alitrip.merchant.galaxy.provider.member.query

星河产品=提供会查询服务 */
type AlitripMerchantGalaxyProviderMemberQueryAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyProviderMemberQueryAPIResponseModel
}

// AlitripMerchantGalaxyProviderMemberQueryAPIResponseModel is 提供会员查询接口 成功返回结果
type AlitripMerchantGalaxyProviderMemberQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_provider_member_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyProviderMemberQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}
