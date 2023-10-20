package wdk

import (
	"encoding/xml"
	"sync"

	"github.com/bububa/opentaobao/model"
)

// AlibabaHmMarketingExpirePromotionCreateAPIResponse 短保优惠创建 API返回值
// alibaba.hm.marketing.expire.promotion.create
//
// 过期优惠优惠信息录入
type AlibabaHmMarketingExpirePromotionCreateAPIResponse struct {
	model.CommonResponse
	AlibabaHmMarketingExpirePromotionCreateAPIResponseModel
}

// Reset 清空结构体
func (m *AlibabaHmMarketingExpirePromotionCreateAPIResponse) Reset() {
	(&m.CommonResponse).Reset()
	(&m.AlibabaHmMarketingExpirePromotionCreateAPIResponseModel).Reset()
}

// AlibabaHmMarketingExpirePromotionCreateAPIResponseModel is 短保优惠创建 成功返回结果
type AlibabaHmMarketingExpirePromotionCreateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_hm_marketing_expire_promotion_create_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// data
	Datas []AlibabaHmMarketingExpirePromotionCreateT `json:"datas,omitempty" xml:"datas>alibaba_hm_marketing_expire_promotion_create_t,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// errorCode
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}

// Reset 清空结构体
func (m *AlibabaHmMarketingExpirePromotionCreateAPIResponseModel) Reset() {
	m.RequestId = ""
	m.Datas = m.Datas[:0]
	m.Message = ""
	m.FailCode = ""
	m.IsSuccess = false
}

var poolAlibabaHmMarketingExpirePromotionCreateAPIResponse = sync.Pool{
	New: func() any {
		return new(AlibabaHmMarketingExpirePromotionCreateAPIResponse)
	},
}

// GetAlibabaHmMarketingExpirePromotionCreateAPIResponse 从 sync.Pool 获取 AlibabaHmMarketingExpirePromotionCreateAPIResponse
func GetAlibabaHmMarketingExpirePromotionCreateAPIResponse() *AlibabaHmMarketingExpirePromotionCreateAPIResponse {
	return poolAlibabaHmMarketingExpirePromotionCreateAPIResponse.Get().(*AlibabaHmMarketingExpirePromotionCreateAPIResponse)
}

// ReleaseAlibabaHmMarketingExpirePromotionCreateAPIResponse 将 AlibabaHmMarketingExpirePromotionCreateAPIResponse 保存到 sync.Pool
func ReleaseAlibabaHmMarketingExpirePromotionCreateAPIResponse(v *AlibabaHmMarketingExpirePromotionCreateAPIResponse) {
	v.Reset()
	poolAlibabaHmMarketingExpirePromotionCreateAPIResponse.Put(v)
}
