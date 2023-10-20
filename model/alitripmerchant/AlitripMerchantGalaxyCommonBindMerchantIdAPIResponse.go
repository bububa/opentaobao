package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyCommonBindMerchantIdAPIResponse 绑定用户和merchantID API返回值
// alitrip.merchant.galaxy.common.bind.merchant.id
//
// 绑定用户和merchantID
type AlitripMerchantGalaxyCommonBindMerchantIdAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyCommonBindMerchantIdAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyCommonBindMerchantIdAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyCommonBindMerchantIdAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyCommonBindMerchantIdAPIResponseModel is 绑定用户和merchantID 成功返回结果
type AlitripMerchantGalaxyCommonBindMerchantIdAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_common_bind_merchant_id_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 结果
	Result *AlitripMerchantGalaxyCommonBindMerchantIdResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyCommonBindMerchantIdAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyCommonBindMerchantIdAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyCommonBindMerchantIdAPIResponse)
	},
}

// GetAlitripMerchantGalaxyCommonBindMerchantIdAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyCommonBindMerchantIdAPIResponse
func GetAlitripMerchantGalaxyCommonBindMerchantIdAPIResponse() *AlitripMerchantGalaxyCommonBindMerchantIdAPIResponse {
	return poolAlitripMerchantGalaxyCommonBindMerchantIdAPIResponse.Get().(*AlitripMerchantGalaxyCommonBindMerchantIdAPIResponse)
}

// ReleaseAlitripMerchantGalaxyCommonBindMerchantIdAPIResponse 将 AlitripMerchantGalaxyCommonBindMerchantIdAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyCommonBindMerchantIdAPIResponse(v *AlitripMerchantGalaxyCommonBindMerchantIdAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyCommonBindMerchantIdAPIResponse.Put(v)
}
