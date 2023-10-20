package alitripmerchant

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlitripMerchantGalaxyCityListAPIResponse 星河-酒店城市列表展示 API返回值
// alitrip.merchant.galaxy.city.list
//
// 雅高酒店城市列表展示，并且首字母列出酒店城市
type AlitripMerchantGalaxyCityListAPIResponse struct {
	model.CommonResponse
	AlitripMerchantGalaxyCityListAPIResponseModel
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyCityListAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlitripMerchantGalaxyCityListAPIResponseModel).Reset()
}

// AlitripMerchantGalaxyCityListAPIResponseModel is 星河-酒店城市列表展示 成功返回结果
type AlitripMerchantGalaxyCityListAPIResponseModel struct {
	XMLName xml.Name `xml:"alitrip_merchant_galaxy_city_list_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 默认描述
	Result *AlitripMerchantGalaxyCityListResponse `json:"result,omitempty" xml:"result,omitempty"`
}

// Reset 清空结构体
func (m *AlitripMerchantGalaxyCityListAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Result = nil
}

var poolAlitripMerchantGalaxyCityListAPIResponse = sync.Pool{
	New: func() any {
		return new(AlitripMerchantGalaxyCityListAPIResponse)
	},
}

// GetAlitripMerchantGalaxyCityListAPIResponse 从 sync.Pool 获取 AlitripMerchantGalaxyCityListAPIResponse
func GetAlitripMerchantGalaxyCityListAPIResponse() *AlitripMerchantGalaxyCityListAPIResponse {
	return poolAlitripMerchantGalaxyCityListAPIResponse.Get().(*AlitripMerchantGalaxyCityListAPIResponse)
}

// ReleaseAlitripMerchantGalaxyCityListAPIResponse 将 AlitripMerchantGalaxyCityListAPIResponse 保存到 sync.Pool
func ReleaseAlitripMerchantGalaxyCityListAPIResponse(v *AlitripMerchantGalaxyCityListAPIResponse) {
	v.Reset()
	poolAlitripMerchantGalaxyCityListAPIResponse.Put(v)
}
