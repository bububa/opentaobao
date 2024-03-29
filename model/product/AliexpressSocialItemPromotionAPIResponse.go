package product

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AliexpressSocialItemPromotionAPIResponse 获取推广链接 API返回值
// aliexpress.social.item.promotion
//
// 获取商品社交推广链接
type AliexpressSocialItemPromotionAPIResponse struct {
	model.CommonResponse
	AliexpressSocialItemPromotionAPIResponseModel
}

// Reset 清空结构体
func (m *AliexpressSocialItemPromotionAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AliexpressSocialItemPromotionAPIResponseModel).Reset()
}

// AliexpressSocialItemPromotionAPIResponseModel is 获取推广链接 成功返回结果
type AliexpressSocialItemPromotionAPIResponseModel struct {
	XMLName xml.Name `xml:"aliexpress_social_item_promotion_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 推广链接
	PromotionUrl string `json:"promotion_url,omitempty" xml:"promotion_url,omitempty"`
	// 错误码
	ResultCode string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 错误信息
	ResultMsg string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 是否成功
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AliexpressSocialItemPromotionAPIResponseModel) Reset() {
	m.RequestId = ""
	m.PromotionUrl = ""
	m.ResultCode = ""
	m.ResultMsg = ""
	m.IsSuccess = false
}

var poolAliexpressSocialItemPromotionAPIResponse = sync.Pool{
	New: func() any {
		return new(AliexpressSocialItemPromotionAPIResponse)
	},
}

// GetAliexpressSocialItemPromotionAPIResponse 从 sync.Pool 获取 AliexpressSocialItemPromotionAPIResponse
func GetAliexpressSocialItemPromotionAPIResponse() *AliexpressSocialItemPromotionAPIResponse {
	return poolAliexpressSocialItemPromotionAPIResponse.Get().(*AliexpressSocialItemPromotionAPIResponse)
}

// ReleaseAliexpressSocialItemPromotionAPIResponse 将 AliexpressSocialItemPromotionAPIResponse 保存到 sync.Pool
func ReleaseAliexpressSocialItemPromotionAPIResponse(v *AliexpressSocialItemPromotionAPIResponse) {
	v.Reset()
	poolAliexpressSocialItemPromotionAPIResponse.Put(v)
}
