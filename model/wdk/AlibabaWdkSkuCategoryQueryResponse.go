package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
商家类目获取接口 APIResponse
alibaba.wdk.sku.category.query

商家类目获取接口
*/
type AlibabaWdkSkuCategoryQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkSkuCategoryQueryResponse
}

type AlibabaWdkSkuCategoryQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_sku_category_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用结果
    
    Result   *AlibabaWdkSkuCategoryQueryApiResult `json:"result,omitempty" xml:"result,omitempty"`

    
}
