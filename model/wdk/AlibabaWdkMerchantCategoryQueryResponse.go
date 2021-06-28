package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
三江erp对接类目查询接口 APIResponse
alibaba.wdk.merchant.category.query

三江erp对接类目查询接口
*/
type AlibabaWdkMerchantCategoryQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_merchant_category_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // success
    
    Suc   bool `json:"suc,omitempty" xml:"