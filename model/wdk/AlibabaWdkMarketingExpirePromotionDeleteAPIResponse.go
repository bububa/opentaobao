package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaWdkMarketingExpirePromotionDeleteAPIResponse 短保优惠删除 API返回值
// alibaba.wdk.marketing.expire.promotion.delete
//
// 短保优惠删除
type AlibabaWdkMarketingExpirePromotionDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingExpirePromotionDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingExpirePromotionDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaWdkMarketingExpirePromotionDeleteAPIResponseModel).Reset()
}

// AlibabaWdkMarketingExpirePromotionDeleteAPIResponseModel is 短保优惠删除 成功返回结果
type AlibabaWdkMarketingExpirePromotionDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_expire_promotion_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Datas []ExpirePromotionResult `json:"datas,omitempty" xml:"datas>expire_promotion_result,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// errorCode
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaWdkMarketingExpirePromotionDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Datas = m.Datas[:0]
	m.Message = ""
	m.FailCode = ""
	m.IsSuccess = false
}

var poolAlibabaWdkMarketingExpirePromotionDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaWdkMarketingExpirePromotionDeleteAPIResponse)
	},
}

// GetAlibabaWdkMarketingExpirePromotionDeleteAPIResponse 从 sync.Pool 获取 AlibabaWdkMarketingExpirePromotionDeleteAPIResponse
func GetAlibabaWdkMarketingExpirePromotionDeleteAPIResponse() *AlibabaWdkMarketingExpirePromotionDeleteAPIResponse {
	return poolAlibabaWdkMarketingExpirePromotionDeleteAPIResponse.Get().(*AlibabaWdkMarketingExpirePromotionDeleteAPIResponse)
}

// ReleaseAlibabaWdkMarketingExpirePromotionDeleteAPIResponse 将 AlibabaWdkMarketingExpirePromotionDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaWdkMarketingExpirePromotionDeleteAPIResponse(v *AlibabaWdkMarketingExpirePromotionDeleteAPIResponse) {
	v.Reset()
	poolAlibabaWdkMarketingExpirePromotionDeleteAPIResponse.Put(v)
}
