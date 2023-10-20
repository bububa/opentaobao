package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailmarketingbuygiftactivitysaveAPIResponse 【同城零售】单品买赠活动保存 API返回值
// alibaba.retail.marketing.buygift.activity.save
//
// 同城零售单品买赠活动保存
type AlibabaretailmarketingbuygiftactivitysaveAPIResponse struct {
	model.CommonResponse
	AlibabaretailmarketingbuygiftactivitysaveAPIResponseModel
}

// AlibabaretailmarketingbuygiftactivitysaveAPIResponseModel is 【同城零售】单品买赠活动保存 成功返回结果
type AlibabaretailmarketingbuygiftactivitysaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_buygift_activity_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	Result *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}
