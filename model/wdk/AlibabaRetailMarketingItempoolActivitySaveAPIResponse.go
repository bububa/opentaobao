package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailmarketingitempoolactivitysaveAPIResponse 【同城零售】保存商品池活动 API返回值
// alibaba.retail.marketing.itempool.activity.save
//
// 同城零售商品池活动保存
type AlibabaretailmarketingitempoolactivitysaveAPIResponse struct {
	model.CommonResponse
	AlibabaretailmarketingitempoolactivitysaveAPIResponseModel
}

// AlibabaretailmarketingitempoolactivitysaveAPIResponseModel is 【同城零售】保存商品池活动 成功返回结果
type AlibabaretailmarketingitempoolactivitysaveAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_itempool_activity_save_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	Result *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}
