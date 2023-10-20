package aesolution

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSolutionMerchantProfileGetAPIResponse aliexpress.solution.merchant.profile.get API返回值
// aliexpress.solution.merchant.profile.get
//
// API for oversea sellers to obtain the normal information, e.g. store id, registration country code.
type AliexpressSolutionMerchantProfileGetAPIResponse struct {
	model.CommonResponse
	AliexpressSolutionMerchantProfileGetAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSolutionMerchantProfileGetAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSolutionMerchantProfileGetAPIResponseModel).Reset()
}

// AliexpressSolutionMerchantProfileGetAPIResponseModel is aliexpress.solution.merchant.profile.get 成功返回结果
type AliexpressSolutionMerchantProfileGetAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_solution_merchant_profile_get_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// country code which the merchant chose when registered
	CountryCode string `json:"country_code,omitempty" xml:"country_code,omitempty"`
	// merchant login id of Aliexpress
	MerchantLoginId string `json:"merchant_login_id,omitempty" xml:"merchant_login_id,omitempty"`
	// shop name
	ShopName string `json:"shop_name,omitempty" xml:"shop_name,omitempty"`
	// shop type
	ShopType string `json:"shop_type,omitempty" xml:"shop_type,omitempty"`
	// shop url
	ShopUrl string `json:"shop_url,omitempty" xml:"shop_url,omitempty"`
	// shop id
	ShopId int64 `json:"shop_id,omitempty" xml:"shop_id,omitempty"`
	// Indicate whether the mechant could post product or not. FALSE means the merchant could normally post product.
	ProductPostingForbidden bool `json:"product_posting_forbidden,omitempty" xml:"product_posting_forbidden,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSolutionMerchantProfileGetAPIResponseModel) Reset() {
	m.RequestId = ""
	m.CountryCode = ""
	m.MerchantLoginId = ""
	m.ShopName = ""
	m.ShopType = ""
	m.ShopUrl = ""
	m.ShopId = 0
	m.ProductPostingForbidden = false
}

var poolAliexpressSolutionMerchantProfileGetAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSolutionMerchantProfileGetAPIResponse)
	},
}

// GetAliexpressSolutionMerchantProfileGetAPIResponse 从 sync.Pool 获取 AliexpressSolutionMerchantProfileGetAPIResponse
func GetAliexpressSolutionMerchantProfileGetAPIResponse() *AliexpressSolutionMerchantProfileGetAPIResponse {
	return poolAliexpressSolutionMerchantProfileGetAPIResponse.Get().(*AliexpressSolutionMerchantProfileGetAPIResponse)
}

// ReleaseAliexpressSolutionMerchantProfileGetAPIResponse 将 AliexpressSolutionMerchantProfileGetAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSolutionMerchantProfileGetAPIResponse(v *AliexpressSolutionMerchantProfileGetAPIResponse) {
	v.Reset()
	poolAliexpressSolutionMerchantProfileGetAPIResponse.Put(v)
}
