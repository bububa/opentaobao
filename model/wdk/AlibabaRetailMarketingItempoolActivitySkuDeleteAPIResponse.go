package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailmarketingitempoolactivityskudeleteAPIResponse 删除商品池活动商品【同城零售】 API返回值
// alibaba.retail.marketing.itempool.activity.sku.delete
//
// 删除商品池活动商品信息【同城零售】
type AlibabaretailmarketingitempoolactivityskudeleteAPIResponse struct {
	model.CommonResponse
	AlibabaretailmarketingitempoolactivityskudeleteAPIResponseModel
}

// AlibabaretailmarketingitempoolactivityskudeleteAPIResponseModel is 删除商品池活动商品【同城零售】 成功返回结果
type AlibabaretailmarketingitempoolactivityskudeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_itempool_activity_sku_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 出参
	Result *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}
