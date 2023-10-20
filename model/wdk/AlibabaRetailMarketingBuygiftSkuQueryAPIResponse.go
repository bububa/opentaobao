package wdk

import (
	"encoding/xml"

	"github.com/bububa/opentaobao/model"
)

// AlibabaretailmarketingbuygiftskuqueryAPIResponse 查询买赠活动商品【同城零售】 API返回值
// alibaba.retail.marketing.buygift.sku.query
//
// 查询买赠活动商品【同城零售】
type AlibabaretailmarketingbuygiftskuqueryAPIResponse struct {
	model.CommonResponse
	AlibabaretailmarketingbuygiftskuqueryAPIResponseModel
}

// AlibabaretailmarketingbuygiftskuqueryAPIResponseModel is 查询买赠活动商品【同城零售】 成功返回结果
type AlibabaretailmarketingbuygiftskuqueryAPIResponseModel struct {
	XMLName xml.Name `xml:"alibaba_retail_marketing_buygift_sku_query_response"`
	// 平台颁发的每次请求访问的唯一标识
	RequestId string `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// 买赠商品查询结果
	Data []BuyGiftActivitySkuDto `json:"data,omitempty" xml:"data>buy_gift_activity_sku_dto,omitempty"`
	// 错误信息
	ErrMessage string `json:"err_message,omitempty" xml:"err_message,omitempty"`
	// 错误码
	ErrNumber string `json:"err_number,omitempty" xml:"err_number,omitempty"`
	// 分页信息
	PageInfo *PageInfoDto `json:"page_info,omitempty" xml:"page_info,omitempty"`
	// 成功标识
	Succeed bool `json:"succeed,omitempty" xml:"succeed,omitempty"`
}
