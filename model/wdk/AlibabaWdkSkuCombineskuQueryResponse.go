package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
组合商品查询接口 APIResponse
alibaba.wdk.sku.combinesku.query

查询组合商品接口
*/
type AlibabaWdkSkuCombineskuQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_sku_combinesku_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用结果
    
    Result   *AlibabaWdkSkuCombineskuQueryApiResults `json:"result,omitempty" xml:"