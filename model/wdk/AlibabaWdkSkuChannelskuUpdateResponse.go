package wdk

import (
    "encoding/xml"

    "github.com/bububa/opentaobao/model"
)

/* 
更新渠道商品 APIResponse
alibaba.wdk.sku.channelsku.update

批量更新渠道商品，商家通过Top接入
*/
type AlibabaWdkSkuChannelskuUpdateAPIResponse struct {
    model.CommonResponse
	RequestId     string         `json:"request_id,omitempty" xml:"alibaba_wdk_sku_channelsku_update_response>request_id,omitempty"`         // 平台颁发的每次请求访问的唯一标识

    // 调用结果
    
    Result   *AlibabaWdkSkuChannelskuUpdateApiResults `json:"result,omitempty" xml:"