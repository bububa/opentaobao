package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
组合商品查询接口 API返回值 
alibaba.wdk.sku.combinesku.query

查询组合商品接口
*/
type AlibabaWdkSkuCombineskuQueryAPIResponse struct {
    model.CommonResponse
    AlibabaWdkSkuCombineskuQueryResponse
}

// 组合商品查询接口 成功返回结果
type AlibabaWdkSkuCombineskuQueryResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_sku_combinesku_query_response"`
    // 平台颁发的每次请求访问的唯一标识
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
    // 调用结果
    Result   *AlibabaWdkSkuCombineskuQueryApiResults `json:"result,omitempty" xml:"result,omitempty"`
}
