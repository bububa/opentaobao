package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingExpirePromotionCreateAPIResponse 短保优惠创建 API返回值
// alibaba.wdk.marketing.expire.promotion.create
//
// 过期优惠优惠信息录入
type AlibabaWdkMarketingExpirePromotionCreateAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingExpirePromotionCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingExpirePromotionCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingExpirePromotionCreateAPIResponseModel).Reset()
}

// AlibabaWdkMarketingExpirePromotionCreateAPIResponseModel is 短保优惠创建 成功返回结果
type AlibabaWdkMarketingExpirePromotionCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_expire_promotion_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Datas []AlibabaWdkMarketingExpirePromotionCreateT `json:"datas,omitempty" xml:"datas>alibaba_wdk_marketing_expire_promotion_create_t,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// errorCode
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingExpirePromotionCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Datas = m.Datas[:0]
	m.Message = ""
	m.FailCode = ""
	m.IsSuccess = false
}

var poolAlibabaWdkMarketingExpirePromotionCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingExpirePromotionCreateAPIResponse)
	},
}

// GetAlibabaWdkMarketingExpirePromotionCreateAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingExpirePromotionCreateAPIResponse
func GetAlibabaWdkMarketingExpirePromotionCreateAPIResponse() *AlibabaWdkMarketingExpirePromotionCreateAPIResponse {
	return poolAlibabaWdkMarketingExpirePromotionCreateAPIResponse.Get().(*AlibabaWdkMarketingExpirePromotionCreateAPIResponse)
}

// ReleaseAlibabaWdkMarketingExpirePromotionCreateAPIResponse 将 AlibabaWdkMarketingExpirePromotionCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingExpirePromotionCreateAPIResponse(v *AlibabaWdkMarketingExpirePromotionCreateAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingExpirePromotionCreateAPIResponse.Put(v)
}
