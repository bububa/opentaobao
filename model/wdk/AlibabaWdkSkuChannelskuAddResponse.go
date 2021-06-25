package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
新增渠道商品 APIResponse
alibaba.wdk.sku.channelsku.add

盒马帮1期新增渠道商品
*/
type AlibabaWdkSkuChannelskuAddAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkSkuChannelskuAddResponse `json:"alibaba_wdk_sku_channelsku_add_response,omitempty"`
}

type AlibabaWdkSkuChannelskuAddResponse struct {

    // 结果
    Result   *AlibabaWdkSkuChannelskuAddApiResults `json:"result,omitempty"`

}
