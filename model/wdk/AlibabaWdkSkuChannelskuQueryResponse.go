package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
查询渠道商品 APIResponse
alibaba.wdk.sku.channelsku.query

查询渠道商品
*/
type AlibabaWdkSkuChannelskuQueryAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_sku_channelsku_query_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用结果
    
    Result   *AlibabaWdkSkuChannelskuQueryApiResults `json:"result,omitempty" xml:"