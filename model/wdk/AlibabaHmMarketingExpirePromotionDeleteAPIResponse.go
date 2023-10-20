package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingExpirePromotionDeleteAPIResponse 短保优惠删除 API返回值
// alibaba.hm.marketing.expire.promotion.delete
//
// 短保优惠删除
type AlibabaHmMarketingExpirePromotionDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingExpirePromotionDeleteAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingExpirePromotionDeleteAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingExpirePromotionDeleteAPIResponseModel).Reset()
}

// AlibabaHmMarketingExpirePromotionDeleteAPIResponseModel is 短保优惠删除 成功返回结果
type AlibabaHmMarketingExpirePromotionDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_expire_promotion_delete_response"`
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
func (m *AlibabaHmMarketingExpirePromotionDeleteAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Datas = m.Datas[:0]
	m.Message = ""
	m.FailCode = ""
	m.IsSuccess = false
}

var poolAlibabaHmMarketingExpirePromotionDeleteAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingExpirePromotionDeleteAPIResponse)
	},
}

// GetAlibabaHmMarketingExpirePromotionDeleteAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingExpirePromotionDeleteAPIResponse
func GetAlibabaHmMarketingExpirePromotionDeleteAPIResponse() *AlibabaHmMarketingExpirePromotionDeleteAPIResponse {
	return poolAlibabaHmMarketingExpirePromotionDeleteAPIResponse.Get().(*AlibabaHmMarketingExpirePromotionDeleteAPIResponse)
}

// ReleaseAlibabaHmMarketingExpirePromotionDeleteAPIResponse 将 AlibabaHmMarketingExpirePromotionDeleteAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingExpirePromotionDeleteAPIResponse(v *AlibabaHmMarketingExpirePromotionDeleteAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingExpirePromotionDeleteAPIResponse.Put(v)
}
