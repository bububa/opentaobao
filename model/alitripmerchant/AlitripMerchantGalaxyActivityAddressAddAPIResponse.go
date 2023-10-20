package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyActivityAddressAddAPIResponse 星河-营销抽奖奖品邮寄地址添加 API返回值
// alitrip.merchant.galaxy.activity.address.add
//
// 营销抽奖活动，奖品邮寄地址填写
type AlitripMerchantGalaxyActivityAddressAddAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyActivityAddressAddAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyActivityAddressAddAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyActivityAddressAddAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyActivityAddressAddAPIResponseModel is 星河-营销抽奖奖品邮寄地址添加 成功返回结果
type AlitripMerchantGalaxyActivityAddressAddAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_activity_address_add_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyActivityAddressAddResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyActivityAddressAddAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyActivityAddressAddAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyActivityAddressAddAPIResponse)
	},
}

// GetAlitripMerchantGalaxyActivityAddressAddAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyActivityAddressAddAPIResponse
func GetAlitripMerchantGalaxyActivityAddressAddAPIResponse() *AlitripMerchantGalaxyActivityAddressAddAPIResponse {
	return poolAlitripMerchantGalaxyActivityAddressAddAPIResponse.Get().(*AlitripMerchantGalaxyActivityAddressAddAPIResponse)
}

// ReleaseAlitripMerchantGalaxyActivityAddressAddAPIResponse 将 AlitripMerchantGalaxyActivityAddressAddAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyActivityAddressAddAPIResponse(v *AlitripMerchantGalaxyActivityAddressAddAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyActivityAddressAddAPIResponse.Put(v)
}
