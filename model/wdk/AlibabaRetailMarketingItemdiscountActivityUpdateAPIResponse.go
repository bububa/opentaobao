package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailmarketingitemdiscountactivityupdateAPIResponse 更新单品特价活动【同城零售】 API返回值
// alibaba.retail.marketing.itemdiscount.activity.update
//
// 同城零售单品特价活动更新
type AlibabaretailmarketingitemdiscountactivityupdateAPIResponse struct {
	model.CommonResponse
	AlibabaretailmarketingitemdiscountactivityupdateAPIResponseModel
}

// AlibabaretailmarketingitemdiscountactivityupdateAPIResponseModel is 更新单品特价活动【同城零售】 成功返回结果
type AlibabaretailmarketingitemdiscountactivityupdateAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_itemdiscount_activity_update_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	Result *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}
