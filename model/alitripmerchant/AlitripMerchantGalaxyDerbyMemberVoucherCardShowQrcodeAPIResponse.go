package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponse 会员权益卡身份识别二维码图片 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.card.show.qrcode
//
// 会员权益卡身份识别二维码图片
type AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponseModel is 会员权益卡身份识别二维码图片 成功返回结果
type AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_card_show_qrcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// response
	Result *AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponse() *AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponse 将 AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherCardShowQrcodeAPIResponse.Put(v)
}
