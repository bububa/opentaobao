package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailmarketingbuygiftactivityqueryAPIResponse 查询单品买赠活动【同城零售】 API返回值
// alibaba.retail.marketing.buygift.activity.query
//
// 查询单品买赠活动【同城零售】
type AlibabaretailmarketingbuygiftactivityqueryAPIResponse struct {
	model.CommonResponse
	AlibabaretailmarketingbuygiftactivityqueryAPIResponseModel
}

// AlibabaretailmarketingbuygiftactivityqueryAPIResponseModel is 查询单品买赠活动【同城零售】 成功返回结果
type AlibabaretailmarketingbuygiftactivityqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_buygift_activity_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 错误码
	ErrNumber string `json:"err_number,omitempty" xml:"err_number,omitempty"`
	// 买赠活动查询结果
	Data *BuyGiftPromotionActivityDto `json:"data,omitempty" xml:"data,omitempty"`
	// 成功标识
	Succeed bool `json:"succeed,omitempty" xml:"succeed,omitempty"`
}
