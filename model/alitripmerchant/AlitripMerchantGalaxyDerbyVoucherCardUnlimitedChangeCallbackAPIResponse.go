package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIResponse 德比无限次券核销通知接口 API返回值
// alitrip.merchant.galaxy.derby.voucher.card.unlimited.change.callback
//
// 德比无限次券核销通知接口
type AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIResponseModel
}

// AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIResponseModel is 德比无限次券核销通知接口 成功返回结果
type AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_voucher_card_unlimited_change_callback_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回值
	Result *AlitripMerchantGalaxyDerbyVoucherCardUnlimitedChangeCallbackResponse `json:"result,omitempty" xml:"result,omitempty"`
}
