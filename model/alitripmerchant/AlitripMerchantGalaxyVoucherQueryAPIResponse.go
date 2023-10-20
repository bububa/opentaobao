package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyVoucherQueryAPIResponse 查询单个代金券信息 API返回值
// alitrip.merchant.galaxy.voucher.query
//
// 查询单个代金券信息
type AlitripMerchantGalaxyVoucherQueryAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyVoucherQueryAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyVoucherQueryAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyVoucherQueryAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyVoucherQueryAPIResponseModel is 查询单个代金券信息 成功返回结果
type AlitripMerchantGalaxyVoucherQueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_voucher_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 返回参数
	Result *AlitripMerchantGalaxyVoucherQueryResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyVoucherQueryAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyVoucherQueryAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyVoucherQueryAPIResponse)
	},
}

// GetAlitripMerchantGalaxyVoucherQueryAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyVoucherQueryAPIResponse
func GetAlitripMerchantGalaxyVoucherQueryAPIResponse() *AlitripMerchantGalaxyVoucherQueryAPIResponse {
	return poolAlitripMerchantGalaxyVoucherQueryAPIResponse.Get().(*AlitripMerchantGalaxyVoucherQueryAPIResponse)
}

// ReleaseAlitripMerchantGalaxyVoucherQueryAPIResponse 将 AlitripMerchantGalaxyVoucherQueryAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyVoucherQueryAPIResponse(v *AlitripMerchantGalaxyVoucherQueryAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyVoucherQueryAPIResponse.Put(v)
}
