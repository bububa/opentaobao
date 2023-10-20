package alitripmerchant

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlitripmerchantgalaxyderbymembervoucherofflineqrcodeAPIResponse 德比线下权益券二维码查询 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.offline.qrcode
//
// 德比线下权益券二维码查询
type AlitripmerchantgalaxyderbymembervoucherofflineqrcodeAPIResponse struct {
	model.CommonResponse
	AlitripmerchantgalaxyderbymembervoucherofflineqrcodeAPIResponseModel
}

// AlitripmerchantgalaxyderbymembervoucherofflineqrcodeAPIResponseModel is 德比线下权益券二维码查询 成功返回结果
type AlitripmerchantgalaxyderbymembervoucherofflineqrcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_offline_qrcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 线下核销券二维码返回结果
	Result *AlitripmerchantgalaxyderbymembervoucherofflineqrcodeResponse `json:"result,omitempty" xml:"result,omitempty"`
}
