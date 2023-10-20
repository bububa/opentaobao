package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponse 生成臻享卡德比分销二维码 API返回值
// alitrip.merchant.galaxy.derby.member.generate.seller.qrcode
//
// 生成臻享卡德比分销二维码
type AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponseModel is 生成臻享卡德比分销二维码 成功返回结果
type AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_generate_seller_qrcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponse
func GetAlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponse() *AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponse {
	return poolAlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponse.Get().(*AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponse 将 AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponse(v *AlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberGenerateSellerQrcodeAPIResponse.Put(v)
}
