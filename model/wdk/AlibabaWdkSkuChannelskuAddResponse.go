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
    AlibabaWdkSkuChannelskuAddResponse
}

type AlibabaWdkSkuChannelskuAddResponse struct {
    XMLName xml.Name `xml:"alibaba_wdk_sku_channelsku_add_response"`
    
	RequestId     string         `json:"request_id,omitempty" xml:"request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识
    

    // 结果
    
    Result   *AlibabaWdkSkuChannelskuAddApiResults `json:"result,omitempty" xml:"result,omitempty"`

    
}
