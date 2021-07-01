package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaWdkMarketingExpirePromotionDeleteAPIResponse
短保优惠删除 API返回值
alibaba.wdk.marketing.expire.promotion.delete

短保优惠删除 */
type AlibabaWdkMarketingExpirePromotionDeleteAPIResponse struct {
	model.CommonResponse
	AlibabaWdkMarketingExpirePromotionDeleteAPIResponseModel
}

// AlibabaWdkMarketingExpirePromotionDeleteAPIResponseModel is 短保优惠删除 成功返回结果
type AlibabaWdkMarketingExpirePromotionDeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_wdk_marketing_expire_promotion_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// message
	Message string `json:"message,omitempty" xml:"message,omitempty"`
	// data
	Datas []ExpirePromotionResult `json:"datas,omitempty" xml:"datas>expire_promotion_result,omitempty"`
	// errorCode
	FailCode string `json:"fail_code,omitempty" xml:"fail_code,omitempty"`
	// success
	IsSuccess bool `json:"is_success,omitempty" xml:"is_success,omitempty"`
}
