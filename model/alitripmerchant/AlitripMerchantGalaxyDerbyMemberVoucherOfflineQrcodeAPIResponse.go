package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIResponse 德比线下权益券二维码查询 API返回值
// alitrip.merchant.galaxy.derby.member.voucher.offline.qrcode
//
// 德比线下权益券二维码查询
type AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIResponseModel is 德比线下权益券二维码查询 成功返回结果
type AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_derby_member_voucher_offline_qrcode_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 线下核销券二维码返回结果
	Result *AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIResponse)
	},
}

// GetAlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIResponse
func GetAlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIResponse() *AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIResponse {
	return poolAlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIResponse.Get().(*AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIResponse)
}

// ReleaseAlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIResponse 将 AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIResponse(v *AlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyDerbyMemberVoucherOfflineQrcodeAPIResponse.Put(v)
}
