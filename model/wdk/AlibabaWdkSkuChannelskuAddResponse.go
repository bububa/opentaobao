package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
新增渠道商品 APIResponse
alibaba.wdk.sku.channelsku.add

盒马帮1期新增渠道商品
*/
type AlibabaWdkSkuChannelskuAddAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_sku_channelsku_add_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 结果
    
    Result   *AlibabaWdkSkuChannelskuAddApiResults `json:"result,omitempty" xml:"