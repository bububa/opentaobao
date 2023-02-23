package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIResponse 德比付费会员卡可购查询 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.card.purchasable.query
//
// 德比付费会员卡可购查询
type AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIResponseModel
}

// AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIResponseModel is 德比付费会员卡可购查询 成功返回结果
type AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_card_purchasable_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 1
	Result *AlitripMerchantGalaxyDerbyMemberVoucherCardPurchasableQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}
