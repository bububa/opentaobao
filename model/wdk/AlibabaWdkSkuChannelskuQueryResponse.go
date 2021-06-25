package wdk

import (
    "github.com/bububa/opentaobao/model"
)

/* 
查询渠道商品 APIResponse
alibaba.wdk.sku.channelsku.query

查询渠道商品
*/
type AlibabaWdkSkuChannelskuQueryAPIResponse struct {
    model.CommonResponse
    Response *AlibabaWdkSkuChannelskuQueryResponse `json:"alibaba_wdk_sku_channelsku_query_response,omitempty"`
}

type AlibabaWdkSkuChannelskuQueryResponse struct {

    // 调用结果
    Result   *AlibabaWdkSkuChannelskuQueryApiResults `json:"result,omitempty"`

}
