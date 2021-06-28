package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询商品 APIResponse
alibaba.wdk.sku.query

查询商品
*/
type AlibabaWdkSkuQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkSkuQueryResponse
}

type AlibabaWdkSkuQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_sku_query_response"`
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用结果
    
    Result   *AlibabaWdkSkuQueryApiResults `json:"result,omitempty" xml:"result,omitempty"`

    
}
