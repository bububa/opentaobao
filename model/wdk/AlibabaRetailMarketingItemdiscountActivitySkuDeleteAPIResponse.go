package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailmarketingitemdiscountactivityskudeleteAPIResponse 删除特价活动商品 API返回值
// alibaba.retail.marketing.itemdiscount.activity.sku.delete
//
// 删除活动商品信息【同城零售】
type AlibabaretailmarketingitemdiscountactivityskudeleteAPIResponse struct {
	model.CommonResponse
	AlibabaretailmarketingitemdiscountactivityskudeleteAPIResponseModel
}

// AlibabaretailmarketingitemdiscountactivityskudeleteAPIResponseModel is 删除特价活动商品 成功返回结果
type AlibabaretailmarketingitemdiscountactivityskudeleteAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_itemdiscount_activity_sku_delete_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 操作结果
	Result *OctopusOpenResult `json:"result,omitempty" xml:"result,omitempty"`
}
