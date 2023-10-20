package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyBrandSearchAPIResponse 星河-品牌搜索 API返回值
// alitrip.merchant.galaxy.brand.search
//
// 星河服务=获取雅高品牌信息
type AlitripMerchantGalaxyBrandSearchAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyBrandSearchAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyBrandSearchAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyBrandSearchAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyBrandSearchAPIResponseModel is 星河-品牌搜索 成功返回结果
type AlitripMerchantGalaxyBrandSearchAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_brand_search_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 接口返回model
	Result *AlitripMerchantGalaxyBrandSearchResult `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyBrandSearchAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyBrandSearchAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyBrandSearchAPIResponse)
	},
}

// GetAlitripMerchantGalaxyBrandSearchAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyBrandSearchAPIResponse
func GetAlitripMerchantGalaxyBrandSearchAPIResponse() *AlitripMerchantGalaxyBrandSearchAPIResponse {
	return poolAlitripMerchantGalaxyBrandSearchAPIResponse.Get().(*AlitripMerchantGalaxyBrandSearchAPIResponse)
}

// ReleaseAlitripMerchantGalaxyBrandSearchAPIResponse 将 AlitripMerchantGalaxyBrandSearchAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyBrandSearchAPIResponse(v *AlitripMerchantGalaxyBrandSearchAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyBrandSearchAPIResponse.Put(v)
}
