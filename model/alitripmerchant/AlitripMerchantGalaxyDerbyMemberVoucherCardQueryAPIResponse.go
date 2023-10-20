package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponse 德比付费会员卡查询 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.card.query
//
// 德比付费会员卡查询
type AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponseModel
}

// AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponseModel is 德比付费会员卡查询 成功返回结果
type AlitripMerchantGalaxyDerbyMemberVoucherCardQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_card_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	Result *AlitripMerchantGalaxyDerbyMemberVoucherCardQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}
